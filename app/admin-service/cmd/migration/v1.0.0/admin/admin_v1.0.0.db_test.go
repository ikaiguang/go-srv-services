package dbv1_0_0_admin

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./cmd/migration/v1.0.0/admin -test.run=TestMigrator_CreateTableAdmin
func TestMigrator_CreateTableAdmin(t *testing.T) {
	err := upHandler.CreateTableAdmin()
	require.Nil(t, err)
}
