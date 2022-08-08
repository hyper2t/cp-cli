/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var operation string

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		kubeConfig, err := InitKubeConfig(false)
		if err != nil {
			fmt.Errorf("ERROR: %s", err)
		}
		// fmt.Println(kubeConfig)
		clientSet := SetNewClientSet()
		err = CopyToPod(kubeConfig, clientSet)
		if err != nil {
			fmt.Errorf("ERROR: %s", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(cpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
