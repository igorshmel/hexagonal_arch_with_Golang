package common

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filePath, url string) error {
	// Get the data
	if resp, err := http.Get(url); err == nil {
		defer func() {
			if err = resp.Body.Close(); err != nil {
				err = errors.New("DownloadFile error resp.Body.Close(): " + fmt.Sprint(err))
			}
		}()
		// Create the file
		if out, err := os.Create(filePath); err == nil {
			defer func() {
				if err = out.Close(); err != nil {
					err = errors.New("DownloadFile error out.Close(): " + fmt.Sprint(err))
				}
			}()
			// Write the body to file
			if _, err = io.Copy(out, resp.Body); err != nil {
				err = errors.New("DownloadFile error io.Copy: " + fmt.Sprint(err))
				return err
			}

		} else {
			err = errors.New("DownloadFile error os.Create: " + fmt.Sprint(err))
			return err
		}
	} else {
		err = errors.New("DownloadFile error http.Get: " + fmt.Sprint(err))
		return err
	}
	return nil
}
