package model

import "github.com/lishimeng/app-starter"

type DeviceRouter struct {
	app.Pk
	DeviceCode        string `orm:"column(device)"`  // 设备Id
	GatewayCode       string `orm:"column(gateway)"` // 网关Id
	DeviceChannelCode string `orm:"column(channel)"` // 设备内部Id
	app.TableChangeInfo
}
