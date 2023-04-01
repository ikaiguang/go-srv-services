package dbv1_0_0_org

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./cmd/migration/v1.0.0/org -test.run=TestMigrator_CreateTableOrg
func TestMigrator_CreateTableOrg(t *testing.T) {
	err := upHandler.CreateTableOrg()
	require.Nil(t, err)
}
