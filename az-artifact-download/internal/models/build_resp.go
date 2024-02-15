package models

type BuildResp struct {
	Count int
	Value []Build
}

type Build struct {
	ID         int
	Definition BuildDef
	Url        string
}

type BuildDef struct {
	Name string
	ID   int
}
