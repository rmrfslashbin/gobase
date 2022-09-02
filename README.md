# gobase
Skeleton Golang project

## Cobra & Viper
This project used Cobra and Viper to generate the CLI commands.
- github.com/spf13/cobra
- github.com/spf13/viper

## Logrus
Logrun is used to log the application events.
- github.com/sirupsen/logrus
This project supports logrus defined log levels: debug, info, warn, error, fatal, panic. Log levels are set via Cobra/Viper flags.

## Config
`config.yaml.DIST` is provided as a sample configuration file.

## Functional Options
This project supports functional options and is implemented/demo'ed in the "base" module.