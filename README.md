# WeatherCLI

<b/>Learning Project:</b> A command-line tool to fetch weather data from Weather API for a Public Location. and convert it to different output formats.

Project based on this [repo](https://github.com/olgazju/weather-project.git)

This project gets also includes a CI/CD pipeline using GitHub actions and pushes the Docker Image to Docker Hub.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [License](#license)

## Requirements

- Golang 1.17 or higher
- A CDO API token (you can get one for free from https://www.ncdc.noaa.gov/cdo-web/)
- DockerHub Account
- Github Actions

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

## Configuration

In order to connect to CDO API set up a `.env` file int the project directory with following variable: <br />
`CDO_TOKEN=token`

In order to use the CI/CD Pipeline ensure you create:

1. DockerHub Account
2. Create Personal Access Token
3. Add them to the Repo's Secrets

## Usage

Run the following command to see the available options:

```bash
go run main.go [flags]
```

<br />
Flags:

-csv: Enable CSV output (default: false) <br />
-json: Enable JSON output (default: false) <br />
Example usage:

```bash

go run main.go -csv -json
```

<br />
To use CI/CD Pipeline make sure you have the requiremens mentioned in the [Configuration](#configuration) then:

1. Run this command to build a development image:

```bash
docker build -t weathercli:dev .
```

2. To run the image use:

```bash
docker run --rm weathercli:dev
```

3. To build the image locally for production:

```bash
docker build -t golang-pipeline:1.0.0 . --build-arg VERSION=1.0.0
```

4. After that push the changes to "branch-name"

5. Check if the Tests ran properly, and if not fix them.

6. Run these final commands to push these to run final tests and push to DockerHub

```bash
git tag v1.0.0
git push --tags
```

## License

This project is licensed under the [MIT License](LICENSE).
