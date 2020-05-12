package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var consumerGroup string

func init() {
	rootCmd.AddCommand(copyCommand)
	copyCommand.Flags().StringVarP(&consumerGroup, "group", "g", "", "consumer group for copy")
}

var copyCommand = &cobra.Command{
	Use:   "cp [SOURCE] [DESTINATION]",
	Short: "Copies kafka topic in a cluster",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("copying...")
	},
}
