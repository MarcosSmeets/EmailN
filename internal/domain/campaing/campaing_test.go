package campaing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampain_Create(t *testing.T) {
	assert := assert.New(t)

	name := "X"
	content := "Body"
	contacts := []string{"emailtest1@teste.com", "emailtest2@teste.com"}

	campaing := NewCampaing(name, content, contacts)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))
}

func Test_NewCampaing_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	name := "X"
	content := "Body"
	contacts := []string{"emailtest1@teste.com", "emailtest2@teste.com"}

	campaing := NewCampaing(name, content, contacts)

	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_CreatedAtNotNil(t *testing.T) {
	assert := assert.New(t)

	name := "X"
	content := "Body"
	contacts := []string{"emailtest1@teste.com", "emailtest2@teste.com"}

	campaing := NewCampaing(name, content, contacts)

	assert.NotNil(campaing.CreatedAt)
}
