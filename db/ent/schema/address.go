package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("city").Optional().MaxLen(255),
		field.String("street_name").Optional().MaxLen(255),
		field.String("house_number").Optional().MaxLen(255),
		field.String("store_name").Optional().MaxLen(255),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("addresses"),
		edge.To("address", Order.Type),
	}
}
