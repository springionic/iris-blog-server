package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("nick_name").MaxLen(255).MinLen(2).Default("").Comment("用户自定义昵称"),
		field.Enum("sex").Values("UNKNOWN", "MALE", "FEMALE").Default("UNKNOWN").Comment("用户性别"),
		field.Time("birthday").Nillable().Optional().Comment("出生日期"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("articles", Article.Type).Comment("用户的文章"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{mixin.Time{}}
}
