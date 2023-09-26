package schema

import (
	"crypto/rand"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid"
)

type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

type UUIDMixin struct {
	mixin.Schema
}

func (UUIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").DefaultFunc(MustUID).Unique(),
		// field.UUID("uuid", uuid.UUID{}).Default(uuid.New).SchemaType(map[string]string{
		// 	dialect.MySQL:    "binary",
		// }),
	}
}

func MustUID() string {
	ms := ulid.Timestamp(time.Now())
	uid, err := ulid.New(ms, rand.Reader)
	if err != nil {
		panic(err)
	}
	return uid.String()
}
