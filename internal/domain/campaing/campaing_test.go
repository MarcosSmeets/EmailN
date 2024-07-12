package campaing

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "teste X"
	content  = "Body teste"
	contacts = []string{"emailtest1@teste.com", "emailtest2@teste.com"}
	fake     = faker.New()
)

func Test_NewCampain_Create(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))
}

func Test_NewCampaing_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_CreatedAtMusBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.Greater(campaing.CreatedAt, now)
}

func Test_NewCampaing_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing("", content, contacts)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaing_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(fake.Lorem().Text(100), content, contacts)

	assert.Equal("name is required with max 50", err.Error())
}

func Test_NewCampaing_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, "", contacts)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaing_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, fake.Lorem().Text(1050), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaing_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, fake.Lorem().Text(100), []string{})

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaing_MustValidateContactsEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, fake.Lorem().Text(100), []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}
