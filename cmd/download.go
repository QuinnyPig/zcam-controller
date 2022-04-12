/*
Copyright © 2022 Corey Quinn corey@lastweekinaws.com

*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/cavaliergopher/grab/v3"
	backoff "github.com/cenkalti/backoff/v4"
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
		timeout, _ := cmd.Flags().GetDuration("timeout")
		retry, _ := cmd.Flags().GetDuration("retry-interval")

		ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
		defer cancel()

		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://"+camUrl+"/DCIM/A001", nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		type Recordings struct {
			Code  int      `json:"code"`
			Desc  string   `json:"desc"`
			Files []string `json:"files"`
		}
		var recordings Recordings
		err = json.Unmarshal(responseData, &recordings)
		if err != nil {
			log.Fatal(err)
		}

		for f := range recordings.Files {
			fmt.Println(recordings.Files[f])
			source := "http://" + camUrl + "/DCIM/A001/" + recordings.Files[f]
			/*This horrible nonsense is because the Z-Cam doesn't seem to have
			a webserver that observes standards. Sometimes it throws bad
			response codes, other times it speaks out of turn. "Fail and retry"
			is sloppy, and yet here we are. I hate this so much.
			*/
			err = backoff.Retry(
				GetOperation(ctx, output, source),
				backoff.WithContext(backoff.NewConstantBackOff(retry), ctx))
			if err != nil {
				log.Fatal(err)
			}
			if delete {
				_, err := http.Get(source + "?act=rm")
				if err != nil {
					log.Printf("Error deleting %s", source)
				}
			}
		}
	},
}

// GetOperation returns a retryable Operation to download a file
func GetOperation(ctx context.Context, output, source string) backoff.Operation {
	return func() error {
		req, err := grab.NewRequest(output, source)
		if err != nil {
			return err
		}
		req = req.WithContext(ctx)

		resp := grab.DefaultClient.Do(req)
		if resp.Err() != nil {
			return resp.Err()
		}

		fmt.Println("Download saved to", resp.Filename)
		return nil
	}
}

func init() {
	downloadCmd.Flags().StringP("output", "o", "", "Directory to download into")
	downloadCmd.Flags().Bool("delete", false, "Delete files after download")
	downloadCmd.Flags().Duration("retry-interval", time.Second*2, "Time between retries")
	downloadCmd.Flags().DurationP("timeout", "t", time.Second*20, "Timeout to cancel the command")
	rootCmd.AddCommand(downloadCmd)
}
