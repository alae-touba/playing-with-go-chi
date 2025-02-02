package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AnswerVote holds the schema definition for the AnswerVote entity.
type AnswerVote struct {
	ent.Schema
}

// Fields of the AnswerVote.
func (AnswerVote) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).
			StorageKey("user_id"),
		field.UUID("answer_id", uuid.UUID{}).
			StorageKey("answer_id"),
		field.Enum("vote_type").Values("upvote", "downvote"),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the AnswerVote.
func (AnswerVote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("answer_votes").
			Required().
			Unique().
			Field("user_id"),
		edge.From("answer", Answer.Type).
			Ref("votes").
			Required().
			Unique().
			Field("answer_id"),
	}
}
