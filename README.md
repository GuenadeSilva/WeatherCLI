# WeatherCLI

<b/>Learning Project:</b> A command-line tool to fetch weather data from Weather API for a Public Location. and convert it to different output formats.

Project based on this [repo](https://github.com/olgazju/weather-project.git)

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [License](#license)

## Requirements

- Golang 1.17 or higher
- A CDO API token (you can get one for free from https://www.ncdc.noaa.gov/cdo-web/)

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/GuenadeSilva/WeatherCLI.git
   cd weathercli
   ```

2. Install necessary dependencies

   ```bash
   go mod tidy
   ```

## Usage

Run the following command to see the available options:

```bash
go run main.go [flags]
```

Flags:

-csv: Enable CSV output (default: false) <br />
-json: Enable JSON output (default: false) <br />
Example usage:

```bash

go run main.go -csv -json
```

## Configuration

In order to connect to CDO API set up a `.env` file int the project directory with following variable:
CDO_TOKEN=token

## License

This project is licensed under the [MIT License](LICENSE).
