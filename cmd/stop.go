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

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stops the Z-Cam recording",
	Long:  `This stops recording a video via the Z-Cam. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")
		camUrl, _ := cmd.Flags().GetString("url")
		ledresponse, err := http.Get("http://" + camUrl + "/ctrl/set?led=Off")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		ledresponseData, err := ioutil.ReadAll(ledresponse.Body)
		if err != nil {
			log.Fatal(err)
		}
		_ = ledresponseData
		response, err := http.Get("http://" + camUrl + "/ctrl/rec?action=stop")
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
	rootCmd.AddCommand(stopCmd)

}
