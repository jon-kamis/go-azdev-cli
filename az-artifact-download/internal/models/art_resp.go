package models

type BuildArtifact struct {
	ID       int
	Name     string
	Resource ArtifactResource
}

type ArtifactResource struct {
	DownloadTicket string
	DownloadUrl    string
	Url            string
}
