#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Copyright (c) 2010, 2012, 2014, 2015.

# SMHI,
# Folkborgsvägen 1,
# Norrköping,
# Sweden

# Author(s):

#   Martin Raspaud <martin.raspaud@smhi.se>
#   Adam Dybbroe <adam.dybbroe@smhi.se>

# This file is part of mpop.

# mpop is free software: you can redistribute it and/or modify it under the
# terms of the GNU General Public License as published by the Free Software
# Foundation, either version 3 of the License, or (at your option) any later
# version.

# mpop is distributed in the hope that it will be useful, but WITHOUT ANY
# WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
# A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

# You should have received a copy of the GNU General Public License along with
# mpop.  If not, see <http://www.gnu.org/licenses/>.

"""Plugin for reading PPS's cloud products hdf files.
"""
import ConfigParser
import os.path
from mpop import CONFIG_PATH
import mpop.channel
import numpy as np
import pyresample.utils

import glob
from mpop.utils import get_logger
from mpop.projector import get_area_def
from datetime import datetime

LOG = get_logger('satin/msg_hdf')
COMPRESS_LVL = 6

ctype_lut = ['0: Not processed',
             '1: Cloud free land',
             '2: Cloud free sea',
             '3: Snow/ice contaminated land',
             '4: Snow/ice contaminated sea',
             '5: Very low cumiliform cloud',
             '6: Very low stratiform cloud',
             '7: Low cumiliform cloud',
             '8: Low stratiform cloud',
             '9: Medium level cumiliform cloud',
             '10: Medium level stratiform cloud',
             '11: High and opaque cumiliform cloud',
             '12: High and opaque stratiform cloud',
             '13:Very high and opaque cumiliform cloud',
             '14: Very high and opaque stratiform cloud',
             '15: Very thin cirrus cloud',
             '16: Thin cirrus cloud',
             '17: Thick cirrus cloud',
             '18: Cirrus above low or medium level cloud',
             '19: Fractional or sub-pixel cloud',
             '20: Undefined']
phase_lut = ['1: Not processed or undefined',
             '2: Water',
             '4: Ice',
             '8: Tb11 below 260K',
             '16: value not defined',
             '32: value not defined',
             '64: value not defined',
             '128: value not defined']
quality_lut = ['1: Land',
               '2: Coast',
               '4: Night',
               '8: Twilight',
               '16: Sunglint',
               '32: High terrain',
               '64: Low level inversion',
               '128: Nwp data present',
               '256: Avhrr channel missing',
               '512: Low quality',
               '1024: Reclassified after spatial smoothing',
               '2048: Stratiform-Cumuliform Distinction performed',
               '4096: bit not defined',
               '8192: bit not defined',
               '16384: bit not defined',
               '32768: bit not defined']
processing_flags_lut = ["1: Not processed",
                        "2: Cloudy",
                        "4: Opaque cloud",
                        "8: RTTOV IR simulations available",
                        "16: Missing NWP data",
                        "32: thermal inversion avaliable",
                        "64: Missing AVHRR data",
                        "128: RTTOV IR simulation applied",
                        "256: Windowing technique applied",
                        "512: bit not defined",
                        "1024: bit not defined",
                        "2048: bit not defined",
                        "4096: bit not defined",
                        "8192: bit not defined",
                        "16384: Quality estimation avaliable",
                        "32768: Low confidence"]


class UnknownChannelError(Exception):
    pass


def pcs_def_from_region(region):
    items = region.proj_dict.items()
    return ' '.join([t[0] + '=' + t[1] for t in items])


def _get_area_extent(cfac, lfac, coff, loff, numcols, numlines):
    """Get the area extent from msg parameters.
    """

    xur = (numcols - coff) * 2 ** 16 / (cfac * 1.0)
    xur = np.deg2rad(xur) * 35785831.0
    xll = (-1 - coff) * 2 ** 16 / (cfac * 1.0)
    xll = np.deg2rad(xll) * 35785831.0
    xres = (xur - xll) / numcols
    xur, xll = xur - xres / 2, xll + xres / 2
    yll = (numlines - loff) * 2 ** 16 / (-lfac * 1.0)
    yll = np.deg2rad(yll) * 35785831.0
    yur = (-1 - loff) * 2 ** 16 / (-lfac * 1.0)
    yur = np.deg2rad(yur) * 35785831.0
    yres = (yur - yll) / numlines
    yll, yur = yll + yres / 2, yur - yres / 2
    return xll, yll, xur, yur


def get_area_extent(filename):
    """Get the area extent of the data in *filename*.
    """
    import h5py
    h5f = h5py.File(filename, 'r')
    aex = _get_area_extent(h5f.attrs["CFAC"],
                           h5f.attrs["LFAC"],
                           h5f.attrs["COFF"],
                           h5f.attrs["LOFF"],
                           h5f.attrs["NC"],
                           h5f.attrs["NL"])
    h5f.close()
    return aex


def _get_palette(h5f, dsname):
    try:
        p = h5f[dsname].attrs['PALETTE']
        return h5f[p].value
    except KeyError:
        return None


class InfoObject(object):

    """Simple data and info container.
    """

    def __init__(self):
        self.info = {}
        self.data = None

# ----------------------------------------


class MsgCloudTypeData(object):

    """NWCSAF/MSG Cloud Type data layer
    """

    def __init__(self):
        self.data = None
        self.scaling_factor = 1
        self.offset = 0
        self.num_of_lines = 0
        self.num_of_columns = 0
        self.product = ""
        self.id = ""


class MsgCloudType(mpop.channel.GenericChannel):

    """NWCSAF/MSG Cloud Type data structure as retrieved from HDF5
    file. Resolution sets the nominal resolution of the data.
    """

    def __init__(self):
        mpop.channel.GenericChannel.__init__(self, "CloudType")
        self.filled = False
        self.name = "CloudType"
        self.package = ""
        self.saf = ""
        self.product_name = ""
        self.num_of_columns = 0
        self.num_of_lines = 0
        self.projection_name = ""
        self.pcs_def = ""
        self.xscale = 0
        self.yscale = 0
        self.ll_lon = 0.0
        self.ll_lat = 0.0
        self.ur_lon = 0.0
        self.ur_lat = 0.0
        self.region_name = ""
        self.cfac = 0
        self.lfac = 0
        self.coff = 0
        self.loff = 0
        self.nb_param = 0
        self.gp_sc_id = 0
        self.image_acquisition_time = 0
        self.spectral_channel_id = 0
        self.nominal_product_time = 0
        self.sgs_product_quality = 0
        self.sgs_product_completeness = 0
        self.product_algorithm_version = ""
        self.cloudtype = None
        self.processing_flags = None
        self.cloudphase = None
        self.shape = None
        self.satid = ""
        self.qc_straylight = -1
        self.cloudtype_palette = None
        self.cloudphase_palette = None

    def __str__(self):
        try:
            shape = self.cloudtype.shape
        except AttributeError:
            shape = self.shape

        return ("'%s: shape %s, resolution %sm'" %
                (self.name,
                 shape,
                 self.resolution))

    def is_loaded(self):
        """Tells if the channel contains loaded data.
        """
        return self.filled

# ------------------------------------------------------------------
    def read(self, filename):
        """Reader for the NWCSAF/MSG cloudtype. Use *filename* to read data.
        """
        import h5py

        self.cloudtype = MsgCloudTypeData()
        self.processing_flags = MsgCloudTypeData()
        self.cloudphase = MsgCloudTypeData()

        LOG.debug("Filename = <" + str(filename) + ">")
        h5f = h5py.File(filename, 'r')
        # pylint: disable-msg=W0212
        self.package = h5f.attrs["PACKAGE"]
        self.saf = h5f.attrs["SAF"]
        self.product_name = h5f.attrs["PRODUCT_NAME"]
        self.num_of_columns = h5f.attrs["NC"]
        self.num_of_lines = h5f.attrs["NL"]
        self.projection_name = h5f.attrs["PROJECTION_NAME"]
        self.region_name = h5f.attrs["REGION_NAME"]
        self.cfac = h5f.attrs["CFAC"]
        self.lfac = h5f.attrs["LFAC"]
        self.coff = h5f.attrs["COFF"]
        self.loff = h5f.attrs["LOFF"]
        self.nb_param = h5f.attrs["NB_PARAMETERS"]
        self.gp_sc_id = h5f.attrs["GP_SC_ID"]
        self.image_acquisition_time = h5f.attrs["IMAGE_ACQUISITION_TIME"]
        self.spectral_channel_id = h5f.attrs["SPECTRAL_CHANNEL_ID"]
        self.nominal_product_time = datetime.strptime(h5f.attrs["NOMINAL_PRODUCT_TIME"],
                                                      "%Y%m%d%H%M")
        self.sgs_product_quality = h5f.attrs["SGS_PRODUCT_QUALITY"]
        self.sgs_product_completeness = h5f.attrs["SGS_PRODUCT_COMPLETENESS"]
        self.product_algorithm_version = h5f.attrs[
            "PACKAGE"] + h5f.attrs["PRODUCT_NAME"] + h5f.attrs["PRODUCT_ALGORITHM_VERSION"]
        # pylint: enable-msg=W0212
        # ------------------------

        # The cloudtype data
        if 'CT' not in h5f.keys():
            raise IOError("No field CT in product " +
                          str(self.product_name) +
                          "\n\tPerhaps you have found the wrong file for this product?")
        h5d = h5f['CT']
        self.cloudtype.data = h5d[:, :]
        self.cloudtype.scaling_factor = h5d.attrs["SCALING_FACTOR"]
        self.cloudtype.offset = h5d.attrs["OFFSET"]
        self.cloudtype.num_of_lines = h5d.attrs["N_LINES"]
        self.cloudtype.num_of_columns = h5d.attrs["N_COLS"]
        self.shape = (self.cloudtype.num_of_lines,
                      self.cloudtype.num_of_columns)
        self.cloudtype.product = h5d.attrs["PRODUCT"]
        self.cloudtype.id = h5d.attrs["ID"]
        self.cloudtype_palette = _get_palette(h5f, 'CT')
        # ------------------------

        # The cloud phase data
        h5d = h5f['CT_PHASE']
        self.cloudphase.data = h5d[:, :]
        self.cloudphase.scaling_factor = h5d.attrs["SCALING_FACTOR"]
        self.cloudphase.offset = h5d.attrs["OFFSET"]
        self.cloudphase.num_of_lines = h5d.attrs["N_LINES"]
        self.cloudphase.num_of_columns = h5d.attrs["N_COLS"]
        self.cloudphase.product = h5d.attrs["PRODUCT"]
        self.cloudphase.id = h5d.attrs["ID"]
        self.cloudphase_palette = _get_palette(h5f, 'CT_PHASE')

        # ------------------------

        # The cloudtype processing/quality flags
        h5d = h5f['CT_QUALITY']
        self.processing_flags.data = h5d[:, :]
        self.processing_flags.scaling_factor = \
            h5d.attrs["SCALING_FACTOR"]
        self.processing_flags.offset = h5d.attrs["OFFSET"]
        self.processing_flags.num_of_lines = h5d.attrs["N_LINES"]
        self.processing_flags.num_of_columns = h5d.attrs["N_COLS"]
        self.processing_flags.product = h5d.attrs["PRODUCT"]
        self.processing_flags.id = h5d.attrs["ID"]
        # ------------------------

        h5f.close()

        self.cloudtype = self.cloudtype.data
        self.cloudphase = self.cloudphase.data
        self.processing_flags = self.processing_flags.data

        self.area = get_area_from_file(filename)

        self.filled = True

    def save(self, filename, **kwargs):
        """Save the current cloudtype object to hdf *filename*, in pps format.
        """
        import h5py
        ctype = self.convert2pps()
        LOG.info("Saving CType hdf file...")
        ctype.save(filename, **kwargs)
        h5f = h5py.File(filename, mode="a")
        h5f.attrs["straylight_contaminated"] = self.qc_straylight
        h5f.close()
        LOG.info("Saving CType hdf file done !")

    def project(self, coverage):
        """Remaps the NWCSAF/MSG Cloud Type to cartographic map-projection on
        area give by a pre-registered area-id. Faster version of msg_remap!
        """
        LOG.info("Projecting channel %s..." % (self.name))

        region = coverage.out_area
        dest_area = region.area_id

        retv = MsgCloudType()

        retv.name = self.name
        retv.package = self.package
        retv.saf = self.saf
        retv.product_name = self.product_name
        retv.region_name = dest_area
        retv.cfac = self.cfac
        retv.lfac = self.lfac
        retv.coff = self.coff
        retv.loff = self.loff
        retv.nb_param = self.nb_param
        retv.gp_sc_id = self.gp_sc_id
        retv.image_acquisition_time = self.image_acquisition_time
        retv.spectral_channel_id = self.spectral_channel_id
        retv.nominal_product_time = self.nominal_product_time
        retv.sgs_product_quality = self.sgs_product_quality
        retv.sgs_product_completeness = self.sgs_product_completeness
        retv.product_algorithm_version = self.product_algorithm_version

        retv.cloudtype = coverage.project_array(self.cloudtype)
        retv.cloudtype_palette = self.cloudtype_palette

        retv.cloudphase = coverage.project_array(self.cloudphase)
        retv.cloudphase_palette = self.cloudphase_palette

        retv.processing_flags = \
            coverage.project_array(self.processing_flags)

        retv.qc_straylight = self.qc_straylight
        retv.region_name = dest_area
        retv.area = region
        retv.projection_name = region.proj_id

        retv.pcs_def = pcs_def_from_region(region)

        retv.num_of_columns = region.x_size
        retv.num_of_lines = region.y_size
        retv.xscale = region.pixel_size_x
        retv.yscale = region.pixel_size_y

        import pyproj
        prj = pyproj.Proj(region.proj4_string)
        aex = region.area_extent
        lonur, latur = prj(aex[2], aex[3], inverse=True)
        lonll, latll = prj(aex[0], aex[1], inverse=True)
        retv.ll_lon = lonll
        retv.ll_lat = latll
        retv.ur_lon = lonur
        retv.ur_lat = latur

        self.shape = region.shape

        retv.filled = True
        retv.resolution = self.resolution

        return retv

    def oldconvert2pps(self):
        """Converts the NWCSAF/MSG Cloud Type to the PPS format,
        in order to have consistency in output format between PPS and MSG.
        """
        retv = PpsCloudType()
        retv.region = SafRegion()
        retv.region.xsize = self.num_of_columns
        retv.region.ysize = self.num_of_lines
        retv.region.id = self.region_name
        retv.region.pcs_id = self.projection_name

        retv.region.pcs_def = pcs_def_from_region(self.area)
        retv.region.area_extent = self.area.area_extent
        retv.satellite_id = self.satid

        luts = pps_luts()
        retv.cloudtype_lut = luts[0]
        retv.phaseflag_lut = []
        retv.qualityflag_lut = []
        retv.cloudtype_des = "MSG SEVIRI Cloud Type"
        retv.qualityflag_des = 'MSG SEVIRI bitwise quality/processing flags'
        retv.phaseflag_des = 'MSG SEVIRI Cloud phase flags'

        retv.cloudtype = self.cloudtype.astype('B')
        retv.phaseflag = self.cloudphase.astype('B')
        retv.qualityflag = ctype_procflags2pps(self.processing_flags)

        return retv

    def convert2pps(self):
        from mpop.satin.nwcsaf_pps import CloudType
        from nwcsaf_formats.pps_conversions import (old_ctype_palette,
                                                    old_ctype_palette_data)

        retv = CloudType()

        region_type = np.dtype([('area_extent', '<f8', (4,)),
                                ('xsize', '<i4'),
                                ('ysize', '<i4'),
                                ('xscale', '<f4'),
                                ('yscale', '<f4'),
                                ('lat_0', '<f4'),
                                ('lon_0', '<f4'),
                                ('lat_ts', '<f4'),
                                ('id', 'S64'),
                                ('name', 'S64'),
                                ('pcs_id', 'S64'),
                                ('pcs_def', 'S128')])

        region = np.zeros((1, ), dtype=region_type)

        region["xsize"] = self.num_of_columns
        region["ysize"] = self.num_of_lines
        region["id"] = self.region_name
        region["pcs_id"] = self.projection_name

        region["pcs_def"] = pcs_def_from_region(self.area)
        region["area_extent"] = self.area.area_extent

        retv.region = InfoObject()
        retv.region.data = region
        retv._keys.append("region")

        retv.Region = region_type
        retv._keys.append("Region")

        retv._md["time_slot"] = self.nominal_product_time
        retv._md["description"] = np.string_(
            "MSG SEVIRI Cloud Type")
        retv._md["satellite_id"] = np.string_(self.satid)
        retv._md["sec_1970"] = np.uint64(0)
        retv._md["version"] = np.string_(self.product_algorithm_version)
        retv._md["orbit_number"] = np.uint64(0)

        retv.PALETTE = InfoObject()
        retv.PALETTE.data = old_ctype_palette_data()
        retv.PALETTE.info["CLASS"] = np.string_("PALETTE\0")
        retv.PALETTE.info["PAL_COLORMODEL"] = np.string_("RGB\0")
        retv.PALETTE.info["PAL_TYPE"] = np.string_("STANDARD8\0")
        retv.PALETTE.info["PAL_VERSION"] = np.string_("1.2\0")
        retv._keys.append("PALETTE")

        namelist = np.dtype([('outval_name', 'S128')])
        retv.OutputValueNameList = namelist
        retv._keys.append("OutputValueNameList")

        retv.cloudtype = InfoObject()
        retv.cloudtype.info["output_value_namelist"] = np.array(ctype_lut,
                                                                dtype=namelist)
        retv.cloudtype.info["CLASS"] = np.string_("IMAGE\0")
        retv.cloudtype.info["IMAGE_VERSION"] = np.string_("1.2\0")
        retv._refs[("cloudtype", "PALETTE")] = np.string_("PALETTE\0")
        retv.cloudtype.info["description"] = np.string_(
            "MSG SEVIRI Cloud Type")
        retv.cloudtype.data = self.cloudtype.astype('B')
        retv._projectables.append("cloudtype")

        # retv.PHASE_PALETTE = InfoObject()
        # retv.PHASE_PALETTE.data = self.cloudphase_palette
        # retv.PHASE_PALETTE.info["CLASS"] = np.string_("PALETTE")
        # retv.PHASE_PALETTE.info["PAL_COLORMODEL"] = np.string_("RGB")
        # retv.PHASE_PALETTE.info["PAL_TYPE"] = np.string_("STANDARD8")
        # retv.PHASE_PALETTE.info["PAL_VERSION"] = np.string_("1.2")
        # retv._keys.append("PHASE_PALETTE")

        retv.phase_flag = InfoObject()
        retv.phase_flag.info["output_value_namelist"] = np.array(phase_lut,
                                                                 dtype=namelist)
        # retv.phase_flag.info["CLASS"] = np.string_("IMAGE")
        # retv.phase_flag.info["IMAGE_VERSION"] = np.string_("1.2")
        # retv._refs[("phase_flag", "PALETTE")] = np.string_("PHASE_PALETTE")
        retv.phase_flag.info["description"] = np.string_(
            'MSG SEVIRI Cloud phase flags')
        retv.phase_flag.data = self.cloudphase.astype('B')
        retv._projectables.append("phase_flag")

        retv.quality_flag = InfoObject()
        retv.quality_flag.info["output_value_namelist"] = np.array(quality_lut,
                                                                   dtype=namelist)
        retv.quality_flag.info[
            "description"] = np.string_('MSG SEVIRI bitwise quality/processing flags')
        retv.quality_flag.data = ctype_procflags2pps(self.processing_flags)
        retv._projectables.append("quality_flag")

        retv.save = retv.write
        return retv

    def convert2nordrad(self):
        return NordRadCType(self)


class MsgCTTHData(object):

    """CTTH data object.
    """

    def __init__(self):
        self.data = None
        self.scaling_factor = 1
        self.offset = 0
        self.num_of_lines = 0
        self.num_of_columns = 0
        self.product = ""
        self.id = ""


class MsgCTTH(mpop.channel.GenericChannel):

    """CTTH channel.
    """

    def __init__(self, resolution=None):
        mpop.channel.GenericChannel.__init__(self, "CTTH")
        self.filled = False
        self.name = "CTTH"
        self.resolution = resolution
        self.package = ""
        self.saf = ""
        self.product_name = ""
        self.num_of_columns = 0
        self.num_of_lines = 0
        self.projection_name = ""
        self.region_name = ""
        self.cfac = 0
        self.lfac = 0
        self.coff = 0
        self.loff = 0
        self.nb_param = 0
        self.gp_sc_id = 0
        self.image_acquisition_time = 0
        self.spectral_channel_id = 0
        self.nominal_product_time = 0
        self.sgs_product_quality = 0
        self.sgs_product_completeness = 0
        self.product_algorithm_version = ""
        self.cloudiness = None  # Effective cloudiness
        self.processing_flags = None
        self.height = None
        self.temperature = None
        self.pressure = None
        self.satid = ""

    def __str__(self):
        return ("'%s: shape %s, resolution %sm'" %
                (self.name,
                 self.shape,
                 self.resolution))

    def is_loaded(self):
        """Tells if the channel contains loaded data.
        """
        return self.filled

    def read(self, filename, calibrate=True):
        import h5py

        self.cloudiness = MsgCTTHData()  # Effective cloudiness
        self.temperature = MsgCTTHData()
        self.height = MsgCTTHData()
        self.pressure = MsgCTTHData()
        self.processing_flags = MsgCTTHData()

        h5f = h5py.File(filename, 'r')

        # The header
        # pylint: disable-msg=W0212
        self.package = h5f.attrs["PACKAGE"]
        self.saf = h5f.attrs["SAF"]
        self.product_name = h5f.attrs["PRODUCT_NAME"]
        self.num_of_columns = h5f.attrs["NC"]
        self.num_of_lines = h5f.attrs["NL"]
        self.projection_name = h5f.attrs["PROJECTION_NAME"]
        self.region_name = h5f.attrs["REGION_NAME"]
        self.cfac = h5f.attrs["CFAC"]
        self.lfac = h5f.attrs["LFAC"]
        self.coff = h5f.attrs["COFF"]
        self.loff = h5f.attrs["LOFF"]
        self.nb_param = h5f.attrs["NB_PARAMETERS"]
        self.gp_sc_id = h5f.attrs["GP_SC_ID"]
        self.image_acquisition_time = h5f.attrs["IMAGE_ACQUISITION_TIME"]
        self.spectral_channel_id = h5f.attrs["SPECTRAL_CHANNEL_ID"]
        self.nominal_product_time = datetime.strptime(h5f.attrs["NOMINAL_PRODUCT_TIME"],
                                                      "%Y%m%d%H%M")
        self.sgs_product_quality = h5f.attrs["SGS_PRODUCT_QUALITY"]
        self.sgs_product_completeness = h5f.attrs["SGS_PRODUCT_COMPLETENESS"]
        self.product_algorithm_version = h5f.attrs[
            "PACKAGE"] + h5f.attrs["PRODUCT_NAME"] + h5f.attrs["PRODUCT_ALGORITHM_VERSION"]
        # pylint: enable-msg=W0212
        # ------------------------

        # The CTTH cloudiness data
        if 'CTTH_EFFECT' not in h5f.keys():
            raise IOError("No field CTTH_EFFECT in product " +
                          str(self.product_name) +
                          "\n\tPerhaps you have found the wrong file for this product?")
        h5d = h5f['CTTH_EFFECT']
        self.cloudiness.data = h5d[:, :]
        self.cloudiness.scaling_factor = \
            h5d.attrs["SCALING_FACTOR"]
        self.cloudiness.offset = h5d.attrs["OFFSET"]
        self.cloudiness.num_of_lines = h5d.attrs["N_LINES"]
        self.cloudiness.num_of_columns = h5d.attrs["N_COLS"]
        self.cloudiness.product = h5d.attrs["PRODUCT"]
        self.cloudiness.id = h5d.attrs["ID"]

        self.cloudiness.data = np.ma.masked_equal(self.cloudiness.data, 255)
        self.cloudiness = np.ma.masked_equal(self.cloudiness.data, 0)
        self.cloudiness_palette = _get_palette(h5f, 'CTTH_EFFECT')

        # ------------------------

        # The CTTH temperature data
        h5d = h5f['CTTH_TEMPER']
        self.temperature.data = h5d[:, :]
        self.temperature.scaling_factor = \
            h5d.attrs["SCALING_FACTOR"]
        self.temperature.offset = h5d.attrs["OFFSET"]
        self.temperature.num_of_lines = h5d.attrs["N_LINES"]
        self.shape = (self.temperature.num_of_lines,
                      self.temperature.num_of_columns)
        self.temperature.num_of_columns = h5d.attrs["N_COLS"]
        self.temperature.product = h5d.attrs["PRODUCT"]
        self.temperature.id = h5d.attrs["ID"]

        self.temperature.data = np.ma.masked_equal(self.temperature.data, 0)
        if calibrate:
            self.temperature = (self.temperature.data *
                                self.temperature.scaling_factor +
                                self.temperature.offset)
        else:
            self.temperature = self.temperature.data
        self.temperature_palette = _get_palette(h5f, 'CTTH_TEMPER')

        # ------------------------

        # The CTTH pressure data
        h5d = h5f['CTTH_PRESS']
        self.pressure.data = h5d[:, :]
        self.pressure.scaling_factor = \
            h5d.attrs["SCALING_FACTOR"]
        self.pressure.offset = h5d.attrs["OFFSET"]
        self.pressure.num_of_lines = h5d.attrs["N_LINES"]
        self.pressure.num_of_columns = h5d.attrs["N_COLS"]
        self.pressure.product = h5d.attrs["PRODUCT"]
        self.pressure.id = h5d.attrs["ID"]

        self.pressure.data = np.ma.masked_equal(self.pressure.data, 255)
        self.pressure.data = np.ma.masked_equal(self.pressure.data, 0)
        if calibrate:
            self.pressure = (self.pressure.data *
                             self.pressure.scaling_factor +
                             self.pressure.offset)
        else:
            self.pressure = self.pressure.data
        self.pressure_palette = _get_palette(h5f, 'CTTH_PRESS')

        # ------------------------

        # The CTTH height data
        h5d = h5f['CTTH_HEIGHT']
        self.height.data = h5d[:, :]
        self.height.scaling_factor = \
            h5d.attrs["SCALING_FACTOR"]
        self.height.offset = h5d.attrs["OFFSET"]
        self.height.num_of_lines = h5d.attrs["N_LINES"]
        self.height.num_of_columns = h5d.attrs["N_COLS"]
        self.height.product = h5d.attrs["PRODUCT"]
        self.height.id = h5d.attrs["ID"]

        self.height.data = np.ma.masked_equal(self.height.data, 255)
        self.height.data = np.ma.masked_equal(self.height.data, 0)
        if calibrate:
            self.height = (self.height.data *
                           self.height.scaling_factor +
                           self.height.offset)
        else:
            self.height = self.height.data
        self.height_palette = _get_palette(h5f, 'CTTH_HEIGHT')

        # ------------------------

        # The CTTH processing/quality flags
        h5d = h5f['CTTH_QUALITY']
        self.processing_flags.data = h5d[:, :]
        self.processing_flags.scaling_factor = \
            h5d.attrs["SCALING_FACTOR"]
        self.processing_flags.offset = h5d.attrs["OFFSET"]
        self.processing_flags.num_of_lines = \
            h5d.attrs["N_LINES"]
        self.processing_flags.num_of_columns = \
            h5d.attrs["N_COLS"]
        self.processing_flags.product = h5d.attrs["PRODUCT"]
        self.processing_flags.id = h5d.attrs["ID"]

        self.processing_flags = \
            np.ma.masked_equal(self.processing_flags.data, 0)

        h5f.close()

        self.shape = self.height.shape

        self.area = get_area_from_file(filename)

        self.filled = True

    def save(self, filename, **kwargs):
        """Save the current CTTH channel to HDF5 format.
        """
        ctth = self.convert2pps()
        LOG.info("Saving CTTH hdf file...")
        ctth.save(filename, **kwargs)
        LOG.info("Saving CTTH hdf file done !")

    def project(self, coverage):
        """Project the current CTTH channel along the *coverage*
        """
        dest_area = coverage.out_area
        dest_area_id = dest_area.area_id

        retv = MsgCTTH()

        retv.temperature = coverage.project_array(self.temperature)
        retv.height = coverage.project_array(self.height)
        retv.pressure = coverage.project_array(self.pressure)
        retv.cloudiness = coverage.project_array(self.cloudiness)
        retv.processing_flags = \
            coverage.project_array(self.processing_flags)

        retv.area = dest_area
        retv.region_name = dest_area_id
        retv.projection_name = dest_area.proj_id
        retv.num_of_columns = dest_area.x_size
        retv.num_of_lines = dest_area.y_size

        retv.shape = dest_area.shape

        retv.name = self.name
        retv.resolution = self.resolution
        retv.filled = True

        return retv

# ------------------------------------------------------------------
    def oldconvert2pps(self):
        """Convert the current CTTH channel to pps format.
        """
        retv = PpsCTTH()
        retv.region = SafRegion()
        retv.region.xsize = self.num_of_columns
        retv.region.ysize = self.num_of_lines
        retv.region.id = self.region_name
        retv.region.pcs_id = self.projection_name
        retv.region.pcs_def = pcs_def_from_region(self.area)
        retv.region.area_extent = self.area.area_extent
        retv.satellite_id = self.satid

        retv.processingflag_lut = []
        retv.des = "MSG SEVIRI Cloud Top Temperature & Height"
        retv.ctt_des = "MSG SEVIRI cloud top temperature (K)"
        retv.ctp_des = "MSG SEVIRI cloud top pressure (hPa)"
        retv.ctp_des = "MSG SEVIRI cloud top height (m)"
        retv.cloudiness_des = "MSG SEVIRI effective cloudiness (%)"
        retv.processingflag_des = 'MSG SEVIRI bitwise quality/processing flags'

        retv.t_gain = 1.0
        retv.t_intercept = 100.0
        retv.t_nodata = 255

        retv.temperature = ((self.temperature - retv.t_intercept) /
                            retv.t_gain).filled(retv.t_nodata).astype('B')

        retv.h_gain = 200.0
        retv.h_intercept = 0.0
        retv.h_nodata = 255

        retv.height = ((self.height - retv.h_intercept) /
                       retv.h_gain).filled(retv.h_nodata).astype('B')

        retv.p_gain = 25.0
        retv.p_intercept = 0.0
        retv.p_nodata = 255

        retv.pressure = ((self.pressure - retv.p_intercept) /
                         retv.p_gain).filled(retv.p_nodata).astype('B')

        retv.cloudiness = self.cloudiness.astype('B')
        retv.c_nodata = 255  # Is this correct? FIXME

        retv.processingflag = ctth_procflags2pps(self.processing_flags)

        return retv

    def convert2pps(self):
        """Convert the current CTTH channel to pps format.
        """
        from mpop.satin.nwcsaf_pps import CloudTopTemperatureHeight
        from nwcsaf_formats.pps_conversions import (old_ctth_press_palette_data,
                                                    old_ctth_temp_palette_data,
                                                    old_ctth_height_palette_data)

        retv = CloudTopTemperatureHeight()

        region_type = np.dtype([('area_extent', '<f8', (4,)),
                                ('xsize', '<i4'),
                                ('ysize', '<i4'),
                                ('xscale', '<f4'),
                                ('yscale', '<f4'),
                                ('lat_0', '<f4'),
                                ('lon_0', '<f4'),
                                ('lat_ts', '<f4'),
                                ('id', 'S64'),
                                ('name', 'S64'),
                                ('pcs_id', 'S128'),
                                ('pcs_def', 'S128')])

        region = np.zeros((1, ), dtype=region_type)

        region["xsize"] = self.num_of_columns
        region["ysize"] = self.num_of_lines
        region["id"] = self.region_name
        region["pcs_id"] = self.projection_name

        region["pcs_def"] = pcs_def_from_region(self.area)
        region["area_extent"] = self.area.area_extent

        retv.Region = region_type
        retv._keys.append("Region")

        namelist = np.dtype([('outval_name', 'S128')])
        retv.OutputValueNameList = namelist
        retv._keys.append("OutputValueNameList")

        retv.region = InfoObject()
        retv.region.data = region
        retv._keys.append("region")

        retv._md["time_slot"] = self.nominal_product_time
        retv._md["description"] = np.string_(
            "MSG SEVIRI Cloud Top Temperature & Height")
        retv._md["satellite_id"] = np.string_(self.satid)
        retv._md["sec_1970"] = np.uint64(0)
        retv._md["version"] = np.string_(self.product_algorithm_version)
        retv._md["orbit_number"] = np.uint64(0)
        retv._md["version"] = np.string_(self.product_algorithm_version)

        retv.processingflag_lut = []

        retv.cloudiness = InfoObject()
        retv.cloudiness.info["description"] = np.string_(
            "MSG SEVIRI effective cloudiness (%)")
        retv.cloudiness.info["gain"] = np.float32(0.0)
        retv.cloudiness.info["intercept"] = np.float32(0.0)
        retv.cloudiness.info["no_data_value"] = np.uint8(255)
        retv.cloudiness.data = self.cloudiness.astype('B')
        retv._projectables.append("cloudiness")

        retv.HEIGHT_PALETTE = InfoObject()
        retv.HEIGHT_PALETTE.data = old_ctth_height_palette_data()
        retv.HEIGHT_PALETTE.info["CLASS"] = np.string_("PALETTE")
        retv.HEIGHT_PALETTE.info["PAL_COLORMODEL"] = np.string_("RGB")
        retv.HEIGHT_PALETTE.info["PAL_TYPE"] = np.string_("STANDARD8")
        retv.HEIGHT_PALETTE.info["PAL_VERSION"] = np.string_("1.2")
        retv._keys.append("HEIGHT_PALETTE")

        retv.height = InfoObject()
        retv.height.info["description"] = np.string_(
            "MSG SEVIRI cloud top height (m)")
        retv.height.info["gain"] = np.float32(200.0)
        retv.height.info["intercept"] = np.float32(0.0)
        retv.height.info["no_data_value"] = np.uint8(255)
        retv.height.data = ((self.height - 0.0) /
                            200.0).filled(255).astype('B')
        retv.height.info["CLASS"] = np.string_("IMAGE")
        retv.height.info["IMAGE_VERSION"] = np.string_("1.2")
        retv._refs[("height", "PALETTE")] = np.string_("HEIGHT_PALETTE")
        retv._projectables.append("height")

        retv.PRESSURE_PALETTE = InfoObject()
        retv.PRESSURE_PALETTE.data = old_ctth_press_palette_data()
        retv.PRESSURE_PALETTE.info["CLASS"] = np.string_("PALETTE")
        retv.PRESSURE_PALETTE.info["PAL_COLORMODEL"] = np.string_("RGB")
        retv.PRESSURE_PALETTE.info["PAL_TYPE"] = np.string_("STANDARD8")
        retv.PRESSURE_PALETTE.info["PAL_VERSION"] = np.string_("1.2")
        retv._keys.append("PRESSURE_PALETTE")

        retv.pressure = InfoObject()
        retv.pressure.info["description"] = \
            np.string_("MSG SEVIRI cloud top pressure (hPa)")
        retv.pressure.info["gain"] = np.float32(25.0)
        retv.pressure.info["intercept"] = np.float32(0.0)
        retv.pressure.info["no_data_value"] = np.uint8(255)
        retv.pressure.data = ((self.pressure - 0.0) /
                              25.0).filled(255).astype('B')
        retv.pressure.info["CLASS"] = np.string_("IMAGE")
        retv.pressure.info["IMAGE_VERSION"] = np.string_("1.2")
        retv._refs[("pressure", "PALETTE")] = np.string_("PRESSURE_PALETTE")
        retv._projectables.append("pressure")

        retv.TEMPERATURE_PALETTE = InfoObject()
        retv.TEMPERATURE_PALETTE.data = old_ctth_temp_palette_data()
        retv.TEMPERATURE_PALETTE.info["CLASS"] = np.string_("PALETTE")
        retv.TEMPERATURE_PALETTE.info["PAL_COLORMODEL"] = np.string_("RGB")
        retv.TEMPERATURE_PALETTE.info["PAL_TYPE"] = np.string_("STANDARD8")
        retv.TEMPERATURE_PALETTE.info["PAL_VERSION"] = np.string_("1.2")
        retv._keys.append("TEMPERATURE_PALETTE")

        retv.temperature = InfoObject()
        retv.temperature.info["description"] = \
            np.string_("MSG SEVIRI cloud top temperature (K)")
        retv.temperature.info["gain"] = np.float32(1.0)
        retv.temperature.info["intercept"] = np.float32(100.0)
        retv.temperature.info["no_data_value"] = np.uint8(255)
        retv.temperature.data = ((self.temperature - 100.0) /
                                 1.0).filled(255).astype('B')
        retv.temperature.info["CLASS"] = np.string_("IMAGE")
        retv.temperature.info["IMAGE_VERSION"] = np.string_("1.2")
        retv._refs[("temperature", "PALETTE")] = np.string_(
            "TEMPERATURE_PALETTE")
        retv._projectables.append("temperature")

        retv.processing_flag = InfoObject()
        # retv.processing_flag.info["output_value_nameslist"] = processing_lut
        retv.processing_flag.info["output_value_namelist"] = np.array(processing_flags_lut,
                                                                      dtype=namelist)
        retv.processing_flag.info["description"] = np.string_(
            'MSG SEVIRI bitwise quality/processing flags')
        retv.processing_flag.data = ctth_procflags2pps(self.processing_flags)
        retv._projectables.append("processing_flag")

        retv.save = retv.write

        return retv

# ----------------------------------------


class MsgPCData(object):

    """NWCSAF/MSG Precipitating Clouds data layer
    """

    def __init__(self):
        self.data = None
        self.scaling_factor = 1
        self.offset = 0
        self.num_of_lines = 0
        self.num_of_columns = 0
        self.product = ""
        self.id = ""


class MsgPC(mpop.channel.GenericChannel):

    """NWCSAF/MSG Precipitating Clouds data structure as retrieved from HDF5
    file. Resolution sets the nominal resolution of the data.
    """

    def __init__(self):
        mpop.channel.GenericChannel.__init__(self, "PC")
        self.filled = False
        self.name = "PC"
        self.package = ""
        self.saf = ""
        self.product_name = ""
        self.num_of_columns = 0
        self.num_of_lines = 0
        self.projection_name = ""
        self.pcs_def = ""
        self.xscale = 0
        self.yscale = 0
        self.ll_lon = 0.0
        self.ll_lat = 0.0
        self.ur_lon = 0.0
        self.ur_lat = 0.0
        self.region_name = ""
        self.cfac = 0
        self.lfac = 0
        self.coff = 0
        self.loff = 0
        self.nb_param = 0
        self.gp_sc_id = 0
        self.image_acquisition_time = 0
        self.spectral_channel_id = 0
        self.nominal_product_time = 0
        self.sgs_product_quality = 0
        self.sgs_product_completeness = 0
        self.product_algorithm_version = ""
        self.probability_1 = None
        self.processing_flags = None
        self.shape = None
        self.satid = ""
        self.qc_straylight = -1

    def __str__(self):
        return ("'%s: shape %s, resolution %sm'" %
                (self.name,
                 self.probability_1.shape,
                 self.resolution))

    def is_loaded(self):
        """Tells if the channel contains loaded data.
        """
        return self.filled

# ------------------------------------------------------------------
    def read(self, filename, calibrate=True):
        """Reader for the NWCSAF/MSG precipitating clouds. Use *filename* to read data.
        """
        import h5py

        self.probability_1 = MsgPCData()
        self.processing_flags = MsgPCData()

        h5f = h5py.File(filename, 'r')
        # pylint: disable-msg=W0212
        self.package = h5f.attrs["PACKAGE"]
        self.saf = h5f.attrs["SAF"]
        self.product_name = h5f.attrs["PRODUCT_NAME"]
        self.num_of_columns = h5f.attrs["NC"]
        self.num_of_lines = h5f.attrs["NL"]
        self.projection_name = h5f.attrs["PROJECTION_NAME"]
        self.region_name = h5f.attrs["REGION_NAME"]
        self.cfac = h5f.attrs["CFAC"]
        self.lfac = h5f.attrs["LFAC"]
        self.coff = h5f.attrs["COFF"]
        self.loff = h5f.attrs["LOFF"]
        self.nb_param = h5f.attrs["NB_PARAMETERS"]
        self.gp_sc_id = h5f.attrs["GP_SC_ID"]
        self.image_acquisition_time = h5f.attrs["IMAGE_ACQUISITION_TIME"]
        self.spectral_channel_id = h5f.attrs["SPECTRAL_CHANNEL_ID"]
        self.nominal_product_time = datetime.strptime(h5f.attrs["NOMINAL_PRODUCT_TIME"],
                                                      "%Y%m%d%H%M")
        self.sgs_product_quality = h5f.attrs["SGS_PRODUCT_QUALITY"]
        self.sgs_product_completeness = h5f.attrs["SGS_PRODUCT_COMPLETENESS"]
        self.product_algorithm_version = h5f.attrs[
            "PACKAGE"] + h5f.attrs["PRODUCT_NAME"] + h5f.attrs["PRODUCT_ALGORITHM_VERSION"]

        # pylint: enable-msg=W0212
        # ------------------------

        # The precipitating clouds data
        h5d = h5f['PC_PROB1']
        self.probability_1.data = h5d[:, :]
        self.probability_1.scaling_factor = h5d.attrs["SCALING_FACTOR"]
        self.probability_1.offset = h5d.attrs["OFFSET"]
        self.probability_1.num_of_lines = h5d.attrs["N_LINES"]
        self.probability_1.num_of_columns = h5d.attrs["N_COLS"]
        self.shape = (self.probability_1.num_of_lines,
                      self.probability_1.num_of_columns)
        self.probability_1.product = h5d.attrs["PRODUCT"]
        self.probability_1.id = h5d.attrs["ID"]
        self.probability_1.data = np.ma.masked_equal(
            self.probability_1.data, 0)
        if calibrate:
            self.probability_1 = (self.probability_1.data *
                                  self.probability_1.scaling_factor +
                                  self.probability_1.offset)
        else:
            self.probability_1 = self.probability_1.data
        self.probability_1_palette = _get_palette(h5f, 'PC_PROB1')

        # ------------------------

        # The cloudtype processing/quality flags
        h5d = h5f['PC_QUALITY']
        self.processing_flags.data = h5d[:, :]
        self.processing_flags.scaling_factor = \
            h5d.attrs["SCALING_FACTOR"]
        self.processing_flags.offset = h5d.attrs["OFFSET"]
        self.processing_flags.num_of_lines = h5d.attrs["N_LINES"]
        self.processing_flags.num_of_columns = h5d.attrs["N_COLS"]
        self.processing_flags.product = h5d.attrs["PRODUCT"]
        self.processing_flags.id = h5d.attrs["ID"]
        self.processing_flags = np.ma.masked_equal(
            self.processing_flags.data, 0)

        # ------------------------
        h5f.close()

        self.area = get_area_from_file(filename)

        self.filled = True

# ------------------------------------------------------------------


def get_bit_from_flags(arr, nbit):
    """I don't know what this function does.
    """
    res = np.bitwise_and(np.right_shift(arr, nbit), 1)
    return res.astype('b')


def ctth_procflags2pps(data):
    """Convert ctth processing flags from MSG to PPS format.
    """

    ones = np.ones(data.shape, "h")

    # 2 bits to define processing status
    # (maps to pps bits 0 and 1:)
    is_bit0_set = get_bit_from_flags(data, 0)
    is_bit1_set = get_bit_from_flags(data, 1)
    proc = (is_bit0_set * np.left_shift(ones, 0) +
            is_bit1_set * np.left_shift(ones, 1))
    del is_bit0_set
    del is_bit1_set

    # Non-processed?
    # If non-processed in msg (0) then set pps bit 0 and nothing else.
    # If non-processed in msg due to FOV is cloud free (1) then do not set any
    # pps bits.
    # If processed (because cloudy) with/without result in msg (2&3) then set
    # pps bit 1.

    arr = np.where(np.equal(proc, 0), np.left_shift(ones, 0), 0)
    arr = np.where(np.equal(proc, 2), np.left_shift(ones, 1), 0)
    arr = np.where(np.equal(proc, 3), np.left_shift(ones, 1), 0)
    retv = np.array(arr)
    del proc

    # 1 bit to define if RTTOV-simulations are available?
    # (maps to pps bit 3:)
    is_bit2_set = get_bit_from_flags(data, 2)
    proc = is_bit2_set

    # RTTOV-simulations available?

    arr = np.where(np.equal(proc, 1), np.left_shift(ones, 3), 0)
    retv = np.add(retv, arr)
    del is_bit2_set

    # 3 bits to describe NWP input data
    # (maps to pps bits 4&5:)
    is_bit3_set = get_bit_from_flags(data, 3)
    is_bit4_set = get_bit_from_flags(data, 4)
    is_bit5_set = get_bit_from_flags(data, 5)
    # Put together the three bits into a nwp-flag:
    nwp_bits = (is_bit3_set * np.left_shift(ones, 0) +
                is_bit4_set * np.left_shift(ones, 1) +
                is_bit5_set * np.left_shift(ones, 2))
    arr = np.where(np.logical_and(np.greater_equal(nwp_bits, 3),
                                  np.less_equal(nwp_bits, 5)),
                   np.left_shift(ones, 4),
                   0)
    arr = np.add(arr, np.where(np.logical_or(np.equal(nwp_bits, 2),
                                             np.equal(nwp_bits, 4)),
                               np.left_shift(ones, 5),
                               0))

    retv = np.add(retv, arr)
    del is_bit3_set
    del is_bit4_set
    del is_bit5_set

    # 2 bits to describe SEVIRI input data
    # (maps to pps bits 6:)
    is_bit6_set = get_bit_from_flags(data, 6)
    is_bit7_set = get_bit_from_flags(data, 7)
    # Put together the two bits into a seviri-flag:
    seviri_bits = (is_bit6_set * np.left_shift(ones, 0) +
                   is_bit7_set * np.left_shift(ones, 1))
    arr = np.where(np.greater_equal(seviri_bits, 2),
                   np.left_shift(ones, 6), 0)

    retv = np.add(retv, arr)
    del is_bit6_set
    del is_bit7_set

    # 4 bits to describe which method has been used
    # (maps to pps bits 7&8 and bit 2:)
    is_bit8_set = get_bit_from_flags(data, 8)
    is_bit9_set = get_bit_from_flags(data, 9)
    is_bit10_set = get_bit_from_flags(data, 10)
    is_bit11_set = get_bit_from_flags(data, 11)
    # Put together the four bits into a method-flag:
    method_bits = (is_bit8_set * np.left_shift(ones, 0) +
                   is_bit9_set * np.left_shift(ones, 1) +
                   is_bit10_set * np.left_shift(ones, 2) +
                   is_bit11_set * np.left_shift(ones, 3))
    arr = np.where(np.logical_or(
        np.logical_and(np.greater_equal(method_bits, 1),
                       np.less_equal(method_bits, 2)),
        np.equal(method_bits, 13)),
        np.left_shift(ones, 2),
        0)
    arr = np.add(arr,
                 np.where(np.equal(method_bits, 1),
                          np.left_shift(ones, 7),
                          0))
    arr = np.add(arr,
                 np.where(np.logical_and(
                     np.greater_equal(method_bits, 3),
                     np.less_equal(method_bits, 12)),
                     np.left_shift(ones, 8),
                     0))

    # (Maps directly - as well - to the spare bits 9-12)
    arr = np.add(arr, np.where(is_bit8_set, np.left_shift(ones, 9), 0))
    arr = np.add(arr, np.where(is_bit9_set,
                               np.left_shift(ones, 10),
                               0))
    arr = np.add(arr, np.where(is_bit10_set,
                               np.left_shift(ones, 11),
                               0))
    arr = np.add(arr, np.where(is_bit11_set,
                               np.left_shift(ones, 12),
                               0))
    retv = np.add(retv, arr)
    del is_bit8_set
    del is_bit9_set
    del is_bit10_set
    del is_bit11_set

    # 2 bits to describe the quality of the processing itself
    # (maps to pps bits 14&15:)
    is_bit12_set = get_bit_from_flags(data, 12)
    is_bit13_set = get_bit_from_flags(data, 13)
    # Put together the two bits into a quality-flag:
    qual_bits = (is_bit12_set * np.left_shift(ones, 0) +
                 is_bit13_set * np.left_shift(ones, 1))
    arr = np.where(np.logical_and(np.greater_equal(qual_bits, 1),
                                  np.less_equal(qual_bits, 2)),
                   np.left_shift(ones, 14), 0)
    arr = np.add(arr,
                 np.where(np.equal(qual_bits, 2),
                          np.left_shift(ones, 15),
                          0))

    retv = np.add(retv, arr)
    del is_bit12_set
    del is_bit13_set

    return retv.astype('h')


def ctype_procflags2pps(data):
    """Converting cloud type processing flags to
    the PPS format, in order to have consistency between
    PPS and MSG cloud type contents.
    """

    ones = np.ones(data.shape, "h")

    # msg illumination bit 0,1,2 (undefined,night,twilight,day,sunglint) maps
    # to pps bits 2, 3 and 4:
    is_bit0_set = get_bit_from_flags(data, 0)
    is_bit1_set = get_bit_from_flags(data, 1)
    is_bit2_set = get_bit_from_flags(data, 2)
    illum = is_bit0_set * np.left_shift(ones, 0) + \
        is_bit1_set * np.left_shift(ones, 1) + \
        is_bit2_set * np.left_shift(ones, 2)
    del is_bit0_set
    del is_bit1_set
    del is_bit2_set
    # Night?
    # If night in msg then set pps night bit and nothing else.
    # If twilight in msg then set pps twilight bit and nothing else.
    # If day in msg then unset both the pps night and twilight bits.
    # If sunglint in msg unset both the pps night and twilight bits and set the
    # pps sunglint bit.
    arr = np.where(np.equal(illum, 1), np.left_shift(ones, 2), 0)
    arr = np.where(np.equal(illum, 2), np.left_shift(ones, 3), arr)
    arr = np.where(np.equal(illum, 3), 0, arr)
    arr = np.where(np.equal(illum, 4), np.left_shift(ones, 4), arr)
    retv = np.array(arr)
    del illum

    # msg nwp-input bit 3 (nwp present?) maps to pps bit 7:
    # msg nwp-input bit 4 (low level inversion?) maps to pps bit 6:
    is_bit3_set = get_bit_from_flags(data, 3)
    is_bit4_set = get_bit_from_flags(data, 4)
    nwp = (is_bit3_set * np.left_shift(ones, 0) +
           is_bit4_set * np.left_shift(ones, 1))
    del is_bit3_set
    del is_bit4_set

    arr = np.where(np.equal(nwp, 1), np.left_shift(ones, 7), 0)
    arr = np.where(np.equal(nwp, 2), np.left_shift(ones, 7) +
                   np.left_shift(ones, 6), arr)
    arr = np.where(np.equal(nwp, 3), 0, arr)
    retv = np.add(arr, retv)
    del nwp

    # msg seviri-input bits 5&6 maps to pps bit 8:
    is_bit5_set = get_bit_from_flags(data, 5)
    is_bit6_set = get_bit_from_flags(data, 6)
    seviri = (is_bit5_set * np.left_shift(ones, 0) +
              is_bit6_set * np.left_shift(ones, 1))
    del is_bit5_set
    del is_bit6_set

    retv = np.add(retv,
                  np.where(np.logical_or(np.equal(seviri, 2),
                                         np.equal(seviri, 3)),
                           np.left_shift(ones, 8), 0))
    del seviri

    # msg quality bits 7&8 maps to pps bit 9&10:
    is_bit7_set = get_bit_from_flags(data, 7)
    is_bit8_set = get_bit_from_flags(data, 8)
    quality = (is_bit7_set * np.left_shift(ones, 0) +
               is_bit8_set * np.left_shift(ones, 1))
    del is_bit7_set
    del is_bit8_set

    arr = np.where(np.equal(quality, 2), np.left_shift(ones, 9), 0)
    arr = np.where(np.equal(quality, 3), np.left_shift(ones, 10), arr)
    retv = np.add(arr, retv)
    del quality

    # msg bit 9 (stratiform-cumuliform distinction?) maps to pps bit 11:
    is_bit9_set = get_bit_from_flags(data, 9)
    retv = np.add(retv,
                  np.where(is_bit9_set,
                           np.left_shift(ones, 11),
                           0))
    del is_bit9_set

    return retv.astype('h')


# def pps_luts():
#     """Gets the LUTs for the PPS Cloud Type data fields.
#     Returns a tuple with Cloud Type lut, Cloud Phase lut, Processing flags lut
#     """

#     return ctype_lut, phase_lut, quality_lut


def add_nullterm_attr(obj, key, val):
    import h5py
    atype = h5py.h5t.C_S1
    tid = h5py.h5t.TypeID.copy(h5py.h5t.C_S1)
    tid.set_size(len(val) + 1)
    tid.set_strpad(h5py.h5t.STR_NULLTERM)
    sid = h5py.h5s.create(h5py.h5s.SCALAR)
    aid = h5py.h5a.create(obj.id, key, tid, sid)
    blob = np.string_(val)
    aid.write(np.array(val))


class NordRadCType(object):

    """Wrapper aroud the msg_ctype channel.
    """

    def __init__(self, ctype_instance):
        self.ctype = ctype_instance
        self.datestr = ctype_instance.image_acquisition_time
        self.info = {}

    def save(self, filename, **kwargs):
        """Save the current instance to nordrad hdf format.
        """
        #import _pyhl
        import h5py

        msgctype = self.ctype

        h5f = h5py.File(filename, mode="w")

        # What
        what = h5f.create_group("what")
        #what.attrs["object"] = np.string_("IMAGE")
        add_nullterm_attr(what, "object", "IMAGE")
        what.attrs["sets"] = np.int(1)
        #what.attrs["version"] = np.string_("H5rad 1.2")
        add_nullterm_attr(what, "version", "H5rad 1.2")

        yyyymmdd = self.datestr[0:8]
        hourminsec = self.datestr[8:12] + '00'
        #what.attrs["date"] = np.string_(yyyymmdd)
        add_nullterm_attr(what, "date", yyyymmdd)
        #what.attrs["time"] = np.string_(hourminsec)
        add_nullterm_attr(what, "time", hourminsec)

        # Where
        where = h5f.create_group("where")
        #where.attrs["projdef"] = np.string_(msgctype.area.proj4_string)
        add_nullterm_attr(where, "projdef", msgctype.area.proj4_string)
        where.attrs["xsize"] = np.int(msgctype.num_of_columns)
        where.attrs["ysize"] = np.int(msgctype.num_of_lines)
        where.attrs["xscale"] = np.float(msgctype.xscale)
        where.attrs["yscale"] = np.float(msgctype.yscale)
        where.attrs["LL_lon"] = np.float(msgctype.ll_lon)
        where.attrs["LL_lat"] = np.float(msgctype.ll_lat)
        where.attrs["UR_lon"] = np.float(msgctype.ur_lon)
        where.attrs["UR_lat"] = np.float(msgctype.ur_lat)

        # How
        how = h5f.create_group("how")
        #how.attrs["area"] = np.string_(msgctype.region_name)
        add_nullterm_attr(how, "area", msgctype.region_name)

        # image1
        image1 = h5f.create_group("image1")
        image1.create_dataset("data", data=msgctype.cloudtype.astype('B'),
                              compression="gzip", compression_opts=COMPRESS_LVL)

        what = image1.create_group("what")
        #what.attrs["product"] = np.string_('MSGCT')
        add_nullterm_attr(what, "product", 'MSGCT')
        what.attrs["prodpar"] = np.float(0.0)
        #what.attrs["quantity"] = np.string_("ct")
        add_nullterm_attr(what, "quantity", "ct")
        #what.attrs["startdate"] = np.string_(yyyymmdd)
        add_nullterm_attr(what, "startdate", yyyymmdd)
        #what.attrs["starttime"] = np.string_(hourminsec)
        add_nullterm_attr(what, "starttime", hourminsec)
        #what.attrs["enddate"] = np.string_(yyyymmdd)
        add_nullterm_attr(what, "enddate", yyyymmdd)
        #what.attrs["endtime"] = np.string_(hourminsec)
        add_nullterm_attr(what, "endtime", hourminsec)
        what.attrs["gain"] = np.float(1.0)
        what.attrs["offset"] = np.float(0.0)
        what.attrs["nodata"] = np.float(0.0)
        # What we call missingdata in PPS:
        what.attrs["undetect"] = np.float(20.0)

        h5f.close()


MSG_PGE_EXTENTIONS = ["PLAX.CTTH.0.h5", "PLAX.CLIM.0.h5", "h5"]


def get_best_product(filename, area_extent):
    """Get the best of the available products for the *filename* template.
    """

    basename = filename
    for ext in MSG_PGE_EXTENTIONS:
        if filename.endswith(ext):
            basename = filename[:-(len(ext) + 1)]
            break
    for ext in MSG_PGE_EXTENTIONS:
        if not basename.endswith(ext):
            match_str = basename + "." + ext
        else:
            match_str = basename
        LOG.debug("glob-string for filename: " + str(match_str))
        flist = glob.glob(match_str)
        if len(flist) == 0:
            LOG.warning("No matching .%s input MSG file."
                        % ext)
        else:
            # File found:
            if area_extent is None:
                LOG.warning("Didn't specify an area, taking " + flist[0])
                return flist[0]
            for fname in flist:
                aex = get_area_extent(fname)
                if np.all(np.max(np.abs(np.array(aex) -
                                        np.array(area_extent))) < 1000):
                    LOG.info("MSG file found: %s" % fname)
                    return fname
            LOG.info("Did not find any MSG file for specified area")


def get_best_products(filename, area_extent):
    """Get the best of the available products for the *filename* template.
    """

    basename = filename
    for ext in MSG_PGE_EXTENTIONS:
        if filename.endswith(ext):
            basename = filename[:-(len(ext) + 1)]
            break

    filenames = []

    for ext in MSG_PGE_EXTENTIONS:
        if not basename.endswith(ext):
            match_str = basename + "." + ext
        else:
            match_str = basename

        LOG.debug('Match string = ' + str(match_str))
        flist = glob.glob(match_str)
        if len(flist) == 0:
            LOG.warning("No matching .%s input MSG file."
                        % ext)
        else:
            # File found:
            if area_extent is None:
                LOG.warning("Didn't specify an area, taking " + flist[0])
                filenames.append(flist[0])
            else:
                found = False
                for fname in flist:
                    aex = get_area_extent(fname)
                    if np.all(np.max(np.abs(np.array(aex) -
                                            np.array(area_extent))) < 1000):
                        found = True
                        LOG.info("MSG file found: %s" % fname)
                        filenames.append(fname)
                    if not found:
                        LOG.info(
                            "Did not find any MSG file for specified area")
    LOG.debug("Sorted filenames: %s", str(sorted(filenames)))
    return sorted(filenames)


def get_area_from_file(filename):
    """Get the area from the h5 file.
    """
    from pyresample.geometry import AreaDefinition
    import h5py

    aex = get_area_extent(filename)
    h5f = h5py.File(filename, 'r')
    pname = h5f.attrs["PROJECTION_NAME"]
    proj = {}
    if pname.startswith("GEOS"):
        proj["proj"] = "geos"
        proj["a"] = "6378169.0"
        proj["b"] = "6356583.8"
        proj["h"] = "35785831.0"
        proj["lon_0"] = str(float(pname.split("<")[1][:-1]))
    else:
        raise NotImplementedError("Only geos projection supported yet.")
    area_def = AreaDefinition(h5f.attrs["REGION_NAME"],
                              h5f.attrs["REGION_NAME"],
                              pname,
                              proj,
                              int(h5f.attrs["NC"]),
                              int(h5f.attrs["NL"]),
                              aex)
    h5f.close()
    return area_def


def load(scene, **kwargs):
    """Load data into the *channels*. *Channels* is a list or a tuple
    containing channels we will load data into. If None, all channels are
    loaded.
    """

    area_extent = kwargs.get("area_extent")
    conf = ConfigParser.ConfigParser()
    conf.read(os.path.join(CONFIG_PATH, scene.fullname + ".cfg"))

    if kwargs.get("filename") is not None:
        full_filename = kwargs["filename"]
        LOG.debug("File name determined from input: " + str(full_filename))
    else:
        directory = conf.get(scene.instrument_name + "-level3",
                             "dir",
                             raw=True)
        filename = conf.get(scene.instrument_name + "-level3", "filename",
                            raw=True)
        pathname = os.path.join(directory, filename)
        full_filename = None

    for chan in scene.channels_to_load:
        LOG.debug("Channel to load = " + str(chan))
        if chan == "CTTH":
            product = 'CTTH_'
            pname = 'CTTH'
            number = '03'
        elif chan == "CloudType":
            product = 'CT___'
            pname = 'CloudType'
            number = '02'
        elif chan == "CloudType_plax":
            product = 'CT___'
            pname = 'CloudType_plax'
            number = '02'
        else:
            LOG.info('Channel ' + str(chan) + ' not known by this reader!')
            continue

        if not full_filename:
            filename = (scene.time_slot.strftime(pathname)
                        % {"number": number,
                           "product": product})
        else:
            filename = full_filename

        ct_chan = CASES[pname]()
        if pname == 'CTTH':
            product = get_best_product(filename, area_extent)
        elif pname == 'CloudType':
            product = get_best_products(filename, area_extent)[-1]
        elif pname == 'CloudType_plax':
            product = get_best_products(filename, area_extent)[0]
            LOG.debug("Parallax corrected file: %s", product)

        if not product or not os.path.exists(product):
            raise IOError("Failed to locate file...")

        ct_chan.read(product)
        ct_chan.name = pname
        ct_chan.satid = (scene.fullname.capitalize())
        ct_chan.resolution = ct_chan.area.pixel_size_x
        scene.channels.append(ct_chan)

        LOG.info("Channel " + str(chan) + " loaded")

    LOG.info("Loading channels done.")


CASES = {
    'CTTH': MsgCTTH,
    'CloudType': MsgCloudType,
    'CloudType_plax': MsgCloudType
}

if __name__ == '__main__':

    FILENAME = "/data/proj/safutv/geo_out/0deg/SAFNWC_MSG3_CT___201505260615_MSG-N_______.h5"

    ct = MsgCloudType()
    ct.read(FILENAME)

    ct.convert2pps().save("blact.h5")

    FILENAME = "/data/proj/safutv/geo_out/0deg/SAFNWC_MSG3_CTTH_201505260615_MSG-N_______.h5"
    FILENAME = "/data/24/saf/geo_out/0deg/SAFNWC_MSG3_CTTH_201505260615_EuropeCanary.h5"
    ct = MsgCTTH()
    ct.read(FILENAME)

    ct.convert2pps().save("blactth.h5")
