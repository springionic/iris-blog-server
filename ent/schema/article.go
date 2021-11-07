package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default("").MaxLen(255).MinLen(1).Comment("文章标题"),
		field.Text("content").Default("").Comment("文章内容"),
		field.Int("user_id").Comment("文章的用户 id"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Field("user_id").Ref("articles").Unique().Required().Comment("文章的用户"),
		edge.To("tags", Tag.Type).Comment("文章的所有标签"),
	}
}

func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{mixin.Time{}}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
