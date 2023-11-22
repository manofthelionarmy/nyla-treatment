/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// medicineCmd represents the medicine command
var medicineCmd = &cobra.Command{
	Use:   "medicine",
	Short: "Get medicine details",
	Long: `A utility for getting nyla's medicine detials. 
You are able to filter by name or type.
	`,
}

func init() {
	rootCmd.AddCommand(medicineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// medicineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// medicineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
