package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderStatus holds the schema definition for the OrderStatus entity.
type OrderStatus struct {
	ent.Schema
}

// Fields of the OrderStatus.
func (OrderStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().MaxLen(255),
	}
}

// Edges of the OrderStatus.
func (OrderStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("status", Order.Type),
	}
}
