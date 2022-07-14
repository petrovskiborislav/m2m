// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/bug/ent/organizationalunit"
	"entgo.io/bug/ent/organizationalunittype"
	"entgo.io/ent/dialect/sql"
)

// OrganizationalUnit is the model entity for the OrganizationalUnit schema.
type OrganizationalUnit struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID string `json:"parent_id,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata []string `json:"metadata,omitempty"`
	// TypeID holds the value of the "type_id" field.
	TypeID string `json:"type_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationalUnitQuery when eager-loading is set.
	Edges OrganizationalUnitEdges `json:"edges"`
}

// OrganizationalUnitEdges holds the relations/edges for other nodes in the graph.
type OrganizationalUnitEdges struct {
	// Parent holds the value of the parent edge.
	Parent *OrganizationalUnit `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*OrganizationalUnit `json:"children,omitempty"`
	// OrganizationalUnitTypeID holds the value of the organizational_unit_type_id edge.
	OrganizationalUnitTypeID *OrganizationalUnitType `json:"organizational_unit_type_id,omitempty"`
	// Specialties holds the value of the specialties edge.
	Specialties []*Specialty `json:"specialties,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]*int
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationalUnitEdges) ParentOrErr() (*OrganizationalUnit, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: organizationalunit.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationalUnitEdges) ChildrenOrErr() ([]*OrganizationalUnit, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// OrganizationalUnitTypeIDOrErr returns the OrganizationalUnitTypeID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationalUnitEdges) OrganizationalUnitTypeIDOrErr() (*OrganizationalUnitType, error) {
	if e.loadedTypes[2] {
		if e.OrganizationalUnitTypeID == nil {
			// The edge organizational_unit_type_id was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: organizationalunittype.Label}
		}
		return e.OrganizationalUnitTypeID, nil
	}
	return nil, &NotLoadedError{edge: "organizational_unit_type_id"}
}

// SpecialtiesOrErr returns the Specialties value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationalUnitEdges) SpecialtiesOrErr() ([]*Specialty, error) {
	if e.loadedTypes[3] {
		return e.Specialties, nil
	}
	return nil, &NotLoadedError{edge: "specialties"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrganizationalUnit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case organizationalunit.FieldMetadata:
			values[i] = new([]byte)
		case organizationalunit.FieldID, organizationalunit.FieldDisplayName, organizationalunit.FieldParentID, organizationalunit.FieldTypeID:
			values[i] = new(sql.NullString)
		case organizationalunit.FieldCreatedAt, organizationalunit.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrganizationalUnit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrganizationalUnit fields.
func (ou *OrganizationalUnit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organizationalunit.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ou.ID = value.String
			}
		case organizationalunit.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				ou.DisplayName = value.String
			}
		case organizationalunit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ou.CreatedAt = value.Time
			}
		case organizationalunit.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ou.DeletedAt = value.Time
			}
		case organizationalunit.FieldParentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				ou.ParentID = value.String
			}
		case organizationalunit.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ou.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		case organizationalunit.FieldTypeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type_id", values[i])
			} else if value.Valid {
				ou.TypeID = value.String
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the OrganizationalUnit entity.
func (ou *OrganizationalUnit) QueryParent() *OrganizationalUnitQuery {
	return (&OrganizationalUnitClient{config: ou.config}).QueryParent(ou)
}

// QueryChildren queries the "children" edge of the OrganizationalUnit entity.
func (ou *OrganizationalUnit) QueryChildren() *OrganizationalUnitQuery {
	return (&OrganizationalUnitClient{config: ou.config}).QueryChildren(ou)
}

// QueryOrganizationalUnitTypeID queries the "organizational_unit_type_id" edge of the OrganizationalUnit entity.
func (ou *OrganizationalUnit) QueryOrganizationalUnitTypeID() *OrganizationalUnitTypeQuery {
	return (&OrganizationalUnitClient{config: ou.config}).QueryOrganizationalUnitTypeID(ou)
}

// QuerySpecialties queries the "specialties" edge of the OrganizationalUnit entity.
func (ou *OrganizationalUnit) QuerySpecialties() *SpecialtyQuery {
	return (&OrganizationalUnitClient{config: ou.config}).QuerySpecialties(ou)
}

// Update returns a builder for updating this OrganizationalUnit.
// Note that you need to call OrganizationalUnit.Unwrap() before calling this method if this OrganizationalUnit
// was returned from a transaction, and the transaction was committed or rolled back.
func (ou *OrganizationalUnit) Update() *OrganizationalUnitUpdateOne {
	return (&OrganizationalUnitClient{config: ou.config}).UpdateOne(ou)
}

// Unwrap unwraps the OrganizationalUnit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ou *OrganizationalUnit) Unwrap() *OrganizationalUnit {
	_tx, ok := ou.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrganizationalUnit is not a transactional entity")
	}
	ou.config.driver = _tx.drv
	return ou
}

// String implements the fmt.Stringer.
func (ou *OrganizationalUnit) String() string {
	var builder strings.Builder
	builder.WriteString("OrganizationalUnit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ou.ID))
	builder.WriteString("display_name=")
	builder.WriteString(ou.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ou.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ou.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(ou.ParentID)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", ou.Metadata))
	builder.WriteString(", ")
	builder.WriteString("type_id=")
	builder.WriteString(ou.TypeID)
	builder.WriteByte(')')
	return builder.String()
}

// OrganizationalUnits is a parsable slice of OrganizationalUnit.
type OrganizationalUnits []*OrganizationalUnit

func (ou OrganizationalUnits) config(cfg config) {
	for _i := range ou {
		ou[_i].config = cfg
	}
}
