package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductPrice holds the schema definition for the ProductPrice entity.
type ProductPrice struct {
	ent.Schema
}

// Fields of the ProductPrice.
func (ProductPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Time("modification_date").Optional(),
		field.Float("new_price").Optional(),
		field.Int("product_id").Optional(),
	}
}

// Edges of the ProductPrice.
func (ProductPrice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).
			Ref("product_price_product").
			Field("product_id").
			Unique(),
	}
}
