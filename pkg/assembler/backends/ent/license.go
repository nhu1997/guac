// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/license"
)

// License is the model entity for the License schema.
type License struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Inline holds the value of the "inline" field.
	Inline *string `json:"inline,omitempty"`
	// ListVersion holds the value of the "list_version" field.
	ListVersion  *string `json:"list_version,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*License) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case license.FieldID:
			values[i] = new(sql.NullInt64)
		case license.FieldName, license.FieldInline, license.FieldListVersion:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the License fields.
func (l *License) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case license.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case license.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case license.FieldInline:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inline", values[i])
			} else if value.Valid {
				l.Inline = new(string)
				*l.Inline = value.String
			}
		case license.FieldListVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field list_version", values[i])
			} else if value.Valid {
				l.ListVersion = new(string)
				*l.ListVersion = value.String
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the License.
// This includes values selected through modifiers, order, etc.
func (l *License) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// Update returns a builder for updating this License.
// Note that you need to call License.Unwrap() before calling this method if this License
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *License) Update() *LicenseUpdateOne {
	return NewLicenseClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the License entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *License) Unwrap() *License {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: License is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *License) String() string {
	var builder strings.Builder
	builder.WriteString("License(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	if v := l.Inline; v != nil {
		builder.WriteString("inline=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := l.ListVersion; v != nil {
		builder.WriteString("list_version=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Licenses is a parsable slice of License.
type Licenses []*License
