/*
Copyright © 2024 Hermann Puzhalin <puzhaling@gmail.com>

*/
package cmd

import (
	"fmt"
	"log"
	"text/tabwriter"
	"os"
	"sort"

	"github.com/puzhaling/todo/cont"
	"github.com/spf13/cobra"
)

var (
	doneOpt bool
	allOpt bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the todos",
	Long: "list the todos in pretty way",
	Run: listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all todos")
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := cont.ReadItems(dataFile)

	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(cont.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, item := range items {
		if allOpt || item.Done == doneOpt {
			fmt.Fprintln(w, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Text+"\t")
		}
		// fmt.Fprintln(w, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Text+"\t")
	}

	w.Flush()
}