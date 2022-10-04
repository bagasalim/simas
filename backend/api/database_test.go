package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupDb(t *testing.T) {
	db, err := SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
