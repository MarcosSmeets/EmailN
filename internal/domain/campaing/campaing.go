package campaing

import (
	"time"

	"github.com/rs/xid"
)

type Campaing struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}

type Contact struct {
	Email string
}

func NewCampaing(name string, content string, emails []string) *Campaing {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	return &Campaing{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}
}
