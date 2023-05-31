package process

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/tree/internal/etc"
)

type Tx struct {
	Gateway string
	Device  string
	Channel string
	Data    map[string]interface{}
}

func AfterConnect(client mqtt.Client) {

	log.Info("mqtt connect established")

	internal()
	external()
}

func OnLostConnection(_ mqtt.Client, err error) {
	log.Info("mqtt connection lost")
	log.Info(err)
}

func Command(payload Tx) (err error) {

	// TODO get gateway sn from router
	to := fmt.Sprintf("prod/%s/tx", payload.Gateway)
	var data []byte
	data, err = json.Marshal(payload)
	if err != nil {
		return
	}
	err = app.GetMqtt().Publish(to, byte(etc.Config.Mqtt.Qos), false, data)
	return
}
