/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ForbiddenR/mqtt-client/pkg/config"
	"github.com/spf13/cobra"
)

var (
	configfile string
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}

func readConfig() (*config.Config, error) {
	if configfile == "" {
		return nil, errors.New("config file is empty")
	}
	if _, err := os.Stat(configfile); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(configfile, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	cfg, err := config.ReadConfig(file)
	return cfg, err
}

func run() {
	cfg, err := readConfig()
	if err != nil {
		println(err)
		return
	}
	fmt.Printf("%+v\n", cfg)
}
