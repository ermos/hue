package hue

type tokenResponse []struct {
	Success struct {
		Username string `json:"username"`
	} `json:"success"`
}
