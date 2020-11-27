package hue

type Bridge struct {
	IPAddr 	string				`json:"ip_addr"`
	Token 	string				`json:"token"`
	Lights 	map[string]Light 	`json:"lights,omitempty"`
	Groups 	map[string]Group 	`json:"groups,omitempty"`
	Sensors map[string]Sensor 	`json:"sensors,omitempty"`
	Fetch 	BridgeFetch
}

type BridgeFetch struct {
	Bridge *Bridge
}