/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rmrfslashbin/gobase/pkg/base"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Catch-all errors
		var err error
		defer func() {
			if err != nil {
				log.WithFields(logrus.Fields{
					"error": err,
				}).Fatal("main crashed")
			}
		}()

		// Run the dump command and cache the error
		if err := runDump(); err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Fatal("error")
		}

	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)

	// Example of how to bind a flag to a configuration variable with Viper
	dumpCmd.Flags().String("dumpMessage", "this is the default message", "dump message")
	viper.BindPFlag("dumpMessage", dumpCmd.Flags().Lookup("dumpMessage"))
}

func runDump() error {
	// Show the dump message via Viper
	log.WithFields(logrus.Fields{
		"dumpMessage": viper.GetString("dumpMessage"),
	}).Info("dump cmd!")

	b, err := base.New(
		base.WithLog(log),
		base.WithHelloString("Hello, world... from the dump command!"),
		base.WithMessage(viper.GetString("dumpMessage")),
	)
	if err != nil {
		return err
	}

	log.WithFields(logrus.Fields{
		"message": b.Message(),
	}).Info("message")
	log.WithFields(logrus.Fields{
		"hello": b.Hello(),
	}).Info("hello")

	return nil
}
