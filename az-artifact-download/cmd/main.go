package main

import (
	"flag"
	"os"

	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/config"
	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/constants"
	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/models"
	"github.com/jon-kamis/go-azdev-cli/az-artifact-download/internal/rest"
	"github.com/jon-kamis/klogger"
	"gopkg.in/yaml.v2"
)

func main() {
	method := "main"
	klogger.Enter(method)

	var cfn string
	var tok string
	var tags string
	var clean bool
	flag.StringVar(&cfn, constants.ConfigFileEnv, "", "Config file name")
	flag.StringVar(&tok, constants.AuthTokenEnv, "", "Authorization Token")
	flag.StringVar(&tags, constants.TagsEnv, "", "Build Tags")
	flag.BoolVar(&clean, constants.CleanEnv, false, "Clean Download Folder")

	flag.Parse()

	if cfn == "" {
		msg := "Config file path is required!"
		klogger.Error(method, msg)
		panic(msg)
	}

	if tok == "" {
		msg := "AuthorizationToken is required!"
		klogger.Error(method, msg)
		panic(msg)
	}

	var conf config.AzArtDwnldCfg
	fdata, err := os.ReadFile(cfn)

	if err != nil {
		klogger.ExitError(method, "Error reading config file:\n%v", err)
	}

	err = yaml.Unmarshal(fdata, &conf)

	if err != nil {
		klogger.ExitError(method, "Error parsing yaml file:\n%v", err)
		panic("Error parsing yaml file")
	}

	klogger.Info(method, "Successfully parsed yaml file data: \n%v", conf)

	klogger.Info(method, "calling external API for artifact information")

	var bl []models.Build

	aMap := make(map[int]config.AzArtifactDetail)
	for _, art := range conf.Artifacts {
		b := rest.GetBuildArtifacts(art, tags, tok)
		bl = append(bl, b...)
		aMap[art.PipelineDef] = art
	}

	rest.DownloadArtifacts(bl, tok, clean, aMap)

	klogger.Exit(method)
}
