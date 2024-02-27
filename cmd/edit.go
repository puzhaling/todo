/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"os"

	"github.com/spf13/cobra"
	"github.com/puzhaling/todo/cont"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit todos",
	Run: editRun,
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func editRun(cmd *cobra.Command, args []string) {
	items, err := cont.ReadItems(dataFile)

	if err != nil {
		fmt.Printf("%v", err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(i, "arg is not convertible to int")
		os.Exit(1)
	}

	if i > 0 && i - 1 < len(items) {
		fmt.Print("Write new todo value: ")

		var text string
		if _, err := fmt.Scanln(&text); err == nil {
			items[i-1].Text = text
		}
	} else {
		fmt.Println("Out of range")
	}

	cont.SaveItems(dataFile, items)
}