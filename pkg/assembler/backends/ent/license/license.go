// Code generated by ent, DO NOT EDIT.

package license

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the license type in the database.
	Label = "license"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldInline holds the string denoting the inline field in the database.
	FieldInline = "inline"
	// FieldListVersion holds the string denoting the list_version field in the database.
	FieldListVersion = "list_version"
	// Table holds the table name of the license in the database.
	Table = "licenses"
)

// Columns holds all SQL columns for license fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldInline,
	FieldListVersion,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the License queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByInline orders the results by the inline field.
func ByInline(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInline, opts...).ToFunc()
}

// ByListVersion orders the results by the list_version field.
func ByListVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldListVersion, opts...).ToFunc()
}
