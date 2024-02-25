/*
Copyright © 2024 Hermann Puzhalin <puzhaling@gmail.com>

*/
package cmd

import (
	"os"
	"fmt"
	"log"
	
	"github.com/spf13/cobra"
	"github.com/mitchellh/go-homedir"
)

var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory.")
	}

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", 
	home+string(os.PathSeparator)+"Desktop"+string(os.PathSeparator)+"src"+string(os.PathSeparator)+"todo"+string(os.PathSeparator)+".todos.json",
	"data file to store todos")
}
