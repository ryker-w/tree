package process

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lishimeng/go-log"
)

type Tx struct {
	Gateway string
	Device  string
	Channel string
	Data    map[string]interface{}
}

func AfterConnect(_ mqtt.Client) {

	log.Info("mqtt connect established")

	internal()
	external()
}

func OnLostConnection(_ mqtt.Client, err error) {
	log.Info("mqtt connection lost")
	log.Info(err)
}
