package campaing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "X"
	content  = "Body"
	contacts = []string{"emailtest1@teste.com", "emailtest2@teste.com"}
)

func Test_NewCampain_Create(t *testing.T) {
	assert := assert.New(t)

	campaing := NewCampaing(name, content, contacts)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))
}

func Test_NewCampaing_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaing := NewCampaing(name, content, contacts)

	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_CreatedAtMusBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaing := NewCampaing(name, content, contacts)

	assert.Greater(campaing.CreatedAt, now)
}
