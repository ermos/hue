package hue

type Sensor struct {
	State struct {
		Daylight    bool   `json:"daylight,omitempty"`
		Buttonevent int    `json:"buttonevent,omitempty"`
		Lastupdated string `json:"lastupdated"`
	} `json:"state"`
	Swupdate struct {
		State       string `json:"state"`
		Lastinstall string `json:"lastinstall"`
	} `json:"swupdate,omitempty"`
	Config struct {
		On            	bool 			`json:"on"`
		Configured    	bool 			`json:"configured,omitempty"`
		Sunriseoffset 	int  			`json:"sunriseoffset,omitempty"`
		Sunsetoffset  	int  			`json:"sunsetoffset,omitempty"`
		Battery   		int           	`json:"battery,omitempty"`
		Reachable 		bool          	`json:"reachable,omitempty"`
		Pending   		[]interface{} 	`json:"pending,omitempty"`
	} `json:"config"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	Modelid          string `json:"modelid"`
	Manufacturername string `json:"manufacturername"`
	Productname      string `json:"productname,omitempty"`
	Diversityid      string `json:"diversityid,omitempty"`
	Swversion        string `json:"swversion"`
	Uniqueid         string `json:"uniqueid,omitempty"`
	Capabilities     struct {
		Certified bool `json:"certified"`
		Primary   bool `json:"primary"`
		Inputs    []struct {
			Repeatintervals []int `json:"repeatintervals"`
			Events          []struct {
				Buttonevent int    `json:"buttonevent"`
				Eventtype   string `json:"eventtype"`
			} `json:"events"`
		} `json:"inputs"`
	} `json:"capabilities,omitempty"`
}