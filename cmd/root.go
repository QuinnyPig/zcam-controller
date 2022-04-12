/*
Copyright © 2022 Corey Quinn corey@lastweekinaws.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zcam-controller",
	Short: "Controls Z-CAMs via API",
	Long:  `zcam-controller uses the z-cam API to control various aspects of the Z-CAM.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zcam-controller.yaml)")
	var camUrl string
	rootCmd.PersistentFlags().StringVarP(&camUrl, "url", "u", "", "URL for the Z-CAM")
	rootCmd.MarkFlagRequired("url")

}
