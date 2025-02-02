package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Answer holds the schema definition for the Answer entity.
type Answer struct {
	ent.Schema
}

// Fields of the Answer.
func (Answer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("content").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		// field.Time("deleted_at").Optional(),
		field.UUID("user_id", uuid.UUID{}),
        field.UUID("question_id", uuid.UUID{}),
	}
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("answers").Unique().Required().Field("user_id"),
		edge.From("question", Question.Type).Ref("answers").Unique().Required().Field("question_id"),
		
		edge.To("votes", AnswerVote.Type),
	}
}
