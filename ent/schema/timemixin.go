package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct {
	mixin.Schema
}

// Fields of the TimeMixin.
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now().UTC).
			Comment("timestamp of creation"),
		field.Time("updated_at").
			Default(time.Now().UTC).
			UpdateDefault(time.Now().UTC).
			Comment("timestamp of last entity update"),
	}
}
