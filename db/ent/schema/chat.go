package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Optional().MaxLen(255),
		field.String("created_at").Optional().MaxLen(255),
		field.Int("user_id").Optional(),
		field.Int("manager_id").Optional(),
	}
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_chat").Field("user_id").Unique(),
		edge.To("chat_message", Message.Type),
	}
}
