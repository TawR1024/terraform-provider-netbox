package netbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T)  {
	config := Config{
		URL: "my.end.point",
		Token: "superSecureToken",
	}
	err := config.Validate()
	assert.NoError(t, err)
}

func TestValidateNoToken(t *testing.T)  {
	config := &Config{}
	expected := "token must be specified"
	actual := config.Validate()

	assert.EqualError(t, actual,expected)
}

func TestValidateNoUrl(t *testing.T)  {
	config := &Config{
		Token: "superToken",
		URL: "",
	}
	expected := "url must be specified"
	actual := config.Validate()

	assert.EqualError(t, actual,expected)
}