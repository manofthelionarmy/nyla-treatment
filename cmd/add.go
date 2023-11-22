/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"nylatreatment/internal/cloud/google"
	"nylatreatment/internal/repository/mysql"
	"nylatreatment/internal/service/calendar"
	"nylatreatment/internal/service/treatment"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
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

		calendarSvc := google.NewCalendarService()
		adapter := calendar.NewAdapter(calendarSvc)

		repo, err := mysql.NewTreatmentDB()
		if err != nil {
			log.Fatal(err)
		}
		treatmentSvc := treatment.NewService(repo)
		if name == "" {
			fmt.Println("add all treatments to the calendar")
			return
		}
		mr, err := treatmentSvc.GetMedicineNextTreatment(name)
		if err != nil {
			log.Fatal(err)
		}

		err = adapter.AddToCalendar(*mr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	calendarCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().String("name", "", "This value is the name of the medicine and it's next treatment time will be added to google calendar")
}
