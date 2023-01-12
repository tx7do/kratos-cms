package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var _ ent.Mixin = (*CreateBy)(nil)

type CreateBy struct{ mixin.Schema }

func (CreateBy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("create_by").
			Comment("创建者").
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*UpdateBy)(nil)

type UpdateBy struct{ mixin.Schema }

func (UpdateBy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("update_by").
			Comment("更新者").
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*DeleteBy)(nil)

type DeleteBy struct{ mixin.Schema }

func (DeleteBy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("delete_by").
			Comment("删除者").
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*CreatedBy)(nil)

type CreatedBy struct{ mixin.Schema }

func (CreatedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("created_by").
			Comment("创建者").
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*UpdatedBy)(nil)

type UpdatedBy struct{ mixin.Schema }

func (UpdatedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("updated_by").
			Comment("更新者").
			Optional().
			Nillable(),
	}
}

var _ ent.Mixin = (*DeletedBy)(nil)

type DeletedBy struct{ mixin.Schema }

func (DeletedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("deleted_by").
			Comment("删除者").
			Optional().
			Nillable(),
	}
}
