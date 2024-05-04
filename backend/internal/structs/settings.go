package structs

type DBSettings struct {
	IsSetupComplete bool   `bson:"is_setup_complete" json:"is_setup_complete"`
	SetupCode       string `bson:"setup_code" json:"setup_code"`
}

type DBUser struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}
