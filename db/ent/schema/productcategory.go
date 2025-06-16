package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductCategory holds the schema definition for the ProductCategory entity.
type ProductCategory struct {
	ent.Schema
}

// Fields of the ProductCategory.
func (ProductCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().MaxLen(255),
	}
}

// Edges of the ProductCategory.
func (ProductCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("category", Product.Type),
	}
}
