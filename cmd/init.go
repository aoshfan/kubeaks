/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	kubeAksConfigFilePath string
	kubeAksConfigFileName string = ".kubeaks.yaml"
)

type AksConfigData struct {
	Name string `yaml:"name"`
	Azure struct {
		Subscription   string `yaml:"subscription"`
		ResourceGroup  string `yaml:"resourceGroup"`
		ClusterName    string `yaml:"clusterName"`
	} `yaml:"azure"`
	Kubeconfig struct {
		KubeConfigName string `yaml:"kubeConfigName"`
	} `yaml:"kubeconfig"`
}

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	// todo: fix Short & Long description
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

		var configs []AksConfigData

		for {
			name := getUserInput("Enter name: ")
			subscription := getUserInput("Enter subscription: ")
			resourceGroup := getUserInput("Enter resourceGroup: ")
			clusterName := getUserInput("Enter clusterName: ")
			kubeConfigName := getUserInput("Enter kubeconfig name: ")
	
			config := AksConfigData{
				Name: name,
				Kubeconfig: struct {
					KubeConfigName string `yaml:"kubeConfigName"`
				}{
					KubeConfigName: kubeConfigName,
				},
				Azure: struct {
					Subscription   string `yaml:"subscription"`
					ResourceGroup  string `yaml:"resourceGroup"`
					ClusterName    string `yaml:"clusterName"`
				}{
					Subscription:   subscription,
					ResourceGroup:  resourceGroup,
					ClusterName:    clusterName,
				},
			}
	
			configs = append(configs, config)
			
			fmt.Println("Add another config ? Press 'n' for no, else continue")
			var userInputAddConfig string

			_, err := fmt.Scan(&userInputAddConfig)
			if err != nil {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}

			userInputAddConfig = strings.ToLower(userInputAddConfig)

			if userInputAddConfig == "n" {
				break
			}
		}
	
		yamlData, err := yaml.Marshal(&configs)
		if err != nil {
			fmt.Println("Error marshaling data to YAML:", err)
			os.Exit(1)
		}
	  
		// todo: change the output.yaml to $HOME/.kubeaks.yaml
		// Write the YAML data to the output.yaml file
		err = os.WriteFile("output.yaml", yamlData, 0644)
		if err != nil {
			fmt.Println("Error writing to output.yaml:", err)
			os.Exit(1)
		}
	
		fmt.Println("YAML data written to 'output.yaml'")
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
