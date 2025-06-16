package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").Optional().MaxLen(255),
		field.String("sended_at").Optional().MaxLen(255),
		field.String("file_path").Optional().MaxLen(255),
		field.String("file_type").Optional().MaxLen(255),
		field.Int("chat_id").Optional(),
		field.Int("user_id").Optional(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chat", Chat.Type).Ref("chat_message").Field("chat_id").Unique(),
		edge.From("user", User.Type).Ref("user_message").Field("user_id").Unique(),
	}
}
