/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"nylatreatment/internal/repository/mysql"
	"nylatreatment/internal/service/treatment"
	"time"

	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Record the day and time nyla took medicine",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		timeRecordedFlagValue, err := cmd.Flags().GetString("time")
		if err != nil {
			log.Fatal(err)
		}
		if timeRecordedFlagValue == "" {
			timeRecordedFlagValue = time.Now().String()
		}

		dayRecordedFlag, err := cmd.Flags().GetString("date")
		if err != nil {
			log.Fatal(err)
		}

		year, month, day, hour, minute := getRecordedTime(timeRecordedFlagValue, dayRecordedFlag)

		db, err := mysql.NewTreatmentDB()
		if err != nil {
			log.Fatal(err)
		}
		svc := treatment.NewService(db)

		recordedTimeEntry := time.Date(
			year, month,
			day, hour,
			minute, 0, 0, time.Now().Location())
		// handle error here
		if err := svc.Record(name, recordedTimeEntry); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// TODO: provide option to query by type
	recordCmd.PersistentFlags().String("name", "", "The name of the medicine")
	recordCmd.MarkFlagRequired("name")

	recordCmd.PersistentFlags().String("time", "", "The time the medicine was taken")
	recordCmd.PersistentFlags().String("date", "", "The day the medicine was taken. YYYY-MM-DD format")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getRecordedTime(recordedTime, recordedDay string) (
	year int,
	month time.Month,
	day int,
	hour int,
	minute int,
) {

	// initialize these values
	year, month, day = time.Now().Date()
	hour, minute, _ = time.Now().Clock()

	// overwrite hour and minute if we pass in a time
	if recordedTime != "" {
		parsedRecordedTime, err := time.Parse("15:04 PM", recordedTime)
		if err != nil {
			log.Fatal(fmt.Errorf("expecting a time format of HH:MM PM"))
		}
		hour, minute, _ = parsedRecordedTime.Clock()
	}
	// overwrite year, month, and ay if we pass in a date
	if recordedDay != "" {
		parsedRecordedDay, err := time.Parse("2006-12-06", recordedDay)
		if err != nil {
			log.Fatal(fmt.Errorf("expecting a date format of YYYY-MM-DD"))
		}
		year, month, day = parsedRecordedDay.Date()
	}
	return
}
