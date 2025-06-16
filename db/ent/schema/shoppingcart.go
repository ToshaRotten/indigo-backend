package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ShoppingCart holds the schema definition for the ShoppingCart entity.
type ShoppingCart struct {
	ent.Schema
}

// Fields of the ShoppingCart.
func (ShoppingCart) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity").Optional(),
		field.Int("product_id").Optional(),
		field.Int("user_id").Optional(),
	}
}

// Edges of the ShoppingCart.
func (ShoppingCart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).
			Ref("shopping_cart_product").
			Field("product_id").
			Unique(),
		edge.From("users", User.Type).Ref("shopping_cart_user").Field("user_id").Unique(),
	}
}
