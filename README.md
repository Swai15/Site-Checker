# HealthChecker CLI Tool

SiteChecker is a command-line tool that allows you to monitor the availability and status of websites.

## Features

- **Single Website Check**: Check the status of a single website.
- **Check All Tracked Websites**: Check the status of all websites being tracked.
- **Add Website**: Add a new website to be monitored.
- **List Tracked Websites**: View a list of all currently tracked websites.
- **Delete Website**: Remove a website from the list of tracked websites.
- **Delete All Websites**: Clear the list of tracked websites.
- **Periodic Checks**: Set intervals for automatic checks of all tracked websites.

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
go get github.com/Swai15/Site-Checker

```

This will install HealthChecker in your Go bin directory.
