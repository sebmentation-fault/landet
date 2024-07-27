package contactinfo

import (
	"github.com/a-h/templ"
)

type ContactInformation struct {
	Email       string
	EmailURL    templ.SafeURL
	Phonenumber string
}

func Get() *ContactInformation {
	return &ContactInformation{
		Email:       "support@landet.com",
		EmailURL:    "mailto:support@landet.com",
		Phonenumber: "(+44) 7777 777777",
	}
}
