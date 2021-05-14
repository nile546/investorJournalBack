package mailer

type Mailer interface {
	Send(string) error
}
