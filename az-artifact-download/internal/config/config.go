package config

type AzArtDwnldCfg struct {
	Artifacts []AzArtifactDetail `yaml:"artifacts"`
}

type AzArtifactDetail struct {
	BaseUrl      string `yaml:"baseUrl"`
	Org          string `yaml:"org"`
	Project      string `yaml:"project"`
	Branch       string `yaml:"branch"`
	PipelineDef  int    `yaml:"pipelineDef"`
	FriendlyName string `yaml:"friendlyName"`
	Name         string `yaml:"name"`
}
