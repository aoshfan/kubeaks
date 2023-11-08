/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"kubeaks/cmd/utils"
	"log"
	"os"
	"os/exec"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	SubscriptionId, ResourceGroup, ClusterName string
)

type Config struct {
	Name     string `yaml:"name"`
	Azure    AzureConfig
	KubeConf KubeConfig `yaml:"kubeconfig"`
}

type AzureConfig struct {
	Subscription  string `yaml:"subscription"`
	ResourceGroup string `yaml:"resourceGroup"`
	ClusterName   string `yaml:"clusterName"`
}

type KubeConfig struct {
	KubeConfigName string `yaml:"kubeConfigName"`
}

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switch called")
		kubeloginPath := utils.CommandExists("kubelogin")

		clustername, _ := cmd.Flags().GetString("clustername")
		if clustername == "" {
			fmt.Println("Please specify localFlag")
			os.Exit(1)
		}

		yamlData, err := os.ReadFile("output.yaml")
		if err != nil {
			log.Fatalf("Failed to read YAML file: %v", err)
		}

		// Define a slice to hold multiple Config entries
		var configs []Config

		// Unmarshal the YAML into the slice of Config structs
		if err := yaml.Unmarshal(yamlData, &configs); err != nil {
			log.Fatalf("Failed to unmarshal YAML: %v", err)
		}

		// Iterate over the Config entries
		i := 0
		for _, config := range configs {

			totalConfigs := len(configs)
			fmt.Println(clustername)
			fmt.Println(config.Name)
			if clustername == config.Name {
				fmt.Printf("Name: %s\n", config.Name)
				fmt.Printf("Azure Subscription: %s\n", config.Azure.Subscription)
				fmt.Printf("Azure ResourceGroup: %s\n", config.Azure.ResourceGroup)
				fmt.Printf("Azure ClusterName: %s\n", config.Azure.ClusterName)
				fmt.Printf("KubeConfig Name: %s\n", config.KubeConf.KubeConfigName)
				fmt.Println("-----")
				SubscriptionId = config.Azure.Subscription
				ResourceGroup = config.Azure.ResourceGroup
				ClusterName = config.Azure.ClusterName
				break
			} else {
				config.Name = ""
				i += 1
				fmt.Println(i)
				if totalConfigs == i {
					fmt.Println("Cluster Name not found")
					os.Exit(1)
				}
			}

		}

		cred, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			log.Fatalf("failed to obtain a credential: %v", err)
		}

		ctx := context.Background()
		clientFactory, err := armcontainerservice.NewClientFactory(SubscriptionId, cred, nil)
		if err != nil {
			log.Fatalf("failed to create client: %v", err)
		}
		res, err := clientFactory.NewManagedClustersClient().ListClusterUserCredentials(ctx, ResourceGroup, ClusterName, &armcontainerservice.ManagedClustersClientListClusterUserCredentialsOptions{ServerFqdn: nil,
			Format: nil,
		})
		if err != nil {
			log.Fatalf("failed to finish the request: %v", err)
		}

		kubeconfig := string(res.Kubeconfigs[0].Value)

		f, err := os.Create("kubeconfig.yaml")
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(kubeconfig)

		if err != nil {
			log.Fatal(err2)
		}

		fmt.Println(kubeconfig)

		// Run kubelogin
		fmt.Println(kubeloginPath)

		// todo: kubeconfig path take from output.yaml
		// todo: try using Stdin, Stdout, Stderr :D
		cmdKubelogin := exec.Command("kubelogin", "convert-kubeconfig", "--kubeconfig", "kubeconfig.yaml", "-l", "azurecli")
		// cmdKubelogin.Stdin = os.Stdin
		// cmdKubelogin.Stdout = os.Stdout
		// cmdKubelogin.Stderr = os.Stderr

		if err := cmdKubelogin.Run(); err != nil {
			panic(err)
		}
		fmt.Println("Success")

	},
}

func init() {
	switchCmd.Flags().String("clustername", "", "a local string flag")
	rootCmd.AddCommand(switchCmd)
	// commandExists("kubelogin")

}
