package authentication

import (
	"encoding/base64"
	"fmt"

	"github.com/jon-kamis/klogger"
)

// Function GetBearerToken accepts a string token and returns a base64 encoded bearer token containing its value
func GetAzAuthTokenString(tok string) string {
	method := "authentication.GetAzAuthTokenString"
	klogger.Enter(method)

	tok = ":" + tok
	b := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(tok)))

	klogger.Exit(method)
	return b
}
