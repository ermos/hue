package hue

type Error []struct {
	Error struct {
		Type 		int		`json:"type"`
		Address 	string 	`json:"address"`
		Description string 	`json:"description"`
	} `json:"error"`
}