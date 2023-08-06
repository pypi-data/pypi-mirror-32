# -*- coding: utf-8 -*-
#
# goespy/__init__.py
#                       BSD 2-Clause License
#
#           Copyright (c) 2018, Paulo Alexandre da Silva Mello
#                       All rights reserved.
#
#        Redistribution and use in source and binary forms, with or without
#      modification, are permitted provided that the following conditions are met:
#
#*      Redistributions of source code must retain the above copyright notice, this
#               list of conditions and the following disclaimer.
#
#*      Redistributions in binary form must reproduce the above copyright notice,
#       this list of conditions and the following disclaimer in the documentation
#               and/or other materials provided with the distribution.
#
#           THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
#           AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
#           IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
#           DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
#           FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
#           DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
#           SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
#           CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
#           OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
#           OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.


from goespy import checkData
from goespy import Downloader
from goespy import utils
import boto3
import botocore
import datetime
import os
import glob
import sys
import threading

__author__ = "Paulo Alexandre S. Mello"
__email__ = "palexandremello@gmail.com"
__date__ = "2018-05-22"
__version__ = '0.2.1'

utils.bannerDisplay(__version__)