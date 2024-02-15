package commonutil

import (
	"os"
	"strings"

	"github.com/jon-kamis/klogger"
)

func DownloadFile(fp string, tok string, clean bool, url string) error {
	method := "util.DownloadFile"
	klogger.Enter(method)

	sArr := strings.Split(fp, "/")
	var dir string

	for i, pi := range sArr {
		if i != len(sArr)-1 {

			if i != 0 {
				dir += "/"
			}

			dir += pi

			if i == len(sArr)-2 && clean {
				os.RemoveAll(dir)
			}

			_ = os.Mkdir(dir, os.ModePerm)
		}
	}

	f, err := os.Create(fp)
	if err != nil {
		klogger.ExitError(method, "failed to create filepath:\n%v", err)
		return err
	}

	defer f.Close()

	//Fetch Data
	resp, err := GetRawHttpResponse(url, tok)

	if err != nil {
		klogger.ExitError(method, "failed to download file data:\n%v", err)
		return err
	}

	_, err = f.Write(resp)

	if err != nil {
		klogger.ExitError(method, "error writing file data:\n%v")
		return err
	}

	klogger.Exit(method)
	return nil
}
