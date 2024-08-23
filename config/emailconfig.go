package config

type EmailConfig struct {
	SmtpHost     string `mapstructure:"smtp-host" json:"smtp-host" yaml:"smtp-host"`
	SmtpPort     string `mapstructure:"smtp-port" json:"smtp-port" yaml:"smtp-port"`
	SmtpEmail    string `mapstructure:"smtp-email" json:"smtp-email" yaml:"smtp-email"`
	SmtpPassword string `mapstructure:"smtp-password" json:"smtp-password" yaml:"smtp-password"`
	MaxClient    int    `mapstructure:"max-client" json:"max-client" yaml:"max-client"`
}
