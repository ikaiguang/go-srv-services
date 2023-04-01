package dbv1_0_0_user

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./cmd/migration/v1.0.0/user -test.run=TestMigrator_CreateTableUser
func TestMigrator_CreateTableUser(t *testing.T) {
	err := upHandler.CreateTableUser()
	require.Nil(t, err)
}
