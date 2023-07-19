package model

type NorthboundUpStream struct {
	Device  string                 `json:"device,omitempty"`
	Gateway GatewayInfo            `json:"gateway,omitempty"`
	Geo     GeoInfo                `json:"geo,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type NorthboundDownStream struct {
	Device string                 `json:"device,omitempty"`
	Data   map[string]interface{} `json:"data,omitempty"`
}

type GatewayInfo struct {
	Gateway string `json:"gateway,omitempty"`
}

type GeoInfo struct {
	Latitude  float64 `json:"latitude,omitempty"`  // 纬度
	Longitude float64 `json:"longitude,omitempty"` // 经度
}
