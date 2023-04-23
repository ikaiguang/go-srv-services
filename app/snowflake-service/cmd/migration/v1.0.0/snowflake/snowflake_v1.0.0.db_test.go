package dbv1_0_0_snowflake

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

// go test -v -count=1 ./app/snowflake-service/cmd/migration/v1.0.0/snowflake -test.run=TestMigrator_CreateTableSnowflakeNodeID
func TestMigrator_CreateTableSnowflakeNodeID(t *testing.T) {
	err := upHandler.Upgrade(context.Background())
	require.Nil(t, err)
}
