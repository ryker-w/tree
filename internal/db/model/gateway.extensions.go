package model

import "github.com/lishimeng/app-starter"

type GatewayExt struct {
	app.Pk
	app.TableChangeInfo
	GatewayCode string `orm:"gateway"`
	Sim         string `orm:"sim"`
}
