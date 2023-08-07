package model

// UpStream 设备上报
type UpStream struct {
	Device    string                 `json:"device,omitempty"`
	Gateway   string                 `json:"gateway,omitempty"`
	ExtSim    string                 `json:"sim,omitempty"`
	Latitude  any                    `json:"latitude,omitempty"`
	Longitude any                    `json:"longitude,omitempty"`
	Channel   string                 `json:"channel,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// DownStream 下发
type DownStream struct {
	Gateway string                 `json:"gateway,omitempty"`
	Device  string                 `json:"device,omitempty"`
	Channel string                 `json:"channel,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}
