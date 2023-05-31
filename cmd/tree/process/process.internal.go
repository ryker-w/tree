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

// up_stream
// 接收南向up消息,发送到北向
func internal() {
	qos := byte(conf.Config.Mqtt.Qos)
	subscriber := model.SouthboundUpTopic
	log.Info("subscriber:%s[%b]", subscriber, qos)
	_ = app.GetMqtt().Subscribe(func(topic string, payload []byte) {
		var err error
		if len(payload) <= 0 {
			return
		}
		var msg model.UpStream

		m, err := tool.TopicResolver(model.SouthboundUpTpl, topic)
		if err != nil {
			log.Info(err)
			return
		}
		var device = m[model.TopicKeyDevice]

		err = json.Unmarshal(payload, &msg)
		if err != nil {
			log.Info(err)
			return
		}
		log.Info(msg)
		if msg.Data == nil {
			msg.Data = make(map[string]any)
		}
		err = updateRouter(device, msg.Gateway, msg.Channel)
		if err != nil {
			log.Info(err)
			return
		}
		var northUp = model.NorthboundDownStream{
			Device: msg.Device,
			Data:   msg.Data,
		}
		topicExternal := fmt.Sprintf(model.NorthboundUpFormat, msg.Device)
		var data []byte
		data, err = json.Marshal(northUp)
		if err != nil {
			log.Info(err)
			return
		}
		err = app.GetMqtt().Publish(topicExternal, qos, false, data)
		if err != nil {
			log.Info(err)
			return
		}
	}, qos, subscriber)
}
