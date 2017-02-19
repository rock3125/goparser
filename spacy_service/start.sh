#!/bin/bash

# start with large timeout allowing for large texts to be parsed
gunicorn --bind 0.0.0.0:9000 --timeout 120 --threads 1 server:app
