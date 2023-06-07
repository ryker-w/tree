package model

import (
	"github.com/lishimeng/tree/internal/tool"
)

// UpStream 设备上报
type UpStream struct {
	Device    string                 `json:"device,omitempty"`
	Gateway   string                 `json:"gateway,omitempty"`
	Latitude  any                    `json:"latitude,omitempty"`
	Longitude any                    `json:"longitude,omitempty"`
	Channel   string                 `json:"channel,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// DownStream 下发
type DownStream struct {
	Gateway string                 `json:"gateway,omitempty"`
	Device  string                 `json:"device,omitempty"`
	Channel string                 `json:"channel,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type NorthboundUpStream struct {
	Device  string                 `json:"device,omitempty"`
	Gateway GatewayInfo            `json:"gateway,omitempty"`
	Geo     GeoInfo                `json:"geo,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type GatewayInfo struct {
	Gateway string `json:"gateway,omitempty"`
}

type GeoInfo struct {
	Latitude  float64 `json:"latitude,omitempty"`  // 纬度
	Longitude float64 `json:"longitude,omitempty"` // 经度
}

type NorthboundDownStream struct {
	Device string                 `json:"device,omitempty"`
	Data   map[string]interface{} `json:"data,omitempty"`
}

const (
	SouthboundDownFormat = "tree/%s/down"
	SouthboundUpFormat   = "tree/%s/up"

	NorthboundDownFormat = "tree/external/%s/down"
	NorthboundUpFormat   = "tree/external/%s/up"

	TopicKeyDevice = "Device"
)

var (
	SouthboundUpTpl   string
	SouthboundUpTopic string

	NorthboundDownTpl   string
	NorthboundDownTopic string
)

func init() {
	SouthboundUpTpl = tool.TopicBuilder(tool.BuilderOption{
		Share: false,
		Tpl:   true,
	}, SouthboundUpFormat, TopicKeyDevice)
	SouthboundUpTopic = tool.TopicBuilder(tool.BuilderOption{
		Share: true,
		Tpl:   false,
	}, SouthboundUpFormat, "+")

	NorthboundDownTpl = tool.TopicBuilder(tool.BuilderOption{
		Share: false,
		Tpl:   true,
	}, NorthboundDownFormat, TopicKeyDevice)
	NorthboundDownTopic = tool.TopicBuilder(tool.BuilderOption{
		Share: true,
		Tpl:   false,
	}, NorthboundDownFormat, "+")
}
