# Catalyst

Catalyst is a command-line tool for controlling smart appliances via the Nature Remo API.

## Features
- Control lights and other appliances registered with Nature Remo
- Simple CLI interface

## Requirements
- Go 1.24
- Nature Remo API token
- Appliance ID (for the device you want to control)

## Getting Started

### Build

```
task build
```

Or manually:

```
go build -o bin/catalyst main.go
```


### Usage

1. **Obtain your Nature Remo API token and Appliance ID:**
   - Visit the [Nature Remo developer site](https://developer.nature.global/) and follow instructions to generate your API token.
   - You can obtain your Appliance ID by calling the [Nature Remo API](https://swagger.nature.global/#/default/get_1_appliances).
   - _You must obtain these yourself. They are not provided by this project._

2. **Configure your credentials:**

You can provide your API token and appliance ID in either of the following ways:

- **Config file:** Create a file named `.catalyst.yaml` in your home directory with the following content:

  ```yaml
  id: <YOUR_APPLIANCE_ID>
  token: <YOUR_API_TOKEN>
  ```

- **Environment variables:**
  - `id`: Appliance ID
  - `token`: Nature Remo API token

3. **Run the CLI:**

To turn the light on:

```
catalyst light on
```

To turn the light off:

```
catalyst light off
```

## License

This project is licensed under the MIT License.
