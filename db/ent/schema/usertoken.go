package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserToken holds the schema definition for the UserToken entity.
type UserToken struct {
	ent.Schema
}

// Fields of the UserToken.
func (UserToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Optional().MaxLen(255),
		field.Time("expiration_date").Optional(),
		field.String("type").Optional().MaxLen(255),
		field.String("status").Optional(),
		field.String("last_used").Optional().MaxLen(255),
		field.Int("user_id").Optional(),
	}
}

// Edges of the UserToken.
func (UserToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_token_user").Field("user_id").Unique(),
	}
}
