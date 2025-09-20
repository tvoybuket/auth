package config

type Settings struct {
	Port            int `config:"env:PORT"`
	ShutdownTimeout int `config:"env:SHUTDOWN_TIMEOUT,default:5"` // seconds
}

type LoggerSettings struct {
	Level       string `config:"env:LOG_LEVEL,default:info"`
	Format      string `config:"env:LOG_FORMAT,default:text"`
	ServiceName string `config:"env:SERVICE_NAME,default:auth"`
}
