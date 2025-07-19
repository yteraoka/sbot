# SwitchBot Device Operations

This document outlines the available commands for various SwitchBot devices through the official API.

## Common Devices

### Bot

*   **`turnOn`**: Turns the Bot on (presses a switch).
*   **`turnOff`**: Turns the Bot off (pulls a switch).
*   **`press`**: A single press action.

### Curtain

*   **`turnOn`**: Opens the curtains.
*   **`turnOff`**: Closes the curtains.
*   **`setPosition`**: Sets the curtain to a specific position.
    *   **Parameter**: `0,ff,{position}` where `position` is 0-100.

### Plug

*   **`turnOn`**: Turns the plug on.
*   **`turnOff`**: Turns the plug off.

### Lock

*   **`lock`**: Locks the door.
*   **`unlock`**: Unlocks the door.

## Lighting

### Color Bulb, Strip Light, Ceiling Light

*   **`turnOn`**: Turns the light on.
*   **`turnOff`**: Turns the light off.
*   **`toggle`**: Toggles the light's power state.
*   **`setBrightness`**: Adjusts the brightness level.
    *   **Parameter**: An integer from 1 to 100.
*   **`setColor`**: Sets the color of the light.
    *   **Parameter**: RGB values in the format `"R:G:B"` (e.g., `"255:128:0"`).
*   **`setColorTemperature`**: Sets the color temperature.
    *   **Parameter**: An integer from 2700 to 6500 (Kelvin).

## Sensors

Sensors like the **Meter**, **Motion Sensor**, and **Contact Sensor** do not have controllable commands. Their data (temperature, humidity, motion, etc.) is retrieved via the device status endpoint.

## Other Appliances

### Humidifier

*   **`turnOn`**: Turns the humidifier on.
*   **`turnOff`**: Turns the humidifier off.
*   **`setMode`**: Sets the humidification mode.
    *   **Parameter**: `auto` or a numeric value representing the percentage (e.g., `50`).

### Robot Vacuum Cleaner

*   **`start`**: Starts the cleaning cycle.
*   **`stop`**: Stops the cleaning cycle.
*   **`dock`**: Sends the vacuum back to its charging dock.

For more details and for commands related to other devices, please refer to the [official SwitchBot API documentation](https://github.com/OpenWonderLabs/SwitchBotAPI).
