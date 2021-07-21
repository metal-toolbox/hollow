package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"

	"go.metalkube.net/hollow/internal/db"
)

func TestCreateAttributes(t *testing.T) {
	s := db.DatabaseTest(t)

	var testCases = []struct {
		testName    string
		a           *db.Attributes
		expectError bool
		errorMsg    string
	}{
		{"missing namespace", &db.Attributes{}, true, "validation failed: namespace is a required attributes attribute"},
		{"happy path", &db.Attributes{ServerID: &db.FixtureServerDory.ID, Namespace: "hollow.test", Data: datatypes.JSON([]byte(`{"value": "set"}`))}, false, ""},
	}

	for _, tt := range testCases {
		err := s.CreateAttributes(tt.a)

		if tt.expectError {
			assert.Error(t, err, tt.testName)
			assert.Contains(t, err.Error(), tt.errorMsg)
		} else {
			assert.NoError(t, err, tt.testName)
		}
	}
}
