package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("content").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("deleted_at").Optional(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("topic_id", uuid.UUID{}),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("questions").Unique().Required().Field("user_id"),
		edge.From("topic", Topic.Type).Ref("questions").Unique().Required().Field("topic_id"),
		edge.To("answers", Answer.Type),
		edge.To("tags", Tag.Type).Through("question_tags", QuestionTag.Type),

		edge.To("votes", QuestionVote.Type),
	}
}
