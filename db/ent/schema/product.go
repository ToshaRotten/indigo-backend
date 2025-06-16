package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().MaxLen(255),
		field.Float("weight").Optional(),
		field.String("product_composition").Optional().MaxLen(255),
		field.Int("min_storage_temp").Optional(),
		field.Int("max_storage_temp").Optional(),
		field.String("shelf_life").Optional().MaxLen(255),
		field.String("picture").Optional().MaxLen(255),
		field.Int("product_category_id").Optional(),
		field.Int("type_of_packaging_id").Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order_detail_product", OrderDetail.Type),
		edge.To("shopping_cart_product", ShoppingCart.Type),
		edge.To("product_price_product", ProductPrice.Type),
		edge.From("product_category", ProductCategory.Type).
			Ref("category").
			Field("product_category_id").
			Unique(),
		edge.From("type_of_packaging_product", TypeOfPackaging.Type).
			Ref("type_of_packaging").
			Field("type_of_packaging_id").
			Unique(),
	}
}
