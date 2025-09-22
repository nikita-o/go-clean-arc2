package config

type Config struct {
	ConfigServer
	ConfigAuth
	ConfigDatabase
	ConfigSmtp
	ConfigRedis
}

type ConfigServer struct {
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
}

type ConfigAuth struct {
	AccessTokenExpiryMinute int    `mapstructure:"ACCESS_TOKEN_EXPIRY_MINUTE"`
	RefreshTokenExpiryHour  int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret       string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret      string `mapstructure:"REFRESH_TOKEN_SECRET"`
	SessionExpiryMinute     int    `mapstructure:"SESSION_EXPIRY_MINUTE"`
}

type ConfigDatabase struct {
	DbUrl string `mapstructure:"DATABASE_URL"`
}

type ConfigRedis struct {
	RedisDatabaseUrl string `mapstructure:"REDIS_DATABASE_URL"`
}

type ConfigClickHouse struct {
	ClickHouseDatabaseUrl string `mapstructure:"CLICK_HOUSE_DATABASE_URL"`
}

type ConfigSmtp struct {
	SmtpServer   string `mapstructure:"SMTP_SERVER"`
	SmtpPort     int    `mapstructure:"SMTP_PORT"`
	SmtpLogin    string `mapstructure:"SMTP_LOGIN"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
	SmtpEmail    string `mapstructure:"SMTP_EMAIL"`
}
