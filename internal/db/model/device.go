package model

import "github.com/lishimeng/app-starter"

type DeviceRouter struct {
	app.Pk              `json:"-"`
	DeviceCode          string `orm:"column(device)" json:"device_code"`          // 设备Id
	GatewayCode         string `orm:"column(gateway)" json:"gateway_code"`        // 网关Id
	DeviceChannelCode   string `orm:"column(channel)" json:"device_channel_code"` // 设备内部Id
	app.TableChangeInfo `json:"-"`
}
