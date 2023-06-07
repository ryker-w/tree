package tool

import (
	"github.com/lishimeng/go-log"
	"strconv"
)

func ConvertGeo(src any) (d float64) {
	var err error
	if src == nil {
		return
	}
	switch src.(type) {
	case string:
		d, err = strconv.ParseFloat(src.(string), 64)
		if err != nil {
			log.Debug(err)
			d = 0
		}
	case float64:
		d = src.(float64)
	case int64:
		var tmp = src.(int64)
		d = float64(tmp)
	default:
		d = 0
	}
	return
}
