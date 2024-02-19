# sitechecker CLI Tool

SiteChecker is a command-line tool that allows you to monitor the availability and status of websites.

## Features

- Check the status of a single website.
- Check the status of all websites being tracked.
- Add a new website to be monitored.
- View a list of all currently tracked websites.
- Remove a single or all websites from list of tracked websites
- Periodic intervals for automatic checks of all tracked websites.

## Usage

To use sitechecker, simply run the executable with appropriate flags. Here are some examples:

1. If you're in the project directory

```
go run . --flag value

```

2. Run it as a build.

```
go build -0 sitechecker

./sitechecker --flag value

```

## Installation

To install sitechecker, you need to have Go installed on your system. Then you can simply run:

```
go get github.com/Swai15/sitechecker

```

This will install HealthChecker in your Go bin directory.
