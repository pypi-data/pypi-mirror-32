# coding: utf-8
from __future__ import unicode_literals

import logging.config
import sys
from os import path


def init_log(log_dir=None):
    if log_dir and not path.exists(log_dir):
        msg = '指定路径不存在:%s' % log_dir
        sys.stdout(msg)
        log_dir = None

    config = {
        'version': 1,
        'disable_existing_loggers': False,
        'formatters': {
            'default': {
                'format': '%(levelname)s %(asctime)s %(module)s:%(funcName)s:%(lineno)d %(message)s'
            },
            'simple': {
                'format': '%(level)s %(message)s'
            }
        },
        'handlers': {
            'console': {
                'level': 'DEBUG',
                'class': 'logging.StreamHandler',
                'formatter': 'default',
            },
        },
        'loggers': {
            '': {
                'handlers': ['console'],
                'level': 'INFO',
                'propagate': True,
            },
        }
    }

    if log_dir:
        config['handles']['file'] = {
            'level': 'DEBUG',
            'class': 'logging.handlers.RotatingFileHandler',
            'filename': path.join(log_dir, 'essay.log'),
            'maxBytes': 1024 * 1024 * 50,
            'backupCount': 5,
            'formatter': 'default',
        }
        config['loggers']['handlers'] = ['console', 'file']

    logging.config.dictConfig(config)
