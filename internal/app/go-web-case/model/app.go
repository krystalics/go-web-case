package model

type App struct {
	Model

	AppId     string `json:"app_id"` //appId
	AppName   string `json:"app_name"`
	OrgId     string `json:"org_id"`
	OrgName   string `json:"org_name"`
	OwnerName string `json:"owner_name"`
	OwnerMail string `json:"owner_mail"`
}
