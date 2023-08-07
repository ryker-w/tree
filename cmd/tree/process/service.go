package process

import (
	"errors"
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/tree/internal/db/model"
)

func updateGatewayExt(gateway string, sim string) (err error) {
	err = app.GetOrm().Transaction(func(context persistence.TxContext) error {

		var extensionInfos []model.GatewayExt
		_, e := context.Context.QueryTable(new(model.GatewayExt)).
			Filter("GatewayCode", gateway).
			All(&extensionInfos)
		if e != nil {
			return e
		}
		var ext model.GatewayExt
		if len(extensionInfos) > 0 {
			ext = extensionInfos[0]
			ext.Sim = sim
			_, e = context.Context.Update(&ext, "Sim")
			if e != nil {
				return e
			}
		} else {
			ext.Sim = sim
			ext.GatewayCode = gateway
			ext.Status = 1
			_, e = context.Context.Insert(&ext)
			if e != nil {
				return e
			}
		}
		return nil
	})
	return
}

// updateRouter 更新router
func updateRouter(device, gateway, channel string) (err error) {

	err = app.GetOrm().Transaction(func(context persistence.TxContext) error {

		var routers []model.DeviceRouter
		_, e := context.Context.QueryTable(new(model.DeviceRouter)).
			Filter("device", device).
			All(&routers)
		if e != nil {
			return e
		}
		var router model.DeviceRouter
		if len(routers) > 0 {
			router = routers[0]
			router.GatewayCode = gateway
			router.DeviceChannelCode = channel
			_, e = context.Context.Update(&router, "gateway", "channel")
			if e != nil {
				return e
			}
		} else {
			router.DeviceCode = device
			router.GatewayCode = gateway
			router.DeviceChannelCode = channel
			router.Status = 1
			_, e = context.Context.Insert(&router)
			if e != nil {
				return e
			}
		}
		return nil
	})
	return
}

func getRouter(device string) (r model.DeviceRouter, err error) {
	var routers []model.DeviceRouter
	_, err = app.GetOrm().Context.QueryTable(new(model.DeviceRouter)).
		Filter("device", device).
		All(&routers)
	if err != nil {
		return
	}
	if len(routers) > 0 {
		r = routers[0]
	} else {
		err = errors.New("not found")
	}
	return
}
