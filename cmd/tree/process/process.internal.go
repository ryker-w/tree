package process

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/tree/internal/conf"
	"github.com/lishimeng/tree/internal/model"
	"github.com/lishimeng/tree/internal/tool"
	"github.com/pkg/errors"
)

// up_stream
// 接收南向up消息,发送到北向
func internal() {
	qos := byte(conf.Config.Mqtt.Qos)
	subscriber := model.SouthboundUpTopic
	log.Info("subscriber:%s[%b]", subscriber, qos)
	_ = app.GetMqtt().Subscribe(onSouthboundUp, qos, subscriber)
}

func onSouthboundUp(topic string, payload []byte) {
	defer func() {
		if e := recover(); e != nil {
			log.Debug(e)
		}
	}()
	var err error
	qos := byte(conf.Config.Mqtt.Qos)
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
		log.Debug(errors.Wrap(err, "message format must be json"))
		return
	}
	log.Debug("msg:%s", string(payload))
	if msg.Data == nil {
		msg.Data = make(map[string]any)
	}
	if len(msg.ExtSim) > 0 {
		err = updateGatewayExt(msg.Gateway, msg.ExtSim)
		if err != nil {
			log.Info(errors.Wrapf(err, "can't update gateway extension information. %s[%s]", msg.Gateway, msg.ExtSim))
			return
		}
	}
	err = updateRouter(device, msg.Gateway, msg.Channel)
	if err != nil {
		log.Info(errors.Wrap(err, "can't update device router"))
		return
	}
	log.Debug("update device router:%s|%s|%s", device, msg.Gateway, msg.Channel)
	var northUp = model.NorthboundUpStream{
		Device: msg.Device,
		Data:   msg.Data,
	}

	if len(msg.Gateway) > 0 {
		northUp.Gateway.Gateway = msg.Gateway
	}
	latitude, longitude := handleGeo(msg)
	if latitude != 0 && longitude != 0 {
		northUp.Geo.Latitude = latitude
		northUp.Geo.Longitude = longitude
	}
	topicExternal := fmt.Sprintf(model.NorthboundUpFormat, msg.Device)
	var data []byte
	data, err = json.Marshal(northUp)
	if err != nil {
		log.Info(err)
		return
	}
	log.Debug("upload to external:%s", string(data))
	err = app.GetMqtt().Publish(topicExternal, qos, false, data)
	if err != nil {
		log.Info(err)
		return
	}
}

func handleGeo(src model.UpStream) (latitude float64, longitude float64) {
	latitude = tool.ConvertGeo(src.Latitude)
	longitude = tool.ConvertGeo(src.Longitude)
	return
}
