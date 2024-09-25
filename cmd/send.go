/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	configfile string
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sendCmd.Flags().StringVarP(&configfile, "configpath", "c", "", "config file")
}

func run() {
	if configfile == "" {
		panic("config file is empty")
	}

	bytes, err := os.ReadFile(configfile)
	if err != nil {
		panic(err)
	}

	config := &Config{}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		panic(err)
	}

	
}
