/*
Copyright © 2022 Corey Quinn corey@lastweekinaws.com

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the Z-Cam recording",
	Long:  `This starts recording a video via the Z-Cam. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		camUrl, _ := cmd.Flags().GetString("url")
		response, err := http.Get("http://" + camUrl + "/ctrl/rec?action=start")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
		fmt.Println("Print: " + camUrl)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

}
