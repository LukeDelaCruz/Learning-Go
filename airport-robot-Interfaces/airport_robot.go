package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(visitorName))
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}
func (s Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct{}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}
func (s Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}
