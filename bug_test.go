package bug

import (
	"context"
	"fmt"
	"testing"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/enttest"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port))
			defer client.Close()
			test(t, client)
		})
	}
}

func test(t *testing.T, client *ent.Client) {
	ctx := context.Background()

	migrationOptions := []schema.MigrateOption{
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	err := client.Debug().Schema.Create(ctx, migrationOptions...)
	assert.NoError(t, err)
}
