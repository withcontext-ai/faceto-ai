#!/bin/bash

# Use this script to run the server in a docker container

set -euo pipefail

./faceto-ai -conf ./configs/config.yaml -env prod