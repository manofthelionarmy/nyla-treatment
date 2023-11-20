/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"nylatreatment/internal/model/medicine"
	"nylatreatment/internal/repository/mysql"
	"nylatreatment/internal/service/medicine/list"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := mysql.NewMedicineDB()
		if err != nil {
			log.Fatal(err)
		}
		svc := list.NewService(repo)
		medicineList, err := svc.List()

		filterByName, err := cmd.Flags().GetString("filter-by-name")
		if err != nil {
			log.Fatal(err)
		}
		if filterByName != "" {
			medicineList = medicineList.Filter(medicine.FilterByName(filterByName))
		}

		filterByType, err := cmd.Flags().GetString("filter-by-type")
		if err != nil {
			log.Fatal(err)
		}

		if filterByType != "" {
			medicineList = medicineList.Filter(medicine.FilterByType(filterByType))
		}

		tbl := tablewriter.NewWriter(os.Stdout)
		tbl.SetHeader([]string{"Name", "Type"})
		for _, m := range medicineList {
			tbl.Append([]string{m.Name, m.Type})
		}
		tbl.Render()
	},
}

func init() {
	medicineCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().String("filter-by-name", "", "Filter medicine by name. Example --filter-by-name='something'")
	listCmd.Flags().String("filter-by-type", "", "Filter medicine by type. Example --filter-by-type='something'")
}
