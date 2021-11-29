package schema

import (
	"entgo.io/ent"
	"iris-blog-server/util"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return nil
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{util.Time{}}
}
