package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").Optional().MaxLen(255),
		field.String("username").Unique().MaxLen(255),
		field.String("email").Optional().MaxLen(255),
		field.String("password_hash").Optional().MaxLen(255),
		field.String("phone_number").Optional().MaxLen(255),
		field.Int("user_category_id").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_category", UserCategory.Type).
			Ref("user_category_user").
			Field("user_category_id").
			Unique(),
		edge.To("shopping_cart_user", ShoppingCart.Type),
		edge.To("user_token_user", UserToken.Type),
		edge.To("order_user", Order.Type),
		edge.To("addresses", Address.Type),
		edge.To("user_chat", Chat.Type),
		edge.To("user_card", Card.Type),
		edge.To("user_message", Message.Type),
	}
}
