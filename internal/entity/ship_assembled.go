package entity

type ShipAssembled struct {
	EventUUID    string `json:"event_uuid"`
	OrderUUID    string `json:"order_uuid"`
	UserUUID     string `json:"user_uuid"`
	BuildTimeSec int64  `json:"build_time_sec"`
}
