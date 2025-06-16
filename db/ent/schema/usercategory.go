package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserCategory holds the schema definition for the UserCategory entity.
type UserCategory struct {
	ent.Schema
}

// Fields of the UserCategory.
func (UserCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().MaxLen(255),
	}
}

// Edges of the UserCategory.
func (UserCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_category_user", User.Type),
	}
}
