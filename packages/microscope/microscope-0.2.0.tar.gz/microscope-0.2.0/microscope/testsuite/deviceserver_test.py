#!/usr/bin/env python
# -*- coding: utf-8 -*-

## Copyright (C) 2017 David Pinto <david.pinto@bioch.ox.ac.uk>
##
## This file is part of Microscope.
##
## Microscope is free software: you can redistribute it and/or modify
## it under the terms of the GNU General Public License as published by
## the Free Software Foundation, either version 3 of the License, or
## (at your option) any later version.
##
## Microscope is distributed in the hope that it will be useful,
## but WITHOUT ANY WARRANTY; without even the implied warranty of
## MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
## GNU General Public License for more details.
##
## You should have received a copy of the GNU General Public License
## along with Microscope.  If not, see <http://www.gnu.org/licenses/>.

import multiprocessing
import time
import unittest

import microscope.deviceserver

from microscope.devices import device
from microscope.testsuite.devices import TestCamera
from microscope.testsuite.devices import TestFilterWheel

class BaseTestServeDevices(unittest.TestCase):
  """Handles start and termination of deviceserver.

  Subclasses may overload class properties defaults as needed.

  Attributes:
    DEVICES (list): list of :class:`microscope.devices` to initialise.
    TIMEOUT (number): time given for service to terminate after
      receiving signal to terminate.
    p (multiprocessing.Process): device server process.
  """
  DEVICES = []
  TIMEOUT = 2
  def setUp(self):
    init = microscope.deviceserver.serve_devices
    self.p = multiprocessing.Process(target=init, args=(self.DEVICES,))
    self.p.start()

  def tearDown(self):
    self.p.terminate()
    self.p.join(self.TIMEOUT)
    self.assertFalse(self.p.is_alive(),
                     "deviceserver not dead after SIGTERM")

class TestStarting(BaseTestServeDevices):
  DEVICES = [
    device(TestCamera, '127.0.0.1', 8001, otherargs=1,),
    device(TestFilterWheel, '127.0.0.1', 8003,
           filters=[(0, 'GFP', 525), (1, 'RFP'), (2, 'Cy5')]),
   ]

  def test_standard(self):
    """Simplest case, start and exit, given enough time to start all devices"""
    time.sleep(2)
    self.assertTrue(self.p.is_alive(), "service dies at start")

  def test_immediate_interrupt(self):
    """Check issues on SIGTERM before starting all devices"""
    pass

class TestInputCheck(BaseTestServeDevices):
  def test_empty_devices(self):
    """Check behaviour if there are no devices."""
    time.sleep(2)
    self.assertTrue(not self.p.is_alive(),
                    "not dying for empty list of devices")

if __name__ == '__main__':
  unittest.main()
