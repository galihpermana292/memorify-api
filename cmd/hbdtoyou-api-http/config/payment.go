package config

import configlib "hbdtoyou/pkg/config"

type Payment struct {
	EmailSMTP      string                 `yaml:"email_smtp"`
	HostSMTP       string                 `yaml:"host_smtp"`
	PortSMTP       int                    `yaml:"port_smtp"`
	SenderNameSMTP string                 `yaml:"sender_name_smtp"`
	PasswordSMTP   string                 `yaml:"password_smtp"`
	HTTP           map[string]PaymentHTTP `yaml:"http"`
}

type PaymentHTTP struct {
	Timeout configlib.Duration `yaml:"timeout"`
}
