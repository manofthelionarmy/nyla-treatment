/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// calendarCmd represents the calendar command
var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "View or add calendar events",
	Long: `Uses google calendar api to add calendar 
events with email reminders for the 
next time nyla has to take her medicine`,
}

func init() {
	rootCmd.AddCommand(calendarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calendarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calendarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
