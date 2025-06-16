package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").Optional().MaxLen(255),
		field.Time("expiration_date").Optional(),
		field.String("name").Optional().MaxLen(255),
		field.String("token").Optional().MaxLen(255),
		field.Int("user_id").Optional(),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_card").Field("user_id").Unique(),
	}
}
