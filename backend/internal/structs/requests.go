package structs

// Create user struct
type RequestSetupUser struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	SetupCode       string `json:"setup_code"`
}
