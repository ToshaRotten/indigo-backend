package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Time("order_date").Optional(),
		field.Time("desired_delivery_date").Optional(),
		field.Float("total_amount").Optional(),
		field.Int("order_status_id").Optional(),
		field.Int("user_id").Optional(),
		field.Int("address_id").Optional(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order_status", OrderStatus.Type).Ref("status").Field("order_status_id").Unique(),
		edge.From("user", User.Type).Ref("order_user").Field("user_id").Unique(),
		edge.From("addres_order", Address.Type).Ref("address").Field("address_id").Unique(),
		edge.To("order", OrderDetail.Type),
	}
}
