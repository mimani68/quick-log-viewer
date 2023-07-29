package model

type SubmitNewData struct {
	Data        interface{} `json:"data" xml:"data" form:"data" query:"data"`
	Project     string      `json:"project" xml:"project" form:"project" query:"project"`
	Environment string      `json:"environment" xml:"environment" form:"environment" query:"environment"`
	Service     string      `json:"service" xml:"service" form:"service" query:"service"`
	Workspace   string      `json:"workspace" xml:"workspace" form:"workspace" query:"workspace"`
	Time        string      `json:"time" xml:"time" form:"time" query:"time"`
}
