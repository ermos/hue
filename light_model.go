package hue

type Light struct {
	Key		string
	Bridge *Bridge
	// Philips Hue API
	State struct {
		On        bool      `json:"on,omitempty"`
		Bri       uint8      `json:"bri,omitempty"`
		Hue       uint16       `json:"hue,omitempty"`
		Sat       uint8       `json:"sat,omitempty"`
		Effect    string    `json:"effect,omitempty"`
		Xy        []float64 `json:"xy,omitempty"`
		Ct        uint16    `json:"ct,omitempty"`
		Alert     string    `json:"alert,omitempty"`
		Colormode string    `json:"colormode"`
		Mode      string    `json:"mode"`
		Reachable bool      `json:"reachable"`
	} `json:"state"`
	Swupdate struct {
		State       string `json:"state"`
		Lastinstall string `json:"lastinstall"`
	} `json:"swupdate"`
	Type             string `json:"type"`
	Name             string `json:"name"`
	Modelid          string `json:"modelid"`
	Manufacturername string `json:"manufacturername"`
	Productname      string `json:"productname"`
	Capabilities     struct {
		Certified bool `json:"certified"`
		Control   struct {
			Mindimlevel    int         `json:"mindimlevel"`
			Maxlumen       int         `json:"maxlumen"`
			Colorgamuttype string      `json:"colorgamuttype"`
			Colorgamut     [][]float64 `json:"colorgamut"`
			Ct             struct {
				Min int `json:"min"`
				Max int `json:"max"`
			} `json:"ct"`
		} `json:"control"`
		Streaming struct {
			Renderer bool `json:"renderer"`
			Proxy    bool `json:"proxy"`
		} `json:"streaming"`
	} `json:"capabilities"`
	Config struct {
		Archetype string `json:"archetype"`
		Function  string `json:"function"`
		Direction string `json:"direction"`
		Startup   struct {
			Mode       string `json:"mode"`
			Configured bool   `json:"configured"`
		} `json:"startup"`
	} `json:"config"`
	Uniqueid   string `json:"uniqueid"`
	Swversion  string `json:"swversion"`
	Swconfigid string `json:"swconfigid,omitempty"`
	Productid  string `json:"productid,omitempty"`
}

type AvailableLight struct {
	ID 	string	`json:"id"`
	Name string	`json:"name"`
}