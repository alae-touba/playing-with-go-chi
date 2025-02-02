package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// QuestionVote holds the schema definition for the QuestionVote entity.
type QuestionVote struct {
	ent.Schema
}

// Fields of the QuestionVote.
func (QuestionVote) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).StorageKey("user_id"),
		field.UUID("question_id", uuid.UUID{}).StorageKey("question_id"),

		field.Enum("vote_type").Values("upvote", "downvote"),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the QuestionVote.
func (QuestionVote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("question_votes").
			Required().
			Unique().
			Field("user_id"),
		edge.From("question", Question.Type).
			Ref("votes").
			Required().
			Unique().
			Field("question_id"),
	}
}
