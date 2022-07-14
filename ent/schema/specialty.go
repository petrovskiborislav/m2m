package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Specialty holds the schema definition for the Specialty entity.
type Specialty struct {
	ent.Schema
}

// Fields of the Specialty.
func (Specialty) Fields() []ent.Field {
	varchar256Type := map[string]string{dialect.Postgres: "VARCHAR(256)"}
	uuidType := map[string]string{dialect.Postgres: "uuid"}

	return []ent.Field{
		field.String("id").MaxLen(16).Default(uuid.NewString()).SchemaType(uuidType),
		field.String("name").Unique().SchemaType(varchar256Type),
	}
}

// Edges of the Specialty.
func (Specialty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organizational_units", OrganizationalUnit.Type).
			Ref("specialties"),
	}
}
