#!/usr/bin/env python
"""
"""
import os,platform,re,sys

extra=dict()
if sys.version_info >= (3,):
    extra['use_2to3'] = True

from distutils.core import setup,Extension
try:
   from distutils.command.build_py import build_py_2to3 as build_py #for Python3
except ImportError:
   from distutils.command.build_py import build_py     # for Python2

rev="$Revision: 0.1dev23$"
sysname=platform.system()
hgtag = "$HGTag: dev.2-deae70e790d3 $"

setup(name="PyUSBTMC",
      version=rev[11:-1],
      author="Noboru Yamamoto, KEK, JAPAN",
      author_email = "Noboru.YAMAMOTO@kek.jp",
      description ="Python module to control USBTMC/USB488 from python",
      long_description=""" Python module to control USBTMC/USB488 from python.
It requires pyusb module and libusb (or openusb) library to run. 
Although it is still in the development stage, it can read/write data from/to the devices.""",
      platforms="tested on MacOSX10.7",
      url="http://www-acc.kek.jp/EPICS_Gr/products/",
      py_modules=["PyUSBTMC","lsusb","usbids","samples.test_DPO", "samples.test_dso_anim"],
      data_files=[("share/misc",['usb.ids'])]
      )

