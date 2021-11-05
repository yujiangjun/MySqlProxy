package dto
type DataSource struct {
	Host string `json:"host"`
	Port int `json:"port"`
	Schema string `json:"schema"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
