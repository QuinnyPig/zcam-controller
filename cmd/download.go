/*
Copyright © 2022 Corey Quinn corey@lastweekinaws.com

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cavaliergopher/grab/v3"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "downloads the Z-Cam recordings",
	Long:  `This downloads videos off of the Z-Cam. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")
		camUrl, _ := cmd.Flags().GetString("url")
		output, _ := cmd.Flags().GetString("output")
		delete, _ := cmd.Flags().GetBool("delete")
		files, err := http.Get("http://" + camUrl + "/DCIM/A001")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(files.Body)
		if err != nil {
			log.Fatal(err)
		}
		type Recordings struct {
			Code  int      `json:"code"`
			Desc  string   `json:"desc"`
			Files []string `json:"files"`
		}
		var recordings Recordings
		json.Unmarshal([]byte(responseData), &recordings)
		for f := range recordings.Files {
			fmt.Println(recordings.Files[f])
			source := "http://" + camUrl + "/DCIM/A001/" + recordings.Files[f]
			/*This horrible nonsense is because the Z-Cam doesn't seem to have
			a webserver that observes standards. Sometimes it throws bad
			response codes, other times it speaks out of turn. "Fail and retry"
			is sloppy, and yet here we are. I hate this so much.
			*/
			for i := 0; i < 10; i++ {

				resp, err := grab.Get(output, source)
				if err != nil {
					if i >= 9 {
						log.Fatal(err)
					}
				} else {
					i = 10

					if delete {
						http.Get(source + "?act=rm")
					}
					fmt.Println("Download saved to", resp.Filename)
					break
				}
			}
		}
	},
}

func init() {
	downloadCmd.Flags().StringP("output", "o", "", "Directory to download into")
	downloadCmd.Flags().Bool("delete", false, "Delete files after download")
	rootCmd.AddCommand(downloadCmd)

}
