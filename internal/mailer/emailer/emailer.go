package emailer

import "net/smtp"

type Config struct {
	Login  string
	Pass   string
	Host   string
	Port   string
	Sender string
}

type Emailer struct {
	config *Config
}

func New(c *Config) *Emailer {
	return &Emailer{
		config: c,
	}
}

func (e *Emailer) Send(recipients []string, msg string) error {

	auth := smtp.PlainAuth(
		"",
		e.config.Login,
		e.config.Pass,
		e.config.Host,
	)

	addr := e.config.Host + ":" + e.config.Port

	err := smtp.SendMail(
		addr,
		auth,
		e.config.Sender,
		recipients,
		[]byte(msg),
	)

	if err != nil {
		return err
	}

	return nil
}
