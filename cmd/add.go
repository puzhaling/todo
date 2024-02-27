/*
Copyright © 2024 Hermann Puzhalin <puzhaling@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/puzhaling/todo/cont"
)

var priority int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "priority:1,2,3")
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := cont.ReadItems(dataFile)
	for _, x := range args {
		item := cont.Item{Text:x}
		item.SetPriority(priority)
		items = append(items, item)
	}
	
	err = cont.SaveItems(dataFile, items)

	if err != nil {
		fmt.Errorf("%v", err)
	}
}
