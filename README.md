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

### Generic Device Commands

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

#### `run-customize [DEVICE_NAME_OR_ID] [BUTTON_NAME]`

Executes a custom button on an infrared remote.

```bash
sbot run-customize "Living Room TV" "Menu"
```

### Color Bulb Commands

#### `bulb set brightness [DEVICE_NAME_OR_ID] [LEVEL]`

Sets the brightness for a Color Bulb. `LEVEL` must be an integer between 1 and 100.

```bash
sbot bulb set brightness "My Bulb" 75
```

#### `bulb set colortemperature [DEVICE_NAME_OR_ID] [KELVIN]`

Sets the color temperature for a Color Bulb. `KELVIN` must be an integer between 2700 and 6500.

```bash
sbot bulb set colortemperature "My Bulb" 4000
```

### TV Commands

#### `tv set-channel [DEVICE_NAME_OR_ID] [CHANNEL]`

Sets the channel for a TV.

```bash
sbot tv set-channel "Living Room TV" 5
```

#### `tv volume-up [DEVICE_NAME_OR_ID]`

Increases the volume of a TV.

```bash
sbot tv volume-up "Living Room TV"
```

#### `tv volume-down [DEVICE_NAME_OR_ID]`

Decreases the volume of a TV.

```bash
sbot tv volume-down "Living Room TV"
```

#### `tv channel-up [DEVICE_NAME_OR_ID]`

Changes to the next channel.

```bash
sbot tv channel-up "Living Room TV"
```

#### `tv channel-down [DEVICE_NAME_OR_ID]`

Changes to the previous channel.

```bash
sbot tv channel-down "Living Room TV"
```

### Air Conditioner Commands

#### `ac on [DEVICE_NAME_OR_ID] --temperature [TEMP] --mode [MODE] --fan-speed [SPEED]`

Turns on an Air Conditioner with specified settings. Temperature is required.

*   `--temperature`, `-t`: Temperature in Celsius.
*   `--mode`, `-m`: `auto`, `cool`, `dry`, `fan`, `heat` (default: `auto`).
*   `--fan-speed`, `-f`: `auto`, `low`, `medium`, `high` (default: `auto`).

```bash
sbot ac on "Bedroom AC" -t 25 --mode cool --fan-speed medium
```

#### `ac off [DEVICE_NAME_OR_ID]`

Turns off an Air Conditioner.

```bash
sbot ac off "Bedroom AC"
```

### Scene Commands

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