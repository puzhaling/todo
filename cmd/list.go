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


// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the todos",
	Long: "list the todos in pretty way",
	Run: listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := cont.ReadItems(dataFile)

	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(cont.ByPri(items))
	
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, item := range items {
		fmt.Fprintln(w, item.Label()+" "+item.PrettyP()+"\t"+item.Text+"\t")
	}

	w.Flush()
}