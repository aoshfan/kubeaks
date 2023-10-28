/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	kubeAksConfigFilePath string
	kubeAksConfigFileName string = ".kubeaks.yaml"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		fmt.Println(os.UserHomeDir())
		// Check if $HOME/.kubeaks.yaml is existing
		/// Check if user have Home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("User didn't have home directory")
		}
		kubeAksConfigFilePath = homeDir + "/" + kubeAksConfigFileName
		fmt.Println(kubeAksConfigFilePath)
		/// Now Check if $HOME/.kubeaks.yaml is existing
		_, err = os.Stat(kubeAksConfigFilePath)

		if os.IsExist(err) {
			fmt.Println("File exists")
			os.Exit(1)
		}
		fmt.Println("File not exists yet")
		// todo: interactive user input using survey ?
		// todo: output as yaml
		// todo: loop again the interactive user input
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
