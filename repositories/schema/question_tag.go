package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// QuestionTag holds the schema definition for the QuestionTag entity.
type QuestionTag struct {
	ent.Schema
}

// Fields of the QuestionTag.
func (QuestionTag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("question_id", uuid.UUID{}).StorageKey("question_id"),
		field.UUID("tag_id", uuid.UUID{}).StorageKey("tag_id"),
	}
}

// Edges of the QuestionTag.
func (QuestionTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("question", Question.Type).
			Required().
			Unique().
			Field("question_id"),
		edge.To("tag", Tag.Type).
			Required().
			Unique().
			Field("tag_id"),
	}
}
