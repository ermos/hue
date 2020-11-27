package hue

type Group struct {
	Name    string        `json:"name"`
	Lights  []string      `json:"lights"`
	Sensors []interface{} `json:"sensors"`
	Type    string        `json:"type"`
	State   struct {
		AllOn bool `json:"all_on"`
		AnyOn bool `json:"any_on"`
	} `json:"state"`
	Recycle bool   `json:"recycle"`
	Class   string `json:"class"`
	Action  struct {
		On        bool      `json:"on"`
		Bri       int       `json:"bri"`
		Hue       int       `json:"hue"`
		Sat       int       `json:"sat"`
		Effect    string    `json:"effect"`
		Xy        []float64 `json:"xy"`
		Ct        int       `json:"ct"`
		Alert     string    `json:"alert"`
		Colormode string    `json:"colormode"`
	} `json:"action"`
}