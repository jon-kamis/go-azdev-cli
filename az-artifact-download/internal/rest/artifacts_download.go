package rest

import (
	"fmt"

	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/config"
	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/constants"
	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/models"
	"github.com/jon-kamis/go-azdev-cli/az-cli-common/pkg/commonutil"
	"github.com/jon-kamis/klogger"
)

func DownloadArtifacts(bl []models.Build, tok string, clean bool, artMap map[int]config.AzArtifactDetail) {
	method := "artifacts_download.DownloadArtifacts"
	klogger.Enter(method)

	for _, b := range bl {

		aUrl := fmt.Sprintf("%s/%s?artifactName=%s", b.Url, constants.ArtifactsApiSuffix, artMap[b.Definition.ID].Name)

		var art models.BuildArtifact
		err := commonutil.MakeRestCall(aUrl, tok, &art)

		if err != nil {
			klogger.ExitError(method, "unexpected error calling rest API:\n%v", err)
			panic("unexpected error calling rest API")
		}

		fp := fmt.Sprintf("downloads/%s/Drop.zip", artMap[b.Definition.ID].FriendlyName)
		err = commonutil.DownloadFile(fp, tok, clean, art.Resource.DownloadUrl)

		if err != nil {
			klogger.ExitError(method, "unexpected error downloading artifact:\n%v", err)
		}

	}

	klogger.Exit(method)
}
