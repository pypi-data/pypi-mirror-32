#!/usr/bin/env python
# -*- coding: utf-8 -*-

MIN_SETUPTOOLS_VERSION = "31.0.0"
import setuptools
from distutils.version import LooseVersion
assert (LooseVersion(setuptools.__version__) >= LooseVersion(MIN_SETUPTOOLS_VERSION)), "LIEF requires a setuptools version '{}' or higher (pip install setuptools --upgrade)".format(MIN_SETUPTOOLS_VERSION)

from setuptools import setup
from setuptools.command.build_ext import build_ext
from distutils.dir_util import remove_tree, mkpath
from setuptools.command.bdist_egg import bdist_egg
from setuptools.command.bdist_egg import log as bdist_egg_log

from setuptools.command.sdist import sdist
from setuptools.command.sdist import log as sdist_log

from setuptools.extension import Extension

import re
import os
import platform
import shutil
import sys
import struct
import zipfile

try:
    from urllib.request import urlopen, Request
except:
    from urllib2 import urlopen, Request


try:
    from io import BytesIO
except:
    try:
        from cStringIO import StringIO as BytesIO
    except:
        from StringIO import StringIO as BytesIO

package_dir       = os.path.dirname(os.path.realpath(__file__))
pkg_info          = os.path.join(package_dir, "PKG-INFO")
in_source_package = os.path.isfile(pkg_info) # (e.g. pip)

is_branch = False

libpylief = {
        'Windows': "_pylief.pyd",
        'Darwin': "_pylief.so",
        'Linux': "_pylief.so",
        'FreeBSD': "_pylief.so",
        }
version_re = r"Version:\s+(?P<major>\d+)\.(?P<minor>\d+)\.(?P<patch>\d+)\.?(.*)"
if in_source_package:
    with open(pkg_info, "r") as f:
        major, minor, patch, branch = re.findall(version_re, f.read(), re.MULTILINE)[0]
        lief_version = "{:d}.{:d}.{:d}".format(int(major), int(minor), int(patch))
        is_branch = len(branch.strip()) > 0
else:
    lief_version = "0.9.0"

package_description = open(os.path.join(package_dir, "README")).read()


def get_lief_platform_name():
    system = platform.system()
    arch   = struct.calcsize('P') * 8

    if system == 'Windows':
        return "windows_x64" if arch == 64 else "windows_x32"
    elif system == 'Darwin':
        return "osx" if arch == 64 else "osx_x32"
    elif system == 'Linux':
        return "linux" if arch == 64 else "linux_x32"
    elif system == 'FreeBSD':
        return "freebsd" if arch == 64 else "freebsd_x32"



class lief_sdist(sdist):

    user_options = sdist.user_options + [
            ('dev', None, "Add a dev marker")
            ]

    def initialize_options(self):
        sdist.initialize_options(self)
        self.dev = 0

    def run(self):
        if self.dev:
            suffix = '.dev-{:s}'.format("a448c5e")
            self.distribution.metadata.version += suffix
        sdist.run(self)

    def make_distribution(self):
        sdist.make_distribution(self)
        base_dir = self.distribution.get_fullname()
        base_name = os.path.join(self.dist_dir, base_dir)
        for fmt in self.formats:
            if fmt == 'gztar':
                fmt = 'tar.gz'

            if fmt == 'bztar':
                fmt = 'tar.bz2'

            if fmt == 'ztar':
                fmt = 'tar.Z'

            new_name = "py{name}-{version}.{ext}".format(
                        name=self.distribution.get_name(),
                        version=lief_version,
                        ext=fmt)

            if self.dev:
                new_name = "py{name}-{version}.dev.{ext}".format(
                        name=self.distribution.get_name(),
                        version=lief_version,
                        ext=fmt)
            new_name = os.path.join(self.dist_dir, new_name)
            shutil.move(base_name + "." + fmt, new_name)




class lief_bdist_egg(bdist_egg):
    def initialize_options(self):
        bdist_egg.initialize_options(self)

        if not in_source_package:
            self.plat_name = get_lief_platform_name()


    def run(self):
        if in_source_package:
            self._build_from_source_package()
        else:
            bdist_egg.run(self)

    def _build_from_source_package(self):
        python_version       = sys.version_info
        python_major_version = python_version[0]
        os_version           = get_lief_platform_name()

        url_branch_fmt  = "https://github.com/lief-project/packages/raw/lief-{branch}-latest/lief-{version}-py{pyversion}-{platform}.{ext}"
        url_release_fmt = "https://github.com/lief-project/LIEF/releases/download/{version}/lief-{version}-py{pyversion}-{platform}.{ext}"
        url_userpath    = "~/lief-{version}-py{pyversion}-{platform}.{ext}"

        url = ""
        if is_branch:
            url = url_branch_fmt.format(
                    branch='master',
                    platform=os_version,
                    version=lief_version,
                    pyversion="{}.{}".format(python_version[0], python_version[1]),
                    ext="egg")
        else:
            url = url_release_fmt.format(
                    platform=os_version,
                    version=lief_version,
                    pyversion="{}.{}".format(python_version[0], python_version[1]),
                    ext="egg")

        bdist_egg_log.info("Url: {}".format(url))
        egg_data = None
        network_error = None
        try:
            egg_data = urlopen(url).read()
        except Exception as e:
            network_error = e


        if network_error is not None:
            bdist_egg_log.warn(network_error)
            url = url_userpath.format(
                    platform=os_version,
                    version=lief_version,
                    pyversion="{}.{}".format(python_version[0], python_version[1]),
                    ext="egg")
            url = os.path.expanduser(url)
            if os.path.isfile(url):
                with open(url, 'rb') as f:
                    egg_data = f.read()
            else:
                raise Exception("Unable to find {}".format(url))


        egg_file = BytesIO(egg_data)

        mkpath(os.path.dirname(self.egg_output))

        bdist_egg_log.info("Output: {}".format(self.egg_output))
        with open(self.egg_output, 'wb') as f:
            f.write(egg_file.getvalue())

        getattr(self.distribution, 'dist_files', []).append(
        ('bdist_egg', "{}.{}".format(python_version[0], python_version[1]), self.egg_output))



class lief_build_ext(build_ext):


    def build_extension(self, ext):
        self.target = self.get_ext_fullpath(ext.name)
        target_dir  = os.path.dirname(self.target)

        try:
            os.makedirs(target_dir)
        except:
            pass

        if in_source_package:
            self._install_from_source_package()
        else:
            shutil.copyfile(libpylief[platform.system()], self.target)


    def _install_from_source_package(self):
        python_version       = sys.version_info
        python_major_version = python_version[0]
        os_version           = get_lief_platform_name()
        target_extension     = os.path.splitext(self.target)[1]

        url_branch_fmt  = "https://github.com/lief-project/packages/raw/lief-{branch}-latest/lief-{version}-py{pyversion}-{platform}.{ext}"
        url_release_fmt = "https://github.com/lief-project/LIEF/releases/download/{version}/lief-{version}-py{pyversion}-{platform}.{ext}"
        url_userpath    = "~/lief-{version}-py{pyversion}-{platform}.{ext}"

        url = ""
        if is_branch:
            url = url_branch_fmt.format(
                    branch='master',
                    platform=os_version,
                    version=lief_version,
                    pyversion="{}.{}".format(python_version[0], python_version[1]),
                    ext="egg")
        else:
            url = url_release_fmt.format(
                    platform=os_version,
                    version=lief_version,
                    pyversion="{}.{}".format(python_version[0], python_version[1]),
                    ext="egg")

        bdist_egg_log.info("Url: {}".format(url))
        egg_data = None
        network_error = None
        try:
            egg_data = urlopen(url).read()
        except Exception as e:
            network_error = e

        if network_error is not None:
            bdist_egg_log.warn(network_error)
            url = url_userpath.format(
                    platform=os_version,
                    version=lief_version,
                    pyversion="{}.{}".format(python_version[0], python_version[1]),
                    ext="egg")
            url = os.path.expanduser(url)
            if os.path.isfile(url):
                with open(url, 'rb') as f:
                    egg_data = f.read()
            else:
                raise Exception("Unable to find {}".format(url))





        egg_file = BytesIO(egg_data)

        egg_zip = zipfile.ZipFile(egg_file)
        extension_member = [info for info in egg_zip.infolist() if info.filename.endswith(target_extension)][0]
        extension_data = egg_zip.read(extension_member)
        with open(self.target, 'wb') as f:
            f.write(extension_data)

setup(
    version              = lief_version,
    ext_modules          = [Extension('_pylief', [])],
    cmdclass={
        'build_ext': lief_build_ext,
        'bdist_egg': lief_bdist_egg,
        'sdist':     lief_sdist
    },
)
