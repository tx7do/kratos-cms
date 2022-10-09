package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*CreateTime)(nil)

type CreateTime struct{ mixin.Schema }

func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		// 创建时间，毫秒
		field.Int64("create_time").
			Comment("创建时间").
			Immutable().
			Optional().
			Nillable().
			DefaultFunc(time.Now().UnixMilli),
	}
}

var _ ent.Mixin = (*UpdateTime)(nil)

type UpdateTime struct{ mixin.Schema }

func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		// 更新时间，毫秒
		// 需要注意的是，如果不是程序自动更新，那么这个字段不会被更新，除非在数据库里面下触发器更新。
		field.Int64("update_time").
			Comment("更新时间").
			Optional().
			Nillable().
			UpdateDefault(time.Now().UnixMilli),
	}
}

var _ ent.Mixin = (*DeleteTime)(nil)

type DeleteTime struct{ mixin.Schema }

func (DeleteTime) Fields() []ent.Field {
	return []ent.Field{
		// 删除时间，毫秒
		field.Int64("delete_time").
			Comment("删除时间").
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*Time)(nil)

type Time struct{ mixin.Schema }

func (Time) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, CreateTime{}.Fields()...)
	fields = append(fields, UpdateTime{}.Fields()...)
	fields = append(fields, DeleteTime{}.Fields()...)
	return fields
}

var _ ent.Mixin = (*CreatedAt)(nil)

type CreatedAt struct{ mixin.Schema }

func (CreatedAt) Fields() []ent.Field {
	return []ent.Field{
		// 创建时间
		field.Time("created_at").
			Comment("创建时间").
			Immutable().
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*UpdatedAt)(nil)

type UpdatedAt struct{ mixin.Schema }

func (UpdatedAt) Fields() []ent.Field {
	return []ent.Field{
		// 更新时间
		field.Time("updated_at").
			Comment("更新时间").
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*DeletedAt)(nil)

type DeletedAt struct{ mixin.Schema }

func (DeletedAt) Fields() []ent.Field {
	return []ent.Field{
		// 删除时间
		field.Time("deleted_at").
			Comment("删除时间").
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*TimeAt)(nil)

type TimeAt struct{ mixin.Schema }

func (TimeAt) Fields() []ent.Field {
	var fields []ent.Field
	fields = append(fields, CreatedAt{}.Fields()...)
	fields = append(fields, UpdatedAt{}.Fields()...)
	fields = append(fields, DeletedAt{}.Fields()...)
	return fields
}
