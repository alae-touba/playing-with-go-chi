package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").Sensitive().NotEmpty(), // Mark password as sensitive
		field.String("image_name").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type),
		edge.To("answers", Answer.Type),
		edge.To("topics", Topic.Type),
		edge.To("roles", Role.Type).Through("user_roles", UserRole.Type),
		edge.To("question_votes", QuestionVote.Type),
		edge.To("answer_votes", AnswerVote.Type),
	}
}
