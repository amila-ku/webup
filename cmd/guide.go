package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// guideCmd represents the guide command
var guideCmd = &cobra.Command{
	Use:   "guide",
	Short: "Interactive creation of website",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("guide called")
	},
}

func init() {
	rootCmd.AddCommand(guideCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// guideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// guideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
