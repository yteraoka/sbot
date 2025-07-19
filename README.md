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

### Devices

#### `list`

List all registered devices.

```bash
sbot list
```

#### `describe [DEVICE_ID]`

Shows details for a specific device in JSON format.

```bash
sbot describe <device-id>
```

#### `on [DEVICE_NAME_OR_ID]`

Turns a device on. You can specify the device by its name or ID.

```bash
sbot on "Bedroom Bot"
```

#### `off [DEVICE_NAME_OR_ID]`

Turns a device off. You can specify the device by its name or ID.

```bash
sbot off "Bedroom Bot"
```

#### `brightness [DEVICE_NAME_OR_ID] [LEVEL]`

Sets the brightness for a light device (e.g., Color Bulb). `LEVEL` must be an integer between 1 and 100.

```bash
sbot brightness "My Bulb" 75
```

### Color Bulb

#### `bulb set colortemperature [DEVICE_NAME_OR_ID] [KELVIN]`

Sets the color temperature for a Color Bulb. `KELVIN` must be an integer between 2700 and 6500.

```bash
sbot bulb set colortemperature "My Bulb" 4000
```

### Scenes

#### `scene list`

List all registered scenes.

```bash
sbot scene list
```

#### `scene exec [SCENE_NAME_OR_ID]`

Executes a specific scene.

```bash
sbot scene exec "Movie Time"
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
