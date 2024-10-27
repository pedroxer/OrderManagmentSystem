package config

type Config struct {
	Postgres      Postgres `json:"postgres"`
	IsDevelopment bool     `env:"DEVELOPMENT"`
	Serv          Api      `json:"api"`
	TokenTTL      string   `json:"token_ttl"`
	Secret        string   `env:"SECRET,notEmpty"`
}
type Postgres struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Db       string `json:"db"`
	User     string `env:"PG_USER,notEmpty"`
	SSLMode  string `json:"sslmode"`
	Password string `env:"PG_PASSWORD,notEmpty"`
}

type Api struct {
	Port int `json:"port"`
}
