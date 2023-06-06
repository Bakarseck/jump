**Usage Guide for the Script**

This script is designed to monitor a directory for file changes and automatically execute a command when changes are detected. It provides options for customizing the command, directory, exclusion, and file pattern restrictions.

**Usage:**

```bash
./watch-run.sh [OPTIONS]
```

**Options:**

- `-c <command>`: Specifies the command to be executed. By default, the command is set to `go run .`.
- `-d <directory>`: Specifies the directory to be monitored. By default, the current directory is monitored.
- `-e <exclude>`: Specifies a template to be excluded from monitoring. Files matching the template will be ignored when detecting changes.
- `-p <pattern>`: Specifies the file pattern restriction for monitoring. Only files matching the pattern will be considered for changes. By default, it is set to `*.go`.

**Functionality:**

1. The script defines a function `run_command` that runs a command in the background. It captures the process ID (PID) of the command for future reference.
2. The script initializes default values for the command, directory, exclusion, and pattern restrictions.
3. It processes the command-line options using `getopts` to override the default values if provided by the user.
4. The initial command is executed using the `run_command` function.
5. The script enters an infinite loop, continuously monitoring the specified directory for file changes.
6. It uses the `find` command to identify files that have changed since the last run. The changes are determined based on the modification timestamp.
7. The changed files are filtered based on the exclusion and pattern restrictions provided by the user.
8. If any changed files are detected, the previous command is stopped by killing its process using the captured PID. Then, the new command is executed using the `run_command` function.
9. The script updates the last run time and repeats the process after a short interval.

**Examples:**

1. Run the script with default options:
   ```bash
   ./watch-run.sh
   ```
   This will monitor the current directory for file changes and execute the default command `go run .`.

2. Specify a custom command and directory:
   ```bash
   ./watch-run.sh -c "python myscript.py" -d /path/to/directory
   ```
   This will monitor the `/path/to/directory` for file changes and execute the command `python myscript.py`.

3. Exclude a specific file or template from monitoring:
   ```bash
   ./watch-run.sh -e "template.txt"
   ```
   This will exclude files matching the template `template.txt` from being considered for changes.

4. Restrict monitoring to specific file patterns:
   ```bash
   ./watch-run.sh -p "*.py"
   ```
   This will monitor files with the `.py` extension only and ignore other file types.

**Note:** Ensure that the script has proper execution permissions (`chmod +x watch-run.sh`) before running it.

Please note that this script is a modified version of the one you provided earlier, with additional options for customization.