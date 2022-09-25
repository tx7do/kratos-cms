package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type Time struct {
	mixin.Schema
}

func (Time) Fields() []ent.Field {
	return []ent.Field{
		// 创建时间,使用毫秒单位
		field.Int64("create_time").
			Comment("创建时间").
			Immutable().
			Optional().
			Nillable().
			DefaultFunc(time.Now().UnixMilli),

		// 更新时间,使用毫秒单位
		// 需要注意的是,如果不是程序自动更新,那么这个字段不会被更新,除非在数据库里面下触发器更新
		field.Int64("update_time").
			Comment("更新时间").
			Optional().
			Nillable().
			UpdateDefault(time.Now().UnixMilli),

		// 删除时间,使用毫秒单位
		field.Int64("delete_time").
			Comment("删除时间").
			Optional().
			Nillable(),
	}
}
