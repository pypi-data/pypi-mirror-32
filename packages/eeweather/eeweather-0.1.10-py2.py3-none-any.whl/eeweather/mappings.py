import numpy as np

from .connections import metadata_db_connection_proxy
from .exceptions import UnrecognizedUSAFIDError
from .stations import ISDStation
from .utils import lazy_property
from .validation import valid_zcta_or_raise

__all__ = (
    'MappingResult',
    'EmptyMapping',
    'ISDStationMapping',
    'zcta_closest_within_climate_zone',
    'zcta_naive_closest_high_quality',
    'zcta_naive_closest_medium_quality',
    'lat_long_naive_closest',
    'lat_long_closest_within_climate_zone',
    'oee_zcta',
    'oee_lat_long',
    'plot_mapping_results',
)


class CachedData(object):

    @lazy_property
    def isd_station_locations(self):
        conn = metadata_db_connection_proxy.get_connection()

        cur = conn.cursor()
        cur.execute('''
          select
            usaf_id
            , latitude
            , longitude
          from
            isd_station_metadata
          where
            quality = 'high'
        ''')

        isd_stations, isd_lats, isd_lngs = zip(*cur.fetchall())
        isd_lats = np.array(isd_lats)
        isd_lngs = np.array(isd_lngs)
        return isd_stations, isd_lats, isd_lngs

    @lazy_property
    def isd_station_metadata(self):
        conn = metadata_db_connection_proxy.get_connection()
        cur = conn.cursor()
        cur.execute('''
          select
            usaf_id
            , latitude
            , longitude
            , iecc_climate_zone
            , iecc_moisture_regime
            , ba_climate_zone
            , ca_climate_zone
          from
            isd_station_metadata
          where
            quality = 'high'
        ''')
        isd_station_metadata = {
            row[0]: {
                col[0]: val
                for col, val in zip(cur.description, row)
            }
            for row in cur.fetchall()
        }
        return isd_station_metadata

cached_data = CachedData()


def _unrecognized_zcta_or_empty_mapping(zcta):
    valid_zcta_or_raise(zcta)
    return EmptyMapping()


class MappingResult(object):
    '''Base class for mapping results.'''

    def is_empty(self):
        '''Returns False'''
        return False


class EmptyMapping(MappingResult):
    '''Returned when there is no valid mapping.

    Attributes
    ----------
    target_latitude : float or None
        latitude of target
    target_longitude : float or None
        longitude of target
    target_coords : tuple of (float, float)
        coordinates of target
    warnings : list of str
        Warnings for this mapping
    '''

    def __init__(self, target_latitude=None, target_longitude=None, warnings=None):
        if warnings is None:
            warnings = ['No mapping result was found.']
        self.target_latitude = target_latitude
        self.target_longitude = target_longitude
        self.target_coords = (target_latitude, target_longitude)
        self.warnings = warnings

    def __repr__(self):
        return 'EmptyMapping(warnings={})'.format(self.warnings)

    def is_empty(self):
        '''Returns True'''
        return True


class ISDStationMapping(MappingResult):
    '''A representation of a weather station mapping.

    Contains information about the target (e.g., lat/long), the mapped weather
    station, and any information about that mapping, such as distance between
    target and station, or warnings about the match.

    Attributes
    ----------
    isd_station : eeweather.ISDStation
        an object representing the ISD station
    target_latitude : float or None
        latitude of target
    target_longitude : float or None
        longitude of target
    target_coords : tuple of (float, float)
        coordinates of target
    distance_meters : int
        distance in meters from target coordinates to ISD station coordinates
    warnings : list of str
        Warnings for this mapping

    '''

    def __init__(
            self, usaf_id, target_latitude, target_longitude,
            distance_meters=None, warnings=None):

        if warnings is None:
            warnings = []

        self.isd_station = ISDStation(usaf_id)
        self.target_latitude = target_latitude
        self.target_longitude = target_longitude
        self.target_coords = (target_latitude, target_longitude)
        self.warnings = warnings

        if distance_meters is None:
            try:
                import pyproj
            except ImportError:  # pragma: no cover
                raise ImportError('Calculating distances requires pyproj.')
            geod = pyproj.Geod(ellps='WGS84')
            distance_meters = int(geod.inv(
                target_longitude, target_latitude,
                self.isd_station.longitude, self.isd_station.latitude)[2])

        if distance_meters > 50000:
            self.warnings.append(
                'Distance from target to weather station is greater than 50km.'
            )

        self.distance_meters = int(distance_meters)

    def __str__(self):
        return self.isd_station.usaf_id

    def __repr__(self):
        return "ISDStationMapping('{}', distance_meters={})".format(
            self.isd_station.usaf_id, self.distance_meters
        )

    def plot(self, target_label='target'):  # pragma: no cover
        ''' Plots this mapping on a map.'''
        try:
            import matplotlib.pyplot as plt
        except ImportError:
            raise ImportError('Plotting requires matplotlib.')

        try:
            import cartopy.crs as ccrs
            import cartopy.feature as cfeature
            import cartopy.io.img_tiles as cimgt
        except ImportError:
            raise ImportError('Plotting requires cartopy.')

        station = self.isd_station
        lat, lng = station.coords
        t_lat, t_lng = map(float, self.target_coords)

        # fiture
        fig = plt.figure(figsize=(16,8))

        # axes
        tiles = cimgt.StamenTerrain()
        ax = plt.subplot(1, 1, 1, projection=tiles.crs)

        # offsets for labels
        x_max = max([lng, t_lng])
        x_min = min([lng, t_lng])
        x_diff = x_max - x_min

        y_max = max([lat, t_lat])
        y_min = min([lat, t_lat])
        y_diff = y_max - y_min

        xoffset = x_diff * 0.05
        yoffset = y_diff * 0.05

        # minimum
        left = x_min - x_diff * 0.5
        right = x_max + x_diff * 0.5
        bottom = y_min - y_diff * 0.3
        top = y_max + y_diff * 0.3

        width_ratio = 2.
        height_ratio = 1.

        if (right - left) / (top - bottom) > width_ratio / height_ratio:
            # too short
            goal = (right - left) * height_ratio / width_ratio
            diff = goal - (top - bottom)
            bottom = bottom - diff/2.
            top = top + diff/2.
        else:
            # too skinny
            goal = (top - bottom) * width_ratio / height_ratio
            diff = goal - (right - left)
            left = left - diff/2.
            right = right + diff/2.

        ax.set_extent([left, right, bottom, top])

        # determine zoom level
        # tile size at level 1 = 64 km
        # level 2 = 32 km, level 3 = 16 km, etc, i.e. 128/(2^n) km
        N_TILES = 600  # (how many tiles approximately fit in distance)
        km = self.distance_meters / 1000.0
        zoom_level = int(np.log2(128 * N_TILES / km))

        ax.add_image(tiles, zoom_level)

        # line between
        plt.plot(
            [lng, t_lng], [lat, t_lat],
            linestyle='-', dashes=[2, 2], transform=ccrs.Geodetic()
        )

        # station
        ax.plot(lng, lat, 'ko', markersize=7, transform=ccrs.Geodetic())

        # target
        ax.plot(t_lng, t_lat, 'ro', markersize=7, transform=ccrs.Geodetic())

        # station label
        station_label = '{} ({})'.format(station.usaf_id, station.name)
        ax.text(lng + xoffset, lat + yoffset, station_label, transform=ccrs.Geodetic())

        # target label
        ax.text(t_lng + xoffset, t_lat + yoffset, target_label, transform=ccrs.Geodetic())

        # distance labels
        mid_lng = (lng + t_lng) / 2
        mid_lat = (lat + t_lat) / 2
        dist_text = '{:.01f} km'.format(km)
        ax.text(mid_lng + xoffset, mid_lat + yoffset, dist_text, transform=ccrs.Geodetic())

        plt.show()


def plot_mapping_results(mapping_results):  # pragma: no cover
    ''' Plot a list of mapping results on a map.

    Requires matplotlib and cartopy.

    Parameters
    ----------
    mapping_results : list of MappingResult objects
        Mapping results to plot
    '''
    try:
        import matplotlib.pyplot as plt
    except ImportError:
        raise ImportError('Plotting requires matplotlib.')

    try:
        import cartopy.crs as ccrs
        import cartopy.feature as cfeature
    except ImportError:
        raise ImportError('Plotting requires cartopy.')

    lats = []
    lngs = []
    t_lats = []
    t_lngs = []
    n_discards = 0
    for mapping_result in mapping_results:
        if not mapping_result.is_empty():
            lat, lng = mapping_result.isd_station.coords
            t_lat, t_lng = map(float, mapping_result.target_coords)
            lats.append(lat)
            lngs.append(lng)
            t_lats.append(t_lat)
            t_lngs.append(t_lng)
        else:
            n_discards += 1

    print('Discarded {} empty mappings'.format(n_discards))

    # figure
    fig = plt.figure(figsize=(60,60))

    # axes
    ax = plt.subplot(1, 1, 1, projection=ccrs.Mercator())

    # offsets for labels
    all_lngs = lngs + t_lngs
    all_lats = lats + t_lats
    x_max = max(all_lngs)  # lists
    x_min = min(all_lngs)
    x_diff = x_max - x_min

    y_max = max(all_lats)
    y_min = min(all_lats)
    y_diff = y_max - y_min

    # minimum
    x_pad = 0.1 * x_diff
    y_pad = 0.1 * y_diff
    left = x_min - x_pad
    right = x_max + x_pad
    bottom = y_min - y_pad
    top = y_max + y_pad

    width_ratio = 2.
    height_ratio = 1.

    if (right - left) / (top - bottom) > height_ratio / width_ratio:
        # too short
        goal = (right - left) * height_ratio / width_ratio
        diff = goal - (top - bottom)
        bottom = bottom - diff/2.
        top = top + diff/2.
    else:
        # too skinny
        goal = (top - bottom) * width_ratio / height_ratio
        diff = goal - (right - left)
        left = left - diff/2.
        right = right + diff/2.

    left = max(left, -179.9)
    right = min(right, 179.9)
    bottom = max([bottom, -89.9])
    top = min([top, 89.9])

    ax.set_extent([left, right, bottom, top])

    # OCEAN
    ax.add_feature(cfeature.NaturalEarthFeature(
        'physical', 'ocean', '50m',
        edgecolor='face', facecolor=cfeature.COLORS['water']
    ))

    # LAND
    ax.add_feature(cfeature.NaturalEarthFeature(
        'physical', 'land', '50m',
        edgecolor='face', facecolor=cfeature.COLORS['land']
    ))

    # BORDERS
    ax.add_feature(cfeature.NaturalEarthFeature(
        'cultural', 'admin_0_boundary_lines_land', '50m',
        edgecolor='black', facecolor='none'
    ))

    # LAKES
    ax.add_feature(cfeature.NaturalEarthFeature(
        'physical', 'lakes', '50m',
        edgecolor='face', facecolor=cfeature.COLORS['water']
    ))

    # COASTLINE
    ax.add_feature(cfeature.NaturalEarthFeature(
        'physical', 'coastline', '50m',
        edgecolor='black', facecolor='none'
    ))

    # lines between
    #for lat, t_lat, lng, t_lng in zip(lats, t_lats, lngs, t_lngs):
    ax.plot(
        [lngs, t_lngs], [lats, t_lats], color='k',
        linestyle='-', transform=ccrs.Geodetic(), linewidth=0.3
    )

    # stations
    ax.plot(lngs, lats, 'bo', markersize=1, transform=ccrs.Geodetic())

    plt.title('ZCTA to weather station mapping')

    plt.show()


def zcta_closest_within_climate_zone(zcta):
    '''Match ZCTA with closest high quality station within the same climate zone.

    Uses ZCTA centroid as target. "Within the same climate zone" means that it shares the same inclusion/exclusion status for each of the following climate zones:

    - IECC Climate Zone
    - IECC Moisture Regime
    - Building America Climate Zone
    - California Building Climate Zone Area

    Parameters
    ----------
    zcta : str
        ID of the target ZCTA.

    Returns
    -------
    mapping_result : eeweather.mappings.ISDStationMapping or eeweather.mappings.EmptyMapping
    '''
    conn = metadata_db_connection_proxy.get_connection()
    cur = conn.cursor()

    # try to find a station in the same climate zone
    cur.execute('''
      select
        z2i.usaf_id
        , z2i.distance_meters
        , zcta.latitude
        , zcta.longitude
      from
        zcta_to_isd_station z2i
        join isd_station_metadata isd on
          z2i.usaf_id = isd.usaf_id
        join zcta_metadata zcta on
          z2i.zcta_id = zcta.zcta_id
      where
        z2i.zcta_id = ?
        and isd.quality = 'high'
        and z2i.iecc_climate_zone_match
        and z2i.iecc_moisture_regime_match
        and z2i.ba_climate_zone_match
        and z2i.ca_climate_zone_match
      order by
        rank
      limit 1
    ''', (zcta,))

    match = cur.fetchone()
    if match is not None:
        (usaf_id, distance_meters, latitude, longitude) = match
        return ISDStationMapping(usaf_id, latitude, longitude, int(distance_meters))
    else:
        return _unrecognized_zcta_or_empty_mapping(zcta)


def zcta_naive_closest_high_quality(zcta):
    '''Match ZCTA with closest high quality station regardless of climate zone inclusion.

    Uses ZCTA centroid as target.

    Parameters
    ----------
    zcta : str
        ID of the target ZCTA.

    Returns
    -------
    mapping_result : eeweather.mappings.ISDStationMapping or eeweather.mappings.EmptyMapping
    '''
    conn = metadata_db_connection_proxy.get_connection()
    cur = conn.cursor()

    # try a naive distance-only search.
    cur.execute('''
      select
        zcta2isd.usaf_id
        , zcta.latitude
        , zcta.longitude
        , zcta2isd.distance_meters
      from
        zcta_to_isd_station zcta2isd
        join isd_station_metadata isd on
          zcta2isd.usaf_id = isd.usaf_id
        join zcta_metadata zcta on
          zcta2isd.zcta_id = zcta.zcta_id
      where
        zcta2isd.zcta_id = ?
        and isd.quality = 'high'
      order by
        rank
      limit 1
    ''', (zcta,))

    match = cur.fetchone()
    if match is not None:
        usaf_id, latitude, longitude, distance_meters = match
        return ISDStationMapping(usaf_id, latitude, longitude, int(distance_meters))
    else:
        return _unrecognized_zcta_or_empty_mapping(zcta)


def zcta_naive_closest_medium_quality(zcta):
    '''Match ZCTA with closest medium quality station regardless of climate zone
    inclusion.

    Uses ZCTA centroid as target.

    Parameters
    ----------
    zcta : str
        ID of the target ZCTA.

    Returns
    -------
    mapping_result : eeweather.mappings.ISDStationMapping or eeweather.mappings.EmptyMapping
    '''
    conn = metadata_db_connection_proxy.get_connection()
    cur = conn.cursor()

    # try a naive distance-only search.
    cur.execute('''
      select
        zcta2isd.usaf_id
        , zcta.latitude
        , zcta.longitude
        , zcta2isd.distance_meters
      from
        zcta_to_isd_station zcta2isd
        join isd_station_metadata isd on
          zcta2isd.usaf_id = isd.usaf_id
        join zcta_metadata zcta on
          zcta2isd.zcta_id = zcta.zcta_id
      where
        zcta2isd.zcta_id = ?
        and isd.quality = 'medium'
      order by
        rank
      limit 1
    ''', (zcta,))

    match = cur.fetchone()
    if match is not None:
        usaf_id, latitude, longitude, distance_meters = match
        return ISDStationMapping(usaf_id, latitude, longitude, int(distance_meters))
    else:
        return _unrecognized_zcta_or_empty_mapping(zcta)


def lat_long_naive_closest(latitude, longitude):
    ''' Find closest high quality ISD station regardless of climate zone match.

    Parameters
    ----------
    latitude : float
        Target latitude.
    longitude : float
        Target longitude.

    Returns
    -------
    mapping_result : eeweather.mappings.ISDStationMapping or eeweather.mappings.EmptyMapping
    '''
    try:
        import pyproj
    except ImportError:  # pragma: no cover
        raise ImportError('Matching by lat/lng requires pyproj.')

    isd_usaf_ids, isd_lats, isd_lngs = cached_data.isd_station_locations
    lats = np.tile(latitude, isd_lats.shape)
    lngs = np.tile(longitude, isd_lngs.shape)

    geod = pyproj.Geod(ellps='WGS84')
    dists = geod.inv(lngs, lats, isd_lngs, isd_lats)[2]
    idx = np.argmin(dists)
    usaf_id = isd_usaf_ids[idx]
    return ISDStationMapping(usaf_id, latitude, longitude)


def lat_long_closest_within_climate_zone(latitude, longitude):
    ''' Find closest ISD station within the same climate zone.

    "Within the same climate zone" means that it shares the same inclusion/exclusion status for each of the following climate zones:

    - IECC Climate Zone
    - IECC Moisture Regime
    - Building America Climate Zone
    - California Building Climate Zone Area

    Parameters
    ----------
    latitude : float
        Target latitude.
    longitude : float
        Target longitude.

    Returns
    -------
    mapping_result : eeweather.mappings.ISDStationMapping or eeweather.mappings.EmptyMapping
    '''
    try:
        import pyproj
    except ImportError:  # pragma: no cover
        raise ImportError('Matching by lat/lng requires pyproj.')

    from .api import get_lat_long_climate_zones
    climate_zones = get_lat_long_climate_zones(latitude, longitude)
    iecc_climate_zone = climate_zones['iecc_climate_zone']
    iecc_moisture_regime = climate_zones['iecc_moisture_regime']
    ca_climate_zone = climate_zones['ca_climate_zone']
    ba_climate_zone = climate_zones['ba_climate_zone']

    # outside all climate zones
    if (iecc_climate_zone is None
        and iecc_moisture_regime is None
        and ca_climate_zone is None
        and ba_climate_zone is None):
        return EmptyMapping(warnings=[
            'Target outside all known climate zones.'
        ])

    isd_station_metadata = cached_data.isd_station_metadata

    isd_usaf_ids, isd_lats, isd_lngs = zip(*[
        (usaf_id, metadata['latitude'], metadata['longitude'])
        for usaf_id, metadata in isd_station_metadata.items()
        if (metadata['iecc_climate_zone'] == iecc_climate_zone
            and metadata['iecc_moisture_regime'] == iecc_moisture_regime
            and metadata['ca_climate_zone'] == ca_climate_zone
            and metadata['ba_climate_zone'] == ba_climate_zone)
    ])

    if len(isd_usaf_ids) == 0:
        # haven't yet found a case where this applies, so untested.
        return EmptyMapping(warnings=[  # pragma: no cover
            'No weather stations in the target climate zone.'
        ])

    isd_lats = np.array(isd_lats)
    isd_lngs = np.array(isd_lngs)

    lats = np.tile(latitude, isd_lats.shape)
    lngs = np.tile(longitude, isd_lngs.shape)

    geod = pyproj.Geod(ellps='WGS84')
    dists = geod.inv(lngs, lats, isd_lngs, isd_lats)[2]
    indices = np.argsort(dists)

    idx = np.argmin(dists)
    distance_meters = dists[idx]
    usaf_id = isd_usaf_ids[idx]
    return ISDStationMapping(
        usaf_id, latitude, longitude, distance_meters=distance_meters)


def oee_zcta(zcta):
    '''Default ZCTA to ISD station mapping method.

    Determine the ZCTA centroid, attempt finding nearest match within climate zone, then fall back to the naive closest if that match is empty. Restricted to high quality stations.

    Parameters
    ----------
    zcta : str
        ID of the target ZCTA.

    Returns
    -------
    mapping_result : eeweather.mappings.ISDStationMapping or eeweather.mappings.EmptyMapping
    '''
    mapping_result = zcta_closest_within_climate_zone(zcta)
    if mapping_result.is_empty():
        mapping_result = zcta_naive_closest_high_quality(zcta)
        if not mapping_result.is_empty():
            mapping_result.warnings.append(
                'Mapped weather station is not in the same'
                ' climate zone as the centroid of the provided ZCTA.'
            )
    return mapping_result


def oee_lat_long(latitude, longitude):
    '''Default lat long to ISD station mapping method.

    First attempt finding nearest match within climate zone, then fall back
    to the naive closest if that match is empty. Restricted to high quality
    stations.

    Parameters
    ----------
    latitude : float
        Target latitude.
    longitude : float
        Target longitude.

    Returns
    -------
    mapping_result : eeweather.mappings.ISDStationMapping or eeweather.mappings.EmptyMapping
    '''
    mapping_result = lat_long_closest_within_climate_zone(latitude, longitude)
    if mapping_result.is_empty():
        mapping_result = lat_long_naive_closest(latitude, longitude)
        if not mapping_result.is_empty():
            mapping_result.warnings.append(
                'Mapped weather station is not in the same climate zone'
                ' as the provided lat/long point.'
            )
    return mapping_result
