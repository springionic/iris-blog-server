// Package util
// created by lilei at 2021/11/29
package util

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Time struct{ mixin.Time }

func (Time) Fields() []ent.Field {
	return append(
		[]ent.Field{
			field.Time("delete_time").Nillable().Optional().Default(nil).StructTag(`json:"-"`).Comment("删除时间"),
		},
		mixin.Time{}.Fields()...,
	)
}
