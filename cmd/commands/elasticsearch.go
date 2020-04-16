package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-elasticsearch/modules/elasticsearch"
)

func newRebuildIndexCmd() *cobra.Command {
	rebuildIndexCmd := &cobra.Command{
		Use: "rebuild",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("rebuild index \n")
			elasticsearch.InitClient("http://localhost:9200")
			if err := elasticsearch.DeleteBooksIndex(); err != nil {
				fmt.Printf("delete failed: %s", err)
			}

			if err := elasticsearch.CreateIndices(); err != nil {
				fmt.Printf("create indices failed: %s", err)
			}
			fmt.Print("finish rebuild index \n")
		},
	}
	return rebuildIndexCmd
}
