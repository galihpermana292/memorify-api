package service

import (
	"errors"
	"hbdtoyou/internal/auth"
	"hbdtoyou/internal/content"
	"hbdtoyou/pkg/provider/email"
	"time"
)

// Followings are the known error returned from service.
var (
	errMissingMandatoryConfig = errors.New("missing mandatory config")
)

var timeFormat = "Monday, 02 Jan 2006 at 15:04"

// service implements subject.Service.
type service struct {
	pgStore PGStore
	mail    *email.Gomail
	user    auth.Service
	content content.Service
	config  Config
	timeNow func() time.Time
}

// Config denotes service configuration
//
// Adding a new field should also add the corresponding default
// value in getDefaultConfig().
type Config struct {
	SMTPEmail    string
	SMTPHost     string
	SMTPPort     int
	SMTPSender   string
	SMTPPassword string
}

// New creates a new service.
func New(pgStore PGStore, user auth.Service, content content.Service, options ...Option) (*service, error) {
	s := &service{
		pgStore: pgStore,
		user:    user,
		content: content,
		config:  Config{},
		timeNow: time.Now,
	}

	// apply options
	for _, opt := range options {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	s.mail = email.NewMailClient(s.config.SMTPHost, s.config.SMTPEmail, s.config.SMTPPassword, s.config.SMTPPort)

	// verify mandatory config
	if s.config.SMTPEmail == "" || s.config.SMTPHost == "" || s.config.SMTPPort == 0 || s.config.SMTPSender == "" || s.config.SMTPPassword == "" {
		return nil, errMissingMandatoryConfig
	}

	return s, nil
}

// Option controls the behavior of service.
type Option func(*service) error

// WithConfig returns Option to set service configuration.
func WithConfig(config Config) Option {
	return func(s *service) error {
		if config.SMTPEmail != "" {
			s.config.SMTPEmail = config.SMTPEmail
		}
		if config.SMTPHost != "" {
			s.config.SMTPHost = config.SMTPHost
		}
		if config.SMTPPort > 0 {
			s.config.SMTPPort = config.SMTPPort
		}
		if config.SMTPSender != "" {
			s.config.SMTPSender = config.SMTPSender
		}
		if config.SMTPPassword != "" {
			s.config.SMTPPassword = config.SMTPPassword
		}
		return nil
	}
}
