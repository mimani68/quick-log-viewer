package model

type Request struct {
	Query       string `json:"query" xml:"query" form:"query" query:"query"`
	Project     string `json:"project" xml:"project" form:"project" query:"project"`
	Environment string `json:"environment" xml:"environment" form:"environment" query:"environment"`
	Service     string `json:"service" xml:"service" form:"service" query:"service"`
	Workspace   string `json:"workspace" xml:"workspace" form:"workspace" query:"workspace"`
	Time        string `json:"time" xml:"time" form:"time" query:"time"`
}
