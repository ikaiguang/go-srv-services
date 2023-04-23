package dbv1_0_0_user

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

// go test -v -count=1 ./app/user-service/cmd/migration/v1.0.0/user -test.run=TestMigrator_Upgrade
func TestMigrator_Upgrade(t *testing.T) {
	err := upHandler.Upgrade(context.Background())
	require.Nil(t, err)
}
