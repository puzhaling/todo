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
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory.")
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", 
	home+string(os.PathSeparator)+"Desktop"+string(os.PathSeparator)+"src"+string(os.PathSeparator)+"todo"+string(os.PathSeparator)+".todos.json",
	"data file to store todos")
}
