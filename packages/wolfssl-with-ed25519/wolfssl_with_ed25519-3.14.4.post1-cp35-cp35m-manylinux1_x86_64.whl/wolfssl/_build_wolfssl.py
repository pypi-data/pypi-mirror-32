#!/usr/bin/env python
# -*- coding: utf-8 -*-
#
# Copyright (C) 2006-2017 wolfSSL Inc.
#
# This file is part of wolfSSL. (formerly known as CyaSSL)
#
# wolfSSL is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; either version 2 of the License, or
# (at your option) any later version.
#
# wolfSSL is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA

import glob
import os
import subprocess
from contextlib import contextmanager
from distutils.util import get_platform

FILE = os.path.abspath(__file__)

def local_path(path):
    """ Return path relative to the root of this project
    """
    current = os.path.dirname(FILE)
    gparent = os.path.dirname(os.path.dirname(current))
    return os.path.abspath(os.path.join(gparent, path))

WOLFSSL_GIT_ADDR = "https://github.com/wolfssl/wolfssl.git"
WOLFSSL_SRC_PATH = local_path("lib/wolfssl/src")


def call(cmd):
    print("Calling: '{}' from working directory {}".format(cmd, os.getcwd()))

    old_env = os.environ["PATH"]
    os.environ["PATH"] = "{}:{}".format(WOLFSSL_SRC_PATH, old_env)
    subprocess.check_call(cmd, shell=True, env=os.environ)
    os.environ["PATH"] = old_env


@contextmanager
def chdir(new_path, mkdir=False):
    old_path = os.getcwd()

    if mkdir:
        try:
            os.mkdir(new_path)
        except OSError:
            pass

    try:
        yield os.chdir(new_path)
    finally:
        os.chdir(old_path)


def clone_wolfssl(version):
    """ Clone wolfSSL C library repository
    """
    call("git clone --depth=1 --branch={} {} {}".format(
        version, WOLFSSL_GIT_ADDR, WOLFSSL_SRC_PATH))
def checkout_version(version):
    """ Ensure that we have the right version
    """
    with chdir(WOLFSSL_SRC_PATH):
        current = subprocess.check_output(
            ["git", "describe", "--all", "--exact-match"]
        ).strip().decode().split('/')[-1]

        if current != version:
            tags = subprocess.check_output(
                ["git", "tag"]
            ).strip().decode().split("\n")

            if version != "master" and version not in tags:
                call("git fetch --depth=1 origin tag {}".format(version))

            call("git checkout --force {}".format(version))

            return True  # rebuild needed

    return False

def apply_patches():
    """ Apply patches to the source
    """
    patch_dir = local_path("patch")
    patches = glob.glob("{}/*.patch".format(patch_dir))
    with chdir(WOLFSSL_SRC_PATH):
        for patch in patches:
            print("Apply patch", patch)
            call("git apply {}".format(patch))

def ensure_wolfssl_src(version):
    """ Ensure that wolfssl sources are presents and up-to-date
    """
    if not os.path.isdir(WOLFSSL_SRC_PATH):
        clone_wolfssl(version)
        return True

    return checkout_version(version)


def make_flags(prefix):
    """ Returns compilation flags
    """
    flags = []

    if get_platform() in ["linux-x86_64", "linux-i686"]:
        flags.append("CFLAGS=-fpic")

    # install location
    flags.append("--prefix={}".format(prefix))

    # lib only
    flags.append("--enable-ed25519")
    flags.append("--enable-curve25519")

    flags.append("--disable-shared")
    flags.append("--disable-examples")

    # tls 1.3
    flags.append("--enable-tls13")

    return " ".join(flags)


def make(configure_flags):
    """ Create a release of wolfSSL C library
    """
    with chdir(WOLFSSL_SRC_PATH):
        call("git clean -fd")
        call("git checkout -- .")
        apply_patches()

        try:
            call("./autogen.sh")
        except subprocess.CalledProcessError:
            call("libtoolize")
            call("./autogen.sh")

        call("./configure {}".format(configure_flags))
        call("make")
        call("make install-exec")


def build_wolfssl(version="master"):
    prefix = local_path("lib/wolfssl/{}/{}".format(
        get_platform(), version))
    libfile = os.path.join(prefix, 'lib/libwolfssl.la')

    rebuild = ensure_wolfssl_src(version)

    if rebuild or not os.path.isfile(libfile):
        make(make_flags(prefix))


if __name__ == "__main__":
    build_wolfssl()
