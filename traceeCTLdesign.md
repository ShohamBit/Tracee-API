
Policy Management:
  traceectl policy create <policy_file>
  traceectl policy describe <policy_name>
  traceectl policy list
  traceectl policy update <updated_policy_file>
  traceectl policy delete <policy_name>
  traceectl policy enable <policy_name>
  traceectl policy disable <policy_name>

Event Management:
  traceectl event list
  traceectl event describe <event_name>
  traceectl event enable <event_name>
  traceectl event disable <event_name>
  traceectl event run <event_name> [--args <arguments>]

Stream Management:
  traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream describe <stream_name>
  traceectl stream list
  traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream delete <stream_name>
  traceectl stream connect <stream_name>
  traceectl stream set-default <stream_name>
  traceectl stream pause <stream_name>
  traceectl stream resume <stream_name>

Plugin Management:
  traceectl plugin install --name <plugin_name> --repo <repository_url>
  traceectl plugin list
  traceectl plugin uninstall <plugin_name>

Additional Commands (Potential):
  traceectl connect [<stream_name>]
  traceectl metrics [--output <format>]
  traceectl diagnose [--component <component_name>]
  traceectl logs [--filter <filter>]
  traceectl status
  traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]
  traceectl version

Profile Management:
  traceectl policy profile [create|stop|list|analyze] -f <policy_file> [--name <profile_name>]

Usage:
 traceectl [flags] [options]
 
 Use "traceectl <command> --help" for more information about a given command.
 Use "traceectl options" for a list of global command-line options (applies to all commands).
 
Flags (for client CLI):
    -h, --help   Help for Tracee
    --request-timeout='0':
	The length of time to wait before giving up on a single server request. Non-zero values should contain a
	corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests.
    -s, --server='':
	The address and port of the Kubernetes API server
    -v, --v=0:
	number for the log level verbosity

#=========================================================================================
# Detailed descriptions of commands
#=========================================================================================
**Policy Management:**

*   `traceectl policy create <policy_file>`
    *   Creates a new policy from the YAML file specified by `<policy_file>`.

*   `traceectl policy describe <policy_name>`
    *   Retrieves the details of a specific policy by its name.

*   `traceectl policy list`
    *   Lists all configured policies, providing a brief summary of each.

*   `traceectl policy update <updated_policy_file>`
    *   Updates an existing policy with the contents of a new YAML file.

*   `traceectl policy delete <policy_name>`
    *   Removes a policy by its name.

*   `traceectl policy enable <policy_name>`
    *   Enables a policy by its name.

*   `traceectl policy disable <policy_name>`
    *   Disables a policy by its name.

**Event Management:**

*   `traceectl event list`
    *   Lists all available event definitions (built-in and plugin-defined), providing a brief summary of each.

*   `traceectl event describe <event_name>`
    *   Retrieves the detailed definition of a specific event, including its fields, types, and other metadata.

*   `traceectl event enable <event_name>`
    *   Enables capturing of a specific event type.

*   `traceectl event disable <event_name>`
    *   Disables capturing of a specific event type.

*   `traceectl event run <event_name> [--args <arguments>]`
    *   Manually triggers a user-space event.

**Stream Management:**

*   `traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Creates a new event stream with a specified name.

*   `traceectl stream describe <stream_name>`
    *   Retrieves the configuration details of a stream.

*   `traceectl stream list`
    *   Lists all configured streams.

*   `traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Updates the configuration of an existing stream.

*   `traceectl stream delete <stream_name>`
    *   Removes a stream.

*   `traceectl stream connect <stream_name>`
    *   Connects to an existing stream and displays events in real time, formatted according to the stream's configuration.
    *   The CLI displays a clear warning message explaining that the stream is not configured for stdout output and that connecting to it might have performance implications:
    WARNING: This stream is configured for output to [destination]. Connecting to it might impact performance. Are you sure you want to proceed? (yes/no):


*   `traceectl stream set-default <stream_name>`
    *   Sets the specified stream as the default stream to connect to.

*   `traceectl stream pause <stream_name>`
    *   Temporarily pauses the specified stream, preventing new events from being sent to its destination.

*   `traceectl stream resume <stream_name>`
    *   Resumes a paused stream, allowing new events to be sent to its destination.

**Plugin Management:**

*   `traceectl plugin install --name <plugin_name> --repo <repository_url>`
    *   Installs a plugin from a remote repository.

*   `traceectl plugin list`
    *   Lists all installed plugins.

*   `traceectl plugin uninstall <plugin_name>`
    *   Uninstalls a plugin.

**Additional Commands (Potential):**

*   `traceectl connect [<stream_name>]`:
    *   Connects to a stream and displays events in real time.

*   `traceectl metrics [--output <format>]`:
    *   Retrieves metrics about Tracee's performance and resource usage.

*   `traceectl diagnose [--component <component_name>]`: 
    *   Collects diagnostic information to help troubleshoot issues.

*   `traceectl logs [--filter <filter>]`:
    *   Displays log messages from Tracee, optionally filtered.

*   `traceectl status`:
    *   Shows the status of the Tracee Daemon and its components.

*   `traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]`
    *   View or modify the Tracee Daemon configuration at runtime.

*   `traceectl version`:
    *   Displays the version of Tracee.

The `tracee` command would be a wrapper script that, by default, starts the `traceed` daemon in the background and then forwards the command and its arguments to the `traceectl` client.

Example traceectl status output:

Tracee Daemon Status:
    Status:        Running
    Uptime:        2h 35m 12s
    Version:       0.10.2
    PID:           12345
    Memory Usage:  128 MB
    CPU Usage:     5%

Event Statistics:
    Total Events Captured: 15384
    Events Processed:      14212
    Events Dropped:        1172

Policy Summary:
    Number of Policies: 2

Artifact Capture Status:
    Enabled:       Yes
    Captured Artifacts: 3 network packets, 1 file write
    Storage Location: /tmp/tracee/artifacts

eBPF Probe Status:
    Loaded Probes:   open, openat, execve, ... (list of probe names)
    Failed Probes:   mmap (reason: permission denied), ... (list of probe names and failure reasons)
Policy Management:
  traceectl policy create <policy_file>
  traceectl policy describe <policy_name>
  traceectl policy list
  traceectl policy update <updated_policy_file>
  traceectl policy delete <policy_name>
  traceectl policy enable <policy_name>
  traceectl policy disable <policy_name>

Event Management:
  traceectl event list
  traceectl event describe <event_name>
  traceectl event enable <event_name>
  traceectl event disable <event_name>
  traceectl event run <event_name> [--args <arguments>]

Stream Management:
  traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream describe <stream_name>
  traceectl stream list
  traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream delete <stream_name>
  traceectl stream connect <stream_name>
  traceectl stream set-default <stream_name>
  traceectl stream pause <stream_name>
  traceectl stream resume <stream_name>

Plugin Management:
  traceectl plugin install --name <plugin_name> --repo <repository_url>
  traceectl plugin list
  traceectl plugin uninstall <plugin_name>

Additional Commands (Potential):
  traceectl connect [<stream_name>]
  traceectl metrics [--output <format>]
  traceectl diagnose [--component <component_name>]
  traceectl logs [--filter <filter>]
  traceectl status
  traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]
  traceectl version

Profile Management:
  traceectl policy profile [create|stop|list|analyze] -f <policy_file> [--name <profile_name>]

Usage:
 traceectl [flags] [options]
 
 Use "traceectl <command> --help" for more information about a given command.
 Use "traceectl options" for a list of global command-line options (applies to all commands).
 
Flags (for client CLI):
    -h, --help   Help for Tracee
    --request-timeout='0':
	The length of time to wait before giving up on a single server request. Non-zero values should contain a
	corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests.
    -s, --server='':
	The address and port of the Kubernetes API server
    -v, --v=0:
	number for the log level verbosity

#=========================================================================================
# Detailed descriptions of commands
#=========================================================================================
**Policy Management:**

*   `traceectl policy create <policy_file>`
    *   Creates a new policy from the YAML file specified by `<policy_file>`.

*   `traceectl policy describe <policy_name>`
    *   Retrieves the details of a specific policy by its name.

*   `traceectl policy list`
    *   Lists all configured policies, providing a brief summary of each.

*   `traceectl policy update <updated_policy_file>`
    *   Updates an existing policy with the contents of a new YAML file.

*   `traceectl policy delete <policy_name>`
    *   Removes a policy by its name.

*   `traceectl policy enable <policy_name>`
    *   Enables a policy by its name.

*   `traceectl policy disable <policy_name>`
    *   Disables a policy by its name.

**Event Management:**

*   `traceectl event list`
    *   Lists all available event definitions (built-in and plugin-defined), providing a brief summary of each.

*   `traceectl event describe <event_name>`
    *   Retrieves the detailed definition of a specific event, including its fields, types, and other metadata.

*   `traceectl event enable <event_name>`
    *   Enables capturing of a specific event type.

*   `traceectl event disable <event_name>`
    *   Disables capturing of a specific event type.

*   `traceectl event run <event_name> [--args <arguments>]`
    *   Manually triggers a user-space event.

**Stream Management:**

*   `traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Creates a new event stream with a specified name.

*   `traceectl stream describe <stream_name>`
    *   Retrieves the configuration details of a stream.

*   `traceectl stream list`
    *   Lists all configured streams.

*   `traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Updates the configuration of an existing stream.

*   `traceectl stream delete <stream_name>`
    *   Removes a stream.

*   `traceectl stream connect <stream_name>`
    *   Connects to an existing stream and displays events in real time, formatted according to the stream's configuration.
    *   The CLI displays a clear warning message explaining that the stream is not configured for stdout output and that connecting to it might have performance implications:
    WARNING: This stream is configured for output to [destination]. Connecting to it might impact performance. Are you sure you want to proceed? (yes/no):


*   `traceectl stream set-default <stream_name>`
    *   Sets the specified stream as the default stream to connect to.

*   `traceectl stream pause <stream_name>`
    *   Temporarily pauses the specified stream, preventing new events from being sent to its destination.

*   `traceectl stream resume <stream_name>`
    *   Resumes a paused stream, allowing new events to be sent to its destination.

**Plugin Management:**

*   `traceectl plugin install --name <plugin_name> --repo <repository_url>`
    *   Installs a plugin from a remote repository.

*   `traceectl plugin list`
    *   Lists all installed plugins.

*   `traceectl plugin uninstall <plugin_name>`
    *   Uninstalls a plugin.

**Additional Commands (Potential):**

*   `traceectl connect [<stream_name>]`:
    *   Connects to a stream and displays events in real time.

*   `traceectl metrics [--output <format>]`:
    *   Retrieves metrics about Tracee's performance and resource usage.

*   `traceectl diagnose [--component <component_name>]`: 
    *   Collects diagnostic information to help troubleshoot issues.

*   `traceectl logs [--filter <filter>]`:
    *   Displays log messages from Tracee, optionally filtered.

*   `traceectl status`:
    *   Shows the status of the Tracee Daemon and its components.

*   `traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]`
    *   View or modify the Tracee Daemon configuration at runtime.

*   `traceectl version`:
    *   Displays the version of Tracee.

The `tracee` command would be a wrapper script that, by default, starts the `traceed` daemon in the background and then forwards the command and its arguments to the `traceectl` client.

Example traceectl status output:

Tracee Daemon Status:
    Status:        Running
    Uptime:        2h 35m 12s
    Version:       0.10.2
    PID:           12345
    Memory Usage:  128 MB
    CPU Usage:     5%

Event Statistics:
    Total Events Captured: 15384
    Events Processed:      14212
    Events Dropped:        1172

Policy Summary:
    Number of Policies: 2

Artifact Capture Status:
    Enabled:       Yes
    Captured Artifacts: 3 network packets, 1 file write
    Storage Location: /tmp/tracee/artifacts

eBPF Probe Status:
    Loaded Probes:   open, openat, execve, ... (list of probe names)
    Failed Probes:   mmap (reason: permission denied), ... (list of probe names and failure reasons)
Policy Management:
  traceectl policy create <policy_file>
  traceectl policy describe <policy_name>
  traceectl policy list
  traceectl policy update <updated_policy_file>
  traceectl policy delete <policy_name>
  traceectl policy enable <policy_name>
  traceectl policy disable <policy_name>

Event Management:
  traceectl event list
  traceectl event describe <event_name>
  traceectl event enable <event_name>
  traceectl event disable <event_name>
  traceectl event run <event_name> [--args <arguments>]

Stream Management:
  traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream describe <stream_name>
  traceectl stream list
  traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream delete <stream_name>
  traceectl stream connect <stream_name>
  traceectl stream set-default <stream_name>
  traceectl stream pause <stream_name>
  traceectl stream resume <stream_name>

Plugin Management:
  traceectl plugin install --name <plugin_name> --repo <repository_url>
  traceectl plugin list
  traceectl plugin uninstall <plugin_name>

Additional Commands (Potential):
  traceectl connect [<stream_name>]
  traceectl metrics [--output <format>]
  traceectl diagnose [--component <component_name>]
  traceectl logs [--filter <filter>]
  traceectl status
  traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]
  traceectl version

Profile Management:
  traceectl policy profile [create|stop|list|analyze] -f <policy_file> [--name <profile_name>]

Usage:
 traceectl [flags] [options]
 
 Use "traceectl <command> --help" for more information about a given command.
 Use "traceectl options" for a list of global command-line options (applies to all commands).
 
Flags (for client CLI):
    -h, --help   Help for Tracee
    --request-timeout='0':
	The length of time to wait before giving up on a single server request. Non-zero values should contain a
	corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests.
    -s, --server='':
	The address and port of the Kubernetes API server
    -v, --v=0:
	number for the log level verbosity

#=========================================================================================
# Detailed descriptions of commands
#=========================================================================================
**Policy Management:**

*   `traceectl policy create <policy_file>`
    *   Creates a new policy from the YAML file specified by `<policy_file>`.

*   `traceectl policy describe <policy_name>`
    *   Retrieves the details of a specific policy by its name.

*   `traceectl policy list`
    *   Lists all configured policies, providing a brief summary of each.

*   `traceectl policy update <updated_policy_file>`
    *   Updates an existing policy with the contents of a new YAML file.

*   `traceectl policy delete <policy_name>`
    *   Removes a policy by its name.

*   `traceectl policy enable <policy_name>`
    *   Enables a policy by its name.

*   `traceectl policy disable <policy_name>`
    *   Disables a policy by its name.

**Event Management:**

*   `traceectl event list`
    *   Lists all available event definitions (built-in and plugin-defined), providing a brief summary of each.

*   `traceectl event describe <event_name>`
    *   Retrieves the detailed definition of a specific event, including its fields, types, and other metadata.

*   `traceectl event enable <event_name>`
    *   Enables capturing of a specific event type.

*   `traceectl event disable <event_name>`
    *   Disables capturing of a specific event type.

*   `traceectl event run <event_name> [--args <arguments>]`
    *   Manually triggers a user-space event.

**Stream Management:**

*   `traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Creates a new event stream with a specified name.

*   `traceectl stream describe <stream_name>`
    *   Retrieves the configuration details of a stream.

*   `traceectl stream list`
    *   Lists all configured streams.

*   `traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Updates the configuration of an existing stream.

*   `traceectl stream delete <stream_name>`
    *   Removes a stream.

*   `traceectl stream connect <stream_name>`
    *   Connects to an existing stream and displays events in real time, formatted according to the stream's configuration.
    *   The CLI displays a clear warning message explaining that the stream is not configured for stdout output and that connecting to it might have performance implications:
    WARNING: This stream is configured for output to [destination]. Connecting to it might impact performance. Are you sure you want to proceed? (yes/no):


*   `traceectl stream set-default <stream_name>`
    *   Sets the specified stream as the default stream to connect to.

*   `traceectl stream pause <stream_name>`
    *   Temporarily pauses the specified stream, preventing new events from being sent to its destination.

*   `traceectl stream resume <stream_name>`
    *   Resumes a paused stream, allowing new events to be sent to its destination.

**Plugin Management:**

*   `traceectl plugin install --name <plugin_name> --repo <repository_url>`
    *   Installs a plugin from a remote repository.

*   `traceectl plugin list`
    *   Lists all installed plugins.

*   `traceectl plugin uninstall <plugin_name>`
    *   Uninstalls a plugin.

**Additional Commands (Potential):**

*   `traceectl connect [<stream_name>]`:
    *   Connects to a stream and displays events in real time.

*   `traceectl metrics [--output <format>]`:
    *   Retrieves metrics about Tracee's performance and resource usage.

*   `traceectl diagnose [--component <component_name>]`: 
    *   Collects diagnostic information to help troubleshoot issues.

*   `traceectl logs [--filter <filter>]`:
    *   Displays log messages from Tracee, optionally filtered.

*   `traceectl status`:
    *   Shows the status of the Tracee Daemon and its components.

*   `traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]`
    *   View or modify the Tracee Daemon configuration at runtime.

*   `traceectl version`:
    *   Displays the version of Tracee.

The `tracee` command would be a wrapper script that, by default, starts the `traceed` daemon in the background and then forwards the command and its arguments to the `traceectl` client.

Example traceectl status output:

Tracee Daemon Status:
    Status:        Running
    Uptime:        2h 35m 12s
    Version:       0.10.2
    PID:           12345
    Memory Usage:  128 MB
    CPU Usage:     5%

Event Statistics:
    Total Events Captured: 15384
    Events Processed:      14212
    Events Dropped:        1172

Policy Summary:
    Number of Policies: 2

Artifact Capture Status:
    Enabled:       Yes
    Captured Artifacts: 3 network packets, 1 file write
    Storage Location: /tmp/tracee/artifacts

eBPF Probe Status:
    Loaded Probes:   open, openat, execve, ... (list of probe names)
    Failed Probes:   mmap (reason: permission denied), ... (list of probe names and failure reasons)
Policy Management:
  traceectl policy create <policy_file>
  traceectl policy describe <policy_name>
  traceectl policy list
  traceectl policy update <updated_policy_file>
  traceectl policy delete <policy_name>
  traceectl policy enable <policy_name>
  traceectl policy disable <policy_name>

Event Management:
  traceectl event list
  traceectl event describe <event_name>
  traceectl event enable <event_name>
  traceectl event disable <event_name>
  traceectl event run <event_name> [--args <arguments>]

Stream Management:
  traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream describe <stream_name>
  traceectl stream list
  traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream delete <stream_name>
  traceectl stream connect <stream_name>
  traceectl stream set-default <stream_name>
  traceectl stream pause <stream_name>
  traceectl stream resume <stream_name>

Plugin Management:
  traceectl plugin install --name <plugin_name> --repo <repository_url>
  traceectl plugin list
  traceectl plugin uninstall <plugin_name>

Additional Commands (Potential):
  traceectl connect [<stream_name>]
  traceectl metrics [--output <format>]
  traceectl diagnose [--component <component_name>]
  traceectl logs [--filter <filter>]
  traceectl status
  traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]
  traceectl version

Profile Management:
  traceectl policy profile [create|stop|list|analyze] -f <policy_file> [--name <profile_name>]

Usage:
 traceectl [flags] [options]
 
 Use "traceectl <command> --help" for more information about a given command.
 Use "traceectl options" for a list of global command-line options (applies to all commands).
 
Flags (for client CLI):
    -h, --help   Help for Tracee
    --request-timeout='0':
	The length of time to wait before giving up on a single server request. Non-zero values should contain a
	corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests.
    -s, --server='':
	The address and port of the Kubernetes API server
    -v, --v=0:
	number for the log level verbosity

#=========================================================================================
# Detailed descriptions of commands
#=========================================================================================
**Policy Management:**

*   `traceectl policy create <policy_file>`
    *   Creates a new policy from the YAML file specified by `<policy_file>`.

*   `traceectl policy describe <policy_name>`
    *   Retrieves the details of a specific policy by its name.

*   `traceectl policy list`
    *   Lists all configured policies, providing a brief summary of each.

*   `traceectl policy update <updated_policy_file>`
    *   Updates an existing policy with the contents of a new YAML file.

*   `traceectl policy delete <policy_name>`
    *   Removes a policy by its name.

*   `traceectl policy enable <policy_name>`
    *   Enables a policy by its name.

*   `traceectl policy disable <policy_name>`
    *   Disables a policy by its name.

**Event Management:**

*   `traceectl event list`
    *   Lists all available event definitions (built-in and plugin-defined), providing a brief summary of each.

*   `traceectl event describe <event_name>`
    *   Retrieves the detailed definition of a specific event, including its fields, types, and other metadata.

*   `traceectl event enable <event_name>`
    *   Enables capturing of a specific event type.

*   `traceectl event disable <event_name>`
    *   Disables capturing of a specific event type.

*   `traceectl event run <event_name> [--args <arguments>]`
    *   Manually triggers a user-space event.

**Stream Management:**

*   `traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Creates a new event stream with a specified name.

*   `traceectl stream describe <stream_name>`
    *   Retrieves the configuration details of a stream.

*   `traceectl stream list`
    *   Lists all configured streams.

*   `traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Updates the configuration of an existing stream.

*   `traceectl stream delete <stream_name>`
    *   Removes a stream.

*   `traceectl stream connect <stream_name>`
    *   Connects to an existing stream and displays events in real time, formatted according to the stream's configuration.
    *   The CLI displays a clear warning message explaining that the stream is not configured for stdout output and that connecting to it might have performance implications:
    WARNING: This stream is configured for output to [destination]. Connecting to it might impact performance. Are you sure you want to proceed? (yes/no):


*   `traceectl stream set-default <stream_name>`
    *   Sets the specified stream as the default stream to connect to.

*   `traceectl stream pause <stream_name>`
    *   Temporarily pauses the specified stream, preventing new events from being sent to its destination.

*   `traceectl stream resume <stream_name>`
    *   Resumes a paused stream, allowing new events to be sent to its destination.

**Plugin Management:**

*   `traceectl plugin install --name <plugin_name> --repo <repository_url>`
    *   Installs a plugin from a remote repository.

*   `traceectl plugin list`
    *   Lists all installed plugins.

*   `traceectl plugin uninstall <plugin_name>`
    *   Uninstalls a plugin.

**Additional Commands (Potential):**

*   `traceectl connect [<stream_name>]`:
    *   Connects to a stream and displays events in real time.

*   `traceectl metrics [--output <format>]`:
    *   Retrieves metrics about Tracee's performance and resource usage.

*   `traceectl diagnose [--component <component_name>]`: 
    *   Collects diagnostic information to help troubleshoot issues.

*   `traceectl logs [--filter <filter>]`:
    *   Displays log messages from Tracee, optionally filtered.

*   `traceectl status`:
    *   Shows the status of the Tracee Daemon and its components.

*   `traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]`
    *   View or modify the Tracee Daemon configuration at runtime.

*   `traceectl version`:
    *   Displays the version of Tracee.

The `tracee` command would be a wrapper script that, by default, starts the `traceed` daemon in the background and then forwards the command and its arguments to the `traceectl` client.

Example traceectl status output:

Tracee Daemon Status:
    Status:        Running
    Uptime:        2h 35m 12s
    Version:       0.10.2
    PID:           12345
    Memory Usage:  128 MB
    CPU Usage:     5%

Event Statistics:
    Total Events Captured: 15384
    Events Processed:      14212
    Events Dropped:        1172

Policy Summary:
    Number of Policies: 2

Artifact Capture Status:
    Enabled:       Yes
    Captured Artifacts: 3 network packets, 1 file write
    Storage Location: /tmp/tracee/artifacts

eBPF Probe Status:
    Loaded Probes:   open, openat, execve, ... (list of probe names)
    Failed Probes:   mmap (reason: permission denied), ... (list of probe names and failure reasons)Policy Management:
  traceectl policy create <policy_file>
  traceectl policy describe <policy_name>
  traceectl policy list
  traceectl policy update <updated_policy_file>
  traceectl policy delete <policy_name>
  traceectl policy enable <policy_name>
  traceectl policy disable <policy_name>

Event Management:
  traceectl event list
  traceectl event describe <event_name>
  traceectl event enable <event_name>
  traceectl event disable <event_name>
  traceectl event run <event_name> [--args <arguments>]

Stream Management:
  traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream describe <stream_name>
  traceectl stream list
  traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]
  traceectl stream delete <stream_name>
  traceectl stream connect <stream_name>
  traceectl stream set-default <stream_name>
  traceectl stream pause <stream_name>
  traceectl stream resume <stream_name>

Plugin Management:
  traceectl plugin install --name <plugin_name> --repo <repository_url>
  traceectl plugin list
  traceectl plugin uninstall <plugin_name>

Additional Commands (Potential):
  traceectl connect [<stream_name>]
  traceectl metrics [--output <format>]
  traceectl diagnose [--component <component_name>]
  traceectl logs [--filter <filter>]
  traceectl status
  traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]
  traceectl version

Profile Management:
  traceectl policy profile [create|stop|list|analyze] -f <policy_file> [--name <profile_name>]

Usage:
 traceectl [flags] [options]
 
 Use "traceectl <command> --help" for more information about a given command.
 Use "traceectl options" for a list of global command-line options (applies to all commands).
 
Flags (for client CLI):
    -h, --help   Help for Tracee
    --request-timeout='0':
	The length of time to wait before giving up on a single server request. Non-zero values should contain a
	corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests.
    -s, --server='':
	The address and port of the Kubernetes API server
    -v, --v=0:
	number for the log level verbosity

#=========================================================================================
# Detailed descriptions of commands
#=========================================================================================
**Policy Management:**

*   `traceectl policy create <policy_file>`
    *   Creates a new policy from the YAML file specified by `<policy_file>`.

*   `traceectl policy describe <policy_name>`
    *   Retrieves the details of a specific policy by its name.

*   `traceectl policy list`
    *   Lists all configured policies, providing a brief summary of each.

*   `traceectl policy update <updated_policy_file>`
    *   Updates an existing policy with the contents of a new YAML file.

*   `traceectl policy delete <policy_name>`
    *   Removes a policy by its name.

*   `traceectl policy enable <policy_name>`
    *   Enables a policy by its name.

*   `traceectl policy disable <policy_name>`
    *   Disables a policy by its name.

**Event Management:**

*   `traceectl event list`
    *   Lists all available event definitions (built-in and plugin-defined), providing a brief summary of each.

*   `traceectl event describe <event_name>`
    *   Retrieves the detailed definition of a specific event, including its fields, types, and other metadata.

*   `traceectl event enable <event_name>`
    *   Enables capturing of a specific event type.

*   `traceectl event disable <event_name>`
    *   Disables capturing of a specific event type.

*   `traceectl event run <event_name> [--args <arguments>]`
    *   Manually triggers a user-space event.

**Stream Management:**

*   `traceectl stream create --name <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Creates a new event stream with a specified name.

*   `traceectl stream describe <stream_name>`
    *   Retrieves the configuration details of a stream.

*   `traceectl stream list`
    *   Lists all configured streams.

*   `traceectl stream update <stream_name> [--destination <destination>] [--format <format>] [--fields <fields>] [--parse-data] [--filter <filter>]`
    *   Updates the configuration of an existing stream.

*   `traceectl stream delete <stream_name>`
    *   Removes a stream.

*   `traceectl stream connect <stream_name>`
    *   Connects to an existing stream and displays events in real time, formatted according to the stream's configuration.
    *   The CLI displays a clear warning message explaining that the stream is not configured for stdout output and that connecting to it might have performance implications:
    WARNING: This stream is configured for output to [destination]. Connecting to it might impact performance. Are you sure you want to proceed? (yes/no):


*   `traceectl stream set-default <stream_name>`
    *   Sets the specified stream as the default stream to connect to.

*   `traceectl stream pause <stream_name>`
    *   Temporarily pauses the specified stream, preventing new events from being sent to its destination.

*   `traceectl stream resume <stream_name>`
    *   Resumes a paused stream, allowing new events to be sent to its destination.

**Plugin Management:**

*   `traceectl plugin install --name <plugin_name> --repo <repository_url>`
    *   Installs a plugin from a remote repository.

*   `traceectl plugin list`
    *   Lists all installed plugins.

*   `traceectl plugin uninstall <plugin_name>`
    *   Uninstalls a plugin.

**Additional Commands (Potential):**

*   `traceectl connect [<stream_name>]`:
    *   Connects to a stream and displays events in real time.

*   `traceectl metrics [--output <format>]`:
    *   Retrieves metrics about Tracee's performance and resource usage.

*   `traceectl diagnose [--component <component_name>]`: 
    *   Collects diagnostic information to help troubleshoot issues.

*   `traceectl logs [--filter <filter>]`:
    *   Displays log messages from Tracee, optionally filtered.

*   `traceectl status`:
    *   Shows the status of the Tracee Daemon and its components.

*   `traceectl config [set|get|update] [<option>=<value>] [--file <config_file>]`
    *   View or modify the Tracee Daemon configuration at runtime.

*   `traceectl version`:
    *   Displays the version of Tracee.

The `tracee` command would be a wrapper script that, by default, starts the `traceed` daemon in the background and then forwards the command and its arguments to the `traceectl` client.

Example traceectl status output:

Tracee Daemon Status:
    Status:        Running
    Uptime:        2h 35m 12s
    Version:       0.10.2
    PID:           12345
    Memory Usage:  128 MB
    CPU Usage:     5%

Event Statistics:
    Total Events Captured: 15384
    Events Processed:      14212
    Events Dropped:        1172

Policy Summary:
    Number of Policies: 2

Artifact Capture Status:
    Enabled:       Yes
    Captured Artifacts: 3 network packets, 1 file write
    Storage Location: /tmp/tracee/artifacts

eBPF Probe Status:
    Loaded Probes:   open, openat, execve, ... (list of probe names)
    Failed Probes:   mmap (reason: permission denied), ... (list of probe names and failure reasons)