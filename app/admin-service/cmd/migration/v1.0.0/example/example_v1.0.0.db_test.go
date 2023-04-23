package dbv1_0_0_example

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

// go test -v -count=1 ./app/admin-service/cmd/migration/v1.0.0/example -test.run=TestMigrator_Upgrade
func TestMigrator_Upgrade(t *testing.T) {
	err := upHandler.Upgrade(context.Background())
	require.Nil(t, err)
}
