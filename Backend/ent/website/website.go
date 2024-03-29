// Code generated by ent, DO NOT EDIT.

package website

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the website type in the database.
	Label = "website"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFilePath holds the string denoting the filepath field in the database.
	FieldFilePath = "file_path"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldBright holds the string denoting the bright field in the database.
	FieldBright = "bright"
	// FieldFlashy holds the string denoting the flashy field in the database.
	FieldFlashy = "flashy"
	// FieldAdult holds the string denoting the adult field in the database.
	FieldAdult = "adult"
	// FieldSmart holds the string denoting the smart field in the database.
	FieldSmart = "smart"
	// FieldBeautiful holds the string denoting the beautiful field in the database.
	FieldBeautiful = "beautiful"
	// FieldLike holds the string denoting the like field in the database.
	FieldLike = "like"
	// Table holds the table name of the website in the database.
	Table = "websites"
)

// Columns holds all SQL columns for website fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldFilePath,
	FieldURL,
	FieldBright,
	FieldFlashy,
	FieldAdult,
	FieldSmart,
	FieldBeautiful,
	FieldLike,
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

// OrderOption defines the ordering options for the Website queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFilePath orders the results by the filePath field.
func ByFilePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFilePath, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByBright orders the results by the bright field.
func ByBright(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBright, opts...).ToFunc()
}

// ByFlashy orders the results by the flashy field.
func ByFlashy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFlashy, opts...).ToFunc()
}

// ByAdult orders the results by the adult field.
func ByAdult(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdult, opts...).ToFunc()
}

// BySmart orders the results by the smart field.
func BySmart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSmart, opts...).ToFunc()
}

// ByBeautiful orders the results by the beautiful field.
func ByBeautiful(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBeautiful, opts...).ToFunc()
}

// ByLike orders the results by the like field.
func ByLike(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLike, opts...).ToFunc()
}
