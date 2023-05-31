package model

import (
	"fmt"
	"github.com/lishimeng/tree/internal/tool"
)

// UpStream 设备上报
type UpStream struct {
	Device    string                 `json:"device,omitempty"`
	Gateway   string                 `json:"gateway,omitempty"`
	Latitude  float64                `json:"latitude,omitempty"`
	Longitude float64                `json:"longitude,omitempty"`
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
	SouthboundDownTpl   string
	SouthboundUpTpl     string
	SouthboundDownTopic string
	SouthboundUpTopic   string

	NorthboundDownTpl   string
	NorthboundUpTpl     string
	NorthboundDownTopic string
	NorthboundUpTopic   string
)

func init() {
	SouthboundDownTpl = tool.TopicBuilder(SouthboundDownFormat, TopicKeyDevice)
	SouthboundUpTpl = tool.TopicBuilder(SouthboundUpFormat, TopicKeyDevice)
	SouthboundDownTopic = fmt.Sprintf(SouthboundDownFormat, "+")
	SouthboundUpTopic = fmt.Sprintf(SouthboundUpFormat, "+")

	NorthboundDownTpl = tool.TopicBuilder(NorthboundDownFormat, TopicKeyDevice)
	NorthboundUpTpl = tool.TopicBuilder(NorthboundUpFormat, TopicKeyDevice)
	NorthboundDownTopic = fmt.Sprintf(NorthboundDownFormat, "+")
	NorthboundUpTopic = fmt.Sprintf(NorthboundUpFormat, "+")
}
