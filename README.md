# sbot

`sbot` is a command-line interface (CLI) tool for interacting with the SwitchBot API.

## Description

This tool allows you to manage and control your SwitchBot devices from the command line.

## Installation

First, ensure you have Go installed on your system. You can then install `sbot` using `go install`:

```bash
go install github.com/yteraoka/sbot@latest
```

## Configuration

To use `sbot`, you need to provide your SwitchBot API credentials via environment variables.

*   `SWITCHBOT_TOKEN`: Your SwitchBot API token.
*   `SWITCHBOT_CLIENT_SECRET`: Your SwitchBot API client secret.

You can obtain these from the SwitchBot app.

## Usage

### `list`

List all registered devices.

```bash
sbot list
```

**Example Output:**

```
Devices:
ID: 01-202304011234-12345678, Name: Living Room Hub, Type: Hub Mini
ID: 02-202304012345-23456789, Name: Bedroom Bot, Type: Bot
```

### `describe [DEVICE_ID]`

Shows details for a specific device in JSON format.

```bash
sbot describe 02-202304012345-23456789
```

### `on [DEVICE_NAME_OR_ID]`

Turns a device on. You can specify the device by its name or ID.

```bash
# By name
sbot on "Bedroom Bot"

# By ID
sbot on 02-202304012345-23456789
```

### `off [DEVICE_NAME_OR_ID]`

Turns a device off. You can specify the device by its name or ID.

```bash
# By name
sbot off "Bedroom Bot"

# By ID
sbot off 02-202304012345-23456789
```

## Shell Completion

`sbot` supports generating shell completion scripts for Bash and Zsh.

### `completion [bash|zsh]`

Generates the completion script. Follow the instructions printed by the command to install it.

**Example for Bash:**

```bash
source <(sbot completion bash)
```

**Example for Zsh:**

```bash
sbot completion zsh > "${fpath[1]}/_sbot"
```

## License

This project is licensed under the terms of the LICENSE file.
