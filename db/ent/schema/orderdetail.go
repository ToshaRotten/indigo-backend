package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderDetail holds the schema definition for the OrderDetail entity.
type OrderDetail struct {
	ent.Schema
}

// Fields of the OrderDetail.
func (OrderDetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity").Optional(),
		field.Int("order_id").Optional(),
		field.Int("product_id").Optional(),
	}
}

// Edges of the OrderDetail.
func (OrderDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("order").Field("order_id").Unique(),
		edge.From("products", Product.Type).Ref("order_detail_product").Field("order_id").Unique(),
	}
}
