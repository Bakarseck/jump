#!/bin/bash

# Define the function to run a command
run_command() {
    echo "Running command: $@"  # Print the command being executed
    "$@" &                     # Execute the command in the background
    sleep 2                    # Sleep for a short interval to allow the command to start
    PID=$(pgrep -P $$)         # Find the process ID of the command
    echo "Command PID: $PID"   # Print the process ID of the command
}

# Initialize variables with default values
COMMAND="go run ."
DIRECTORY="./"
EXCLUDE=""
RESTRICT="*"
CHECK_INTERVAL=1

# Function to stop the running command gracefully
stop_command() {
    if [ -n "$PID" ]; then
        echo "Stopping previous command..."
        pkill -P $PID 2>/dev/null  # Send SIGTERM to stop the previous command children process
        kill -TERM $PID 2>/dev/null  # Send SIGTERM to stop the previous command gracefully
        wait $PID 2>/dev/null  # Wait for the command to finish
    fi
}

# Process command-line options
while getopts "c:d:e:p:i" opt; do
  case $opt in
    c) COMMAND="$OPTARG" ;;
    d) DIRECTORY="$OPTARG" ;;
    e) EXCLUDE="$OPTARG" ;;
    p) RESTRICT="$OPTARG" ;;
    i) CHECK_INTERVAL="$OPTARG" ;;
    \?) echo "Invalid option: -$OPTARG" >&2; exit 1 ;;
  esac
done

run_command $COMMAND  # Initial command execution

LAST_RUN_TIME=$(date +%T)  # Set the initial last run time

while true; do
    CHANGED_FILES=$(find $DIRECTORY -type f -newermt "$LAST_RUN_TIME" ! -name "$EXCLUDE" -name "$RESTRICT" 2>/dev/null)  # Find files changed since the last run time, excluding the specified template and matching the specified pattern

   if [[ -n "$CHANGED_FILES" ]]; then
        echo "Detected file changes."
        stop_command  # Stop the previous command gracefully
        run_command $COMMAND  # Run the new command
    fi

    LAST_RUN_TIME=$(date +%T)  # Update the last run time
    sleep "$CHECK_INTERVAL"  # Sleep for a short interval before checking for file changes again
done
