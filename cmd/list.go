/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/himanshuk42/chash/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the notes saved in the notes database",
	Long:  `See the list of all the notes created and saved in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DisplayAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
