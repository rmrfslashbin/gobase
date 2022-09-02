package base

// This module demos the implementation of functional options

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Used to manage varidic options
type Option func(c *Config)

// database configs
type Config struct {
	log         *logrus.Logger
	message     string
	helloString string
}

// New returns a new Config with the given options
func New(opts ...func(*Config)) (*Config, error) {
	config := &Config{}

	// Set a default hello message
	config.helloString = "Hello, World!"

	// apply options
	for _, opt := range opts {
		opt(config)
	}

	// message must be set
	if config.message == "" {
		return nil, fmt.Errorf("message is required")
	}

	// Set up default logger if not set
	if config.log == nil {
		config.log = logrus.New()
		config.log.SetLevel(logrus.InfoLevel)
		config.log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return config, nil
}

// Sets the logger to use
func WithLog(log *logrus.Logger) func(*Config) {
	return func(c *Config) {
		c.log = log
	}
}

// Set the message to use
func WithMessage(message string) func(*Config) {
	return func(c *Config) {
		c.message = message
	}
}

// Set the hello message to use
func WithHelloString(helloString string) func(*Config) {
	return func(c *Config) {
		c.helloString = helloString
	}
}

// Return the hello message
func (c *Config) Hello() string {
	c.log.Info("Hello function")
	return c.helloString
}

// Return the message
func (c *Config) Message() string {
	c.log.Info("Message function")
	return c.message
}
