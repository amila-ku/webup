package cmd

import (
	"context"
	"fmt"

	"github.com/amila-ku/webup/aws"
	"github.com/spf13/cobra"
)

var webSiteName string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new website in AWS",
	Long:  `Creates a new website for the given dns name.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		err := aws.MakeWebResources(context.TODO(), webSiteName)
		if err != nil {
			fmt.Println("failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().StringVarP(&webSiteName, "domain-name", "n", "", "Web site name")
}
