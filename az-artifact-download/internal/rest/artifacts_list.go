package rest

import (
	"fmt"

	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/config"
	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/constants"
	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/models"
	"github.com/jon-kamis/go-azdev-cli/az-cli-common/pkg/commonutil"
	"github.com/jon-kamis/klogger"
)

func GetBuildArtifacts(art config.AzArtifactDetail, tgs string, tok string) []models.Build {
	method := "artifacts_list.GetBuildArtifacts"
	klogger.Enter(method)

	if art.BaseUrl == "" || art.Org == "" || art.Project == "" {
		klogger.ExitError(method, constants.UnableToProcessRestReqMissingParamsErr)
		panic(constants.UnableToProcessRestReqMissingParamsErr)
	}

	url := fmt.Sprintf(constants.FetchBuildsAPI, art.BaseUrl, art.Org, art.Project)
	var ql bool

	// Check if we should search by branchName
	if art.Branch != "" {
		url = fmt.Sprintf("%s?branchName=%s", url, art.Branch)
		ql = true
	}

	//Check if we should search by build tags
	if tgs != "" && ql {
		klogger.Debug(method, "searching for tags: %s", tgs)
		url = fmt.Sprintf("%s&tagFilters=%s", url, tgs)
	} else if tgs != "" {
		klogger.Debug(method, "searching for tags: %s", tgs)
		url = fmt.Sprintf("%s?tagFilters=%s", url, tgs)
	}

	//Check if we should search by pipeline definition
	if art.PipelineDef != 0 && ql {
		klogger.Debug(method, "filtering by pipeline definition: %s", art.PipelineDef)
		url = fmt.Sprintf("%s&definitions=%d", url, art.PipelineDef)
	} else if art.PipelineDef != 0 {
		klogger.Debug(method, "filtering by pipeline definition: %s", art.PipelineDef)
		url = fmt.Sprintf("%s?definitions=%d", url, art.PipelineDef)
	}

	//Filter to only return top result
	if ql {
		url = fmt.Sprintf("%s&queryOrder=queueTimeDescending&$top=1", url)
	} else {
		url = fmt.Sprintf("%s?queryOrder=queueTimeDescending&$top=1", url)
	}

	var respBody models.BuildResp
	err := commonutil.MakeRestCall(url, tok, &respBody)

	if err != nil {
		klogger.ExitError(method, "unexpected error occured making external call:\n%v", err)
		panic("unexpected error occured making external call")
	}

	klogger.Debug(method, "Found %d artifacts", respBody.Count)

	if respBody.Count > 0 {
		klogger.Exit(method)
		return respBody.Value
	}

	klogger.Exit(method)
	return nil
}
