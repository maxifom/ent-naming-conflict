package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"ent-naming-conflict/internal/message"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(message.MessageType("")),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
