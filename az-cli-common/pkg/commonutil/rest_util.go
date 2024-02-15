package commonutil

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jon-kamis/go-azdev-cli/az-cli-common/pkg/authentication"
	"github.com/jon-kamis/klogger"
)

func MakeRestCall(url string, tok string, respBody any) error {
	method := "util.MakeRestCall"
	klogger.Enter(method)

	data, err := GetRawHttpResponse(url, tok)

	if err != nil {
		klogger.ExitError(method, "failed to read response body: \n%v", err)
		return err
	}

	err = json.Unmarshal(data, &respBody)

	if err != nil {
		klogger.ExitError(method, "failed to parse response body: \n%v", err)
		return err
	}

	klogger.Exit(method)
	return nil
}

func GetRawHttpResponse(url string, tok string) ([]byte, error) {
	method := "util.GetRawHttpResponse"
	klogger.Enter(method)

	//Create Request
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		klogger.ExitError(method, "unexpected error generating http request: \n%v", err)
		return nil, err
	}

	req.Header.Add("Authorization", authentication.GetAzAuthTokenString(tok))

	klogger.Debug(method, "making external call to: %s", url)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		klogger.ExitError(method, "external call returned error: \n%v", err)
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		klogger.ExitError(method, "failed to read response body: \n%v", err)
		return nil, err
	}

	klogger.Exit(method)
	return data, nil
}
