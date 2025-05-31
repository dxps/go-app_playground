package config

type Config struct {
	Db struct {
		Driver       string `env:"DB_DRIVER, default=postgres"`
		DSN          string `env:"DB_DSN, default=postgres://data_portal:data_portal@localhost:5456/data_portal?sslmode=disable&timezone=Europe/Bucharest"`
		MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS, default=10"`
		MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS, default=2"`
		MaxIdleTime  string `env:"DB_MAX_IDLE_TIME, default=3m"`
	}
	Servers struct {
		FrontendPort int `env:"FRONTEND_PORT, default=8001"`
		BackendPort  int `env:"BACKEND_PORT, default=8002"`
	}
}
