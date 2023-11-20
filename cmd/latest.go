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

// lastestCmd represents the lastest command
var latestCmd = &cobra.Command{
	Use:   "latest",
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
		medicineRecord, err := getLatestMedicineRecord(svc, nameFlagValue)
		if err != nil {
			log.Fatal(err)
		}
		tbl := tablewriter.NewWriter(os.Stdout)
		tbl.SetHeader([]string{"Name", "Time Taken"})

		defaultFmt := "Mon, Jan _2, 2006 3:04PM"

		timeTaken := medicineRecord.TimeTaken.Format(defaultFmt)
		tbl.Append([]string{medicineRecord.Name, timeTaken})
		tbl.Render()
	},
}

func getLatestMedicineRecord(svc treatment.Service, name string) (*medicine.MedicineRecord, error) {
	if name != "" {
		return svc.GetMedicineLastTreatment(name)
	}
	return svc.GetLatestTreatment()
}

func init() {
	rootCmd.AddCommand(latestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lastestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	latestCmd.Flags().String("name", "", "Specify name of medicine to get the last time it was taken")
}
