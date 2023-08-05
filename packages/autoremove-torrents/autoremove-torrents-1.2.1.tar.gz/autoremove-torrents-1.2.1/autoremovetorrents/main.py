#!/usr/bin/python3
# -*- coding:utf-8 -*-
import sys
import getopt
import traceback
import yaml
from . import logger
from .task import Task

def main():
    # View Mode
    view_mode = False
    # The path of the configuration file
    conf_path = 'config.yml'
    # Task
    task = None
    # Logger
    lg = logger.register(__name__)

    # Get arguments
    try:
        opts = getopt.getopt(sys.argv[1:], 'vc:t:', ['view', 'conf=', 'task='])[0]
    except getopt.GetoptError:
        print('Invalid arguments.')
        sys.exit(255)
    for opt,arg in opts:
        if opt in ('-v', '--view'): # View mode (without deleting)
            view_mode = True
        elif opt in ('-c', '--conf'):
            conf_path = arg
        elif opt in ('-t', '--task'):
            task = arg

    # Run autoremove
    try:
        # Load configurations
        lg.info('Loading configurations...')
        with open(conf_path, 'r') as stream:
            result = yaml.safe_load(stream)
        lg.info('Found %d task(s) in the file.' % len(result))

        # Run tasks
        if task == None: # Task name specified
            for task_name in result:
                Task(task_name, result[task_name], not view_mode).execute()
        else:
            Task(task, result[task], not view_mode).execute()
    except Exception:
        lg.error(traceback.format_exc().splitlines()[-1])
        lg.debug('Exception Logged', exc_info=True)
        lg.critical('An error occured. Please contact the administrator for more information.')

if __name__ == '__main__':
    main()