/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"nylatreatment/internal/model/medicine"
	"nylatreatment/internal/repository/mysql"
	"nylatreatment/internal/service/treatment"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// nextTimeCmd represents the nextTime command
var nextTimeCmd = &cobra.Command{
	Use:   "next-time",
	Short: "A brief description of your command",
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
		nameFlagValue, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		// TODO: get all next times for all medicines
		medicineRecords, err := getNextTreatment(svc, nameFlagValue)
		if err != nil {
			log.Fatal(err)
		}
		tbl := tablewriter.NewWriter(os.Stdout)
		tbl.SetHeader([]string{"Name", "Time Taken"})

		defaultFmt := "Mon, Jan _2, 2006 3:04PM"

		for _, mr := range medicineRecords {
			timeTaken := mr.TimeTaken.Format(defaultFmt)
			tbl.Append([]string{mr.Name, timeTaken})
		}
		tbl.Render()
	},
}

func getNextTreatment(svc treatment.Service, name string) ([]medicine.MedicineRecord, error) {
	if name == "" {
		return svc.GetAllMedicinesNextTreatment()
	}
	mr, err := svc.GetMedicineNextTreatment(name)
	if err != nil {
		return nil, err
	}
	return []medicine.MedicineRecord{*mr}, nil
}

func init() {
	rootCmd.AddCommand(nextTimeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nextTimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	nextTimeCmd.Flags().String("name", "", "Get the next time the medicine needs to be taken")
}
