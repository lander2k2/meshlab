#!/bin/bash
locust -f /locust_file.py --host=http://$MESHLAB_URL

