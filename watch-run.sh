#!/bin/bash

# Define the function to run a command
run_command() {
    echo "Running command: $@"  # Print the command being executed
    "$@" &                     # Execute the command in the background
    sleep 2                    # Sleep for a short interval to allow the command to start
    PID=$(ps aux | grep "/tmp/go" | grep -v grep | awk '{print $2}')  # Find the process ID of the command
    echo "Command PID: $PID"   # Print the process ID of the command
}

# Initialize variables with default values
COMMAND="go run ."
DIRECTORY="./"
EXCLUDE=""
RESTRICT="*.go"

# Process command-line options
while getopts "c:d:e:p:" opt; do
  case $opt in
    c) COMMAND="$OPTARG" ;;
    d) DIRECTORY="$OPTARG" ;;
    e) EXCLUDE="$OPTARG" ;;
    p) RESTRICT="$OPTARG" ;;
    \?) echo "Invalid option: -$OPTARG" >&2; exit 1 ;;
  esac
done

run_command $COMMAND  # Initial command execution

LAST_RUN_TIME=$(date +%T)  # Set the initial last run time

while true; do
    CHANGED_FILES=$(find $DIRECTORY -type f -newermt "$LAST_RUN_TIME" ! -name "$EXCLUDE" -name "$RESTRICT" 2>/dev/null)  # Find files changed since the last run time, excluding the specified template and matching the specified pattern

    if [[ -n "$CHANGED_FILES" ]]; then
        echo "Detected file changes."
        echo "Stopping previous command..."
        kill -9 $PID 2>/dev/null  # Stop the previous command by killing its process
        run_command $COMMAND  # Run the new command
    fi

    LAST_RUN_TIME=$(date +%T)  # Update the last run time
    sleep 1  # Sleep for a short interval before checking for file changes again
done
