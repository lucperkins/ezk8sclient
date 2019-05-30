package ezk8sclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient(t *testing.T) {
	is := assert.New(t)
	client, err := NewClient()
	is.NoError(err)
	is.NotNil(client)
}
