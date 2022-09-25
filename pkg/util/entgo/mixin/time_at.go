package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeAt struct {
	mixin.Schema
}

func (TimeAt) Fields() []ent.Field {
	return []ent.Field{
		// 创建时间
		field.Time("created_at").
			Comment("创建时间").
			Immutable().
			Optional().
			Nillable(),

		// 更新时间
		field.Time("updated_at").
			Comment("更新时间").
			Optional().
			Nillable(),

		// 删除时间
		field.Time("deleted_at").
			Comment("删除时间").
			Optional().
			Nillable(),
	}
}
