package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/app-starter/mqtt"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/tree/cmd/tree/process"
	"github.com/lishimeng/tree/internal/api"
	"github.com/lishimeng/tree/internal/conf"
	"github.com/lishimeng/tree/internal/db/model"
	"time"
)
import _ "github.com/lib/pq"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.SetLevelAll(log.DEBUG)

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 500)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&conf.Config, func(loader etc.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}
		if len(conf.Config.Web.Listen) == 0 {
			conf.Config.Web.Listen = ":80"
		}
		dbConfig := persistence.PostgresConfig{
			UserName:  conf.Config.Db.User,
			Password:  conf.Config.Db.Password,
			Host:      conf.Config.Db.Host,
			Port:      conf.Config.Db.Port,
			DbName:    conf.Config.Db.Database,
			InitDb:    true,
			AliasName: "default",
			SSL:       conf.Config.Db.Ssl,
		}

		cfgMqtt := conf.Config.Mqtt
		builder.
			EnableDatabase(dbConfig.Build(),
				model.Tables()...).
			//SetWebLogLevel("debug").
			EnableMqtt(mqtt.WithAuth(cfgMqtt.Username, cfgMqtt.Password),
				mqtt.WithBroker(cfgMqtt.Broker),
				mqtt.WithRandomClientId(),
				mqtt.WithOnConnectHandler(process.AfterConnect),
				mqtt.WithOnLostHandler(process.OnLostConnection)).
			EnableWeb(conf.Config.Web.Listen, api.Route).
			PrintVersion()
		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}
