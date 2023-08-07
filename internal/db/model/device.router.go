package model

import "github.com/lishimeng/app-starter"

// DeviceRouter 设备路由表, 每次上报后, 更新相应设备的网关编号与网络地址
type DeviceRouter struct {
	app.Pk
	DeviceCode        string `orm:"column(device)"`  // 设备Id
	GatewayCode       string `orm:"column(gateway)"` // 网关Id
	DeviceChannelCode string `orm:"column(channel)"` // 设备内部Id
	app.TableChangeInfo
}
