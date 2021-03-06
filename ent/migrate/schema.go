// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrganizationalUnitsColumns holds the columns for the "organizational_units" table.
	OrganizationalUnitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 16, SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "display_name", Type: field.TypeString, Default: "", SchemaType: map[string]string{"postgres": "VARCHAR(256)"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "TIMESTAMPTZ"}},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "TIMESTAMPTZ"}},
		{Name: "metadata", Type: field.TypeJSON},
		{Name: "parent_id", Type: field.TypeString, Nullable: true, Size: 16, SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "type_id", Type: field.TypeString, Size: 16, SchemaType: map[string]string{"postgres": "uuid"}},
	}
	// OrganizationalUnitsTable holds the schema information for the "organizational_units" table.
	OrganizationalUnitsTable = &schema.Table{
		Name:       "organizational_units",
		Columns:    OrganizationalUnitsColumns,
		PrimaryKey: []*schema.Column{OrganizationalUnitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organizational_units_organizational_units_children",
				Columns:    []*schema.Column{OrganizationalUnitsColumns[5]},
				RefColumns: []*schema.Column{OrganizationalUnitsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "organizational_units_organizational_unit_types_organizational_units",
				Columns:    []*schema.Column{OrganizationalUnitsColumns[6]},
				RefColumns: []*schema.Column{OrganizationalUnitTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrganizationalUnitTypesColumns holds the columns for the "organizational_unit_types" table.
	OrganizationalUnitTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 16, SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "name", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "VARCHAR(256)"}},
	}
	// OrganizationalUnitTypesTable holds the schema information for the "organizational_unit_types" table.
	OrganizationalUnitTypesTable = &schema.Table{
		Name:       "organizational_unit_types",
		Columns:    OrganizationalUnitTypesColumns,
		PrimaryKey: []*schema.Column{OrganizationalUnitTypesColumns[0]},
	}
	// SpecialtiesColumns holds the columns for the "specialties" table.
	SpecialtiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 16, SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "name", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "VARCHAR(256)"}},
	}
	// SpecialtiesTable holds the schema information for the "specialties" table.
	SpecialtiesTable = &schema.Table{
		Name:       "specialties",
		Columns:    SpecialtiesColumns,
		PrimaryKey: []*schema.Column{SpecialtiesColumns[0]},
	}
	// OrganizationalUnitSpecialtiesColumns holds the columns for the "organizational_unit_specialties" table.
	OrganizationalUnitSpecialtiesColumns = []*schema.Column{
		{Name: "organizational_unit_id", Type: field.TypeString, Size: 16},
		{Name: "specialty_id", Type: field.TypeString, Size: 16},
	}
	// OrganizationalUnitSpecialtiesTable holds the schema information for the "organizational_unit_specialties" table.
	OrganizationalUnitSpecialtiesTable = &schema.Table{
		Name:       "organizational_unit_specialties",
		Columns:    OrganizationalUnitSpecialtiesColumns,
		PrimaryKey: []*schema.Column{OrganizationalUnitSpecialtiesColumns[0], OrganizationalUnitSpecialtiesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organizational_unit_specialties_organizational_unit_id",
				Columns:    []*schema.Column{OrganizationalUnitSpecialtiesColumns[0]},
				RefColumns: []*schema.Column{OrganizationalUnitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "organizational_unit_specialties_specialty_id",
				Columns:    []*schema.Column{OrganizationalUnitSpecialtiesColumns[1]},
				RefColumns: []*schema.Column{SpecialtiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrganizationalUnitsTable,
		OrganizationalUnitTypesTable,
		SpecialtiesTable,
		OrganizationalUnitSpecialtiesTable,
	}
)

func init() {
	OrganizationalUnitsTable.ForeignKeys[0].RefTable = OrganizationalUnitsTable
	OrganizationalUnitsTable.ForeignKeys[1].RefTable = OrganizationalUnitTypesTable
	OrganizationalUnitSpecialtiesTable.ForeignKeys[0].RefTable = OrganizationalUnitsTable
	OrganizationalUnitSpecialtiesTable.ForeignKeys[1].RefTable = SpecialtiesTable
}
