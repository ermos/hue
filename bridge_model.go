package hue

type Bridge struct {
	IPAddr 	string				`json:"ip_addr"`
	Token 	string				`json:"token"`
	Lights 	map[string]*Light 	`json:"lights,omitempty"`
	Groups 	map[string]*Group 	`json:"groups,omitempty"`
	Sensors map[string]*Sensor 	`json:"sensors,omitempty"`
	Fetch 	BridgeFetch
	Config 	*BridgeConfig
}

type BridgeFetch struct {
	Bridge *Bridge
}

type BridgeConfig struct {
	Name             string `json:"name"`
	Zigbeechannel    int    `json:"zigbeechannel"`
	Bridgeid         string `json:"bridgeid"`
	Mac              string `json:"mac"`
	Dhcp             bool   `json:"dhcp"`
	Ipaddress        string `json:"ipaddress"`
	Netmask          string `json:"netmask"`
	Gateway          string `json:"gateway"`
	Proxyaddress     string `json:"proxyaddress"`
	Proxyport        int    `json:"proxyport"`
	UTC              string `json:"UTC"`
	Localtime        string `json:"localtime"`
	Timezone         string `json:"timezone"`
	Modelid          string `json:"modelid"`
	Datastoreversion string `json:"datastoreversion"`
	Swversion        string `json:"swversion"`
	Apiversion       string `json:"apiversion"`
	Swupdate         struct {
		Updatestate    int  `json:"updatestate"`
		Checkforupdate bool `json:"checkforupdate"`
		Devicetypes    struct {
			Bridge  bool          `json:"bridge"`
			Lights  []interface{} `json:"lights"`
			Sensors []interface{} `json:"sensors"`
		} `json:"devicetypes"`
		URL    string `json:"url"`
		Text   string `json:"text"`
		Notify bool   `json:"notify"`
	} `json:"swupdate"`
	Swupdate2 struct {
		Checkforupdate bool   `json:"checkforupdate"`
		Lastchange     string `json:"lastchange"`
		Bridge         struct {
			State       string `json:"state"`
			Lastinstall string `json:"lastinstall"`
		} `json:"bridge"`
		State       string `json:"state"`
		Autoinstall struct {
			Updatetime string `json:"updatetime"`
			On         bool   `json:"on"`
		} `json:"autoinstall"`
	} `json:"swupdate2"`
	Linkbutton       bool   `json:"linkbutton"`
	Portalservices   bool   `json:"portalservices"`
	Portalconnection string `json:"portalconnection"`
	Portalstate      struct {
		Signedon      bool   `json:"signedon"`
		Incoming      bool   `json:"incoming"`
		Outgoing      bool   `json:"outgoing"`
		Communication string `json:"communication"`
	} `json:"portalstate"`
	Internetservices struct {
		Internet     string `json:"internet"`
		Remoteaccess string `json:"remoteaccess"`
		Time         string `json:"time"`
		Swupdate     string `json:"swupdate"`
	} `json:"internetservices"`
	Factorynew       bool        `json:"factorynew"`
	Replacesbridgeid interface{} `json:"replacesbridgeid"`
	Backup           struct {
		Status    string `json:"status"`
		Errorcode int    `json:"errorcode"`
	} `json:"backup"`
	Starterkitid string `json:"starterkitid"`
	Whitelist    map[string]struct {
		LastUseDate string `json:"last use date"`
		CreateDate  string `json:"create date"`
		Name        string `json:"name"`
	} `json:"whitelist"`
}