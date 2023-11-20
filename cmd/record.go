/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
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
		db, err := mysql.NewTreatmentDB()
		if err != nil {
			log.Fatal(err)
		}
		svc := treatment.NewService(db)

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

		recordedTimeEntry := time.Date(
			year, month,
			day, hour,
			minute, 0, 0, time.Now().Location())
		svc.Record(name, recordedTimeEntry)
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")
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

	if recordedTime != "" {
		parsedRecordedTime, err := time.Parse("15:04 PM", recordedTime)
		if err != nil && errors.Is(err, &time.ParseError{}) {
			log.Fatal(fmt.Errorf("expecting a time format of HH:MM PM"))
		}
		hour = parsedRecordedTime.Hour()
		minute = parsedRecordedTime.Minute()
	}
	if recordedDay != "" {
		parsedRecordedDay, err := time.Parse("2006-12-06", recordedDay)
		if err != nil && errors.Is(err, &time.ParseError{}) {
			log.Fatal(fmt.Errorf("expecting a date format of YYYY-MM-DD"))
		}
		year = parsedRecordedDay.Day()
		month = parsedRecordedDay.Month()
		day = parsedRecordedDay.Day()
	}
	if year == 0 {
		year = time.Now().Year()
	}
	if month == 0 {
		month = time.Now().Month()
	}
	if day == 0 {
		day = time.Now().Day()
	}
	if hour == 0 {
		hour = time.Now().Hour()
	}
	if minute == 0 {
		minute = time.Now().Minute()
	}
	return
}
