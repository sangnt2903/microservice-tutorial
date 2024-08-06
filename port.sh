#!/bin/sh

service=$1
env=$2

# Define the configuration file path
CONFIG_FILE="$service/$env.ini"

# Extract the port value from the configuration file
PORT=$(grep -oP '(?<=port=)\d+' "$CONFIG_FILE")

# Print the port value
echo "$PORT"
