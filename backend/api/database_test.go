package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupDb(t *testing.T) {
	os.Setenv("AUTO_MIGRATE", "Y")
	db, err := SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
