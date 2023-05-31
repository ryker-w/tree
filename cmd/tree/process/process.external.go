package process

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/tree/internal/conf"
	"github.com/lishimeng/tree/internal/model"
	"github.com/lishimeng/tree/internal/tool"
)

// down_stream
// 接收北向命令, 发送到南向down_stream
func external() {
	qos := byte(conf.Config.Mqtt.Qos)
	subscriber := model.NorthboundDownTopic
	log.Info("subscriber:%s[%b]", subscriber, qos)
	_ = app.GetMqtt().Subscribe(func(topic string, payload []byte) {
		northboundDown(topic, payload)
	}, qos, subscriber)
}

func northboundDown(topic string, payload []byte) {
	m, err := tool.TopicResolver(model.NorthboundDownTpl, topic)
	if err != nil {
		log.Info(err)
		return
	}
	// 转换成内部接口
	var device = m[model.TopicKeyDevice]
	var downMsg model.NorthboundDownStream
	err = json.Unmarshal(payload, &downMsg)
	if err != nil {
		log.Info(err)
		return
	}
	internalTx(device, downMsg)
}

// internalTx 下发数据到设备
func internalTx(device string, payload model.NorthboundDownStream) {

	qos := byte(conf.Config.Mqtt.Qos)
	// find router
	// transform link message
	// send data
	r, err := getRouter(device)
	if err != nil {
		log.Info(err)
		return
	}

	ds := model.DownStream{
		Gateway: r.GatewayCode,
		Device:  r.DeviceCode,
		Channel: r.DeviceChannelCode,
		Data:    payload.Data,
	}

	bs, err := json.Marshal(ds)
	if err != nil {
		log.Info(err)
		return
	}

	topic := fmt.Sprintf(model.SouthboundDownFormat, r.GatewayCode)
	err = app.GetMqtt().Publish(topic, qos, false, bs)
	if err != nil {
		log.Info(err)
		return
	}
}
