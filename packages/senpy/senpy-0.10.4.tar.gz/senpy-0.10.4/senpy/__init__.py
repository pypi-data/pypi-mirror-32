#!/usr/bin/python
# -*- coding: utf-8 -*-
#    Copyright 2014 J. Fernando Sánchez Rada - Grupo de Sistemas Inteligentes
#                                                       DIT, UPM
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS,
#    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#    See the License for the specific language governing permissions and
#    limitations under the License.
"""
Sentiment analysis server in Python
"""
from .version import __version__

import logging

logger = logging.getLogger(__name__)

logger.info('Using senpy version: {}'.format(__version__))

from .utils import easy, easy_load, easy_test  # noqa: F401

from .models import *  # noqa: F401,F403
from .plugins import *  # noqa: F401,F403
from .extensions import *  # noqa: F401,F403

__all__ = ['api', 'blueprints', 'cli', 'extensions', 'models', 'plugins']
