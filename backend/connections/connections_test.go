package connections

import (
	"testing"
)

// TestDB -
func TestDB(t *testing.T) {
	// returns an error
	ping := GetDB().DB().Ping()

	if err := ping; err != nil {
		t.Errorf("cannot connect to db : %v ", err)
	}
}
