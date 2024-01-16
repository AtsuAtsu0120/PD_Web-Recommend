package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Website holds the schema definition for the Website entity.
type Website struct {
	ent.Schema
}

// Fields of the Website.
func (Website) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("filePath").Unique(),
		field.String("url").Unique(),
		field.Float("bright").Nillable(),
		field.Float("flashy").Nillable(),
		field.Float("adult").Nillable(),
		field.Float("smart").Nillable(),
		field.Float("beautiful").Nillable(),
		field.Float("like").Nillable(),
	}
}

// Edges of the Website.
func (Website) Edges() []ent.Edge {
	return nil
}
