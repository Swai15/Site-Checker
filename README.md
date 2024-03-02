# sitechecker CLI Tool

SiteChecker is a command-line tool that allows you to monitor the availability and status of websites.

## Features

- Check the status of a single/ multiple tracked websites.
- Check the status of all websites being tracked.
- Add, list and delete a collection of monitored websites handled via a JSON file
- Periodic intervals for automatic checks of all tracked websites.

## Pending

Notification system upon status change on a monitor mode for windows, linux and mac.

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

This will install sitechecker in your Go bin directory.
