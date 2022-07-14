package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// OrganizationalUnit holds the schema definition for the OrganizationalUnit entity.
type OrganizationalUnit struct {
	ent.Schema
}

// Fields of the OrganizationalUnit.
func (OrganizationalUnit) Fields() []ent.Field {
	timestamptzType := map[string]string{dialect.Postgres: "TIMESTAMPTZ"}
	varchar256Type := map[string]string{dialect.Postgres: "VARCHAR(256)"}
	uuidType := map[string]string{dialect.Postgres: "uuid"}

	return []ent.Field{
		field.String("id").MaxLen(16).Default(uuid.NewString()).SchemaType(uuidType),
		field.String("display_name").SchemaType(varchar256Type).Default(""),
		field.Time("created_at").SchemaType(timestamptzType).Default(time.Now),
		field.Time("deleted_at").SchemaType(timestamptzType).Optional(),
		field.String("parent_id").SchemaType(uuidType).Optional(),
		field.JSON("metadata", []string{}).Default([]string{}),
		field.String("type_id").SchemaType(uuidType),
	}
}

// Edges of the OrganizationalUnit.
func (OrganizationalUnit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", OrganizationalUnit.Type).
			From("parent").Field("parent_id").Unique(),
		edge.From("organizational_unit_type_id", OrganizationalUnitType.Type).
			Ref("organizational_units").Field("type_id").Unique().Required(),
		edge.To("specialties", Specialty.Type),
	}
}
