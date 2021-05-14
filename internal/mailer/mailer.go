package mailer

type Mailer interface {
	Send(recipients []string, msg string) error
}
