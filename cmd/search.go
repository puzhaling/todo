/*
Copyright © 2024 Hermann Puzhalin <puzhaling@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/puzhaling/todo/cont"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search todos",
	Run: searchRun,
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func searchRun(cmd *cobra.Command, args []string) {
	items, err := cont.ReadItems(dataFile)

	if err != nil {
		fmt.Printf("%v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)

	for _, item := range items {
		if regexpObj, _:= regexp.MatchString(args[0], item.Text); regexpObj {
			fmt.Fprintln(w, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Text+"\t")
		}
	}

	w.Flush()
}
