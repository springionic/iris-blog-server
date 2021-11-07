package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MinLen(1).MaxLen(255).Default("").Comment("标签名称"),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("articles", Article.Type).Ref("tags").Comment("当前标签下的所有文章"),
	}
}

func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{mixin.Time{}}
}
