package model

import (
	"github.com/lishimeng/tree/internal/tool"
)

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
