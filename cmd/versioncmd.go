/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/broker/generated"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Provides version information about the current binary",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Provider version: %s, based on Semp version %s", broker.ProviderVersion, generated.SempVersion))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
