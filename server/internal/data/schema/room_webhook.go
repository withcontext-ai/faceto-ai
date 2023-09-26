// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/livekit/protocol/livekit"
)

type RoomWebhook struct {
	ent.Schema
}

func (RoomWebhook) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Default(""),
		field.String("sid").Default(""),
		field.String("event").Default(""),
		field.Time("event_time").Optional(),
		field.JSON("room", &livekit.Room{}).Optional(),
		field.JSON("participant", &livekit.ParticipantInfo{}).Optional(),
		field.JSON("track", &livekit.TrackInfo{}).Optional(),
		field.JSON("egressInfo", &livekit.EgressInfo{}).Optional(),
		field.JSON("ingressInfo", &livekit.IngressInfo{}).Optional(),
	}
}

func (RoomWebhook) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (RoomWebhook) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "room_webhook"}}
}

func (RoomWebhook) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UUIDMixin{},
		TimeMixin{},
	}
}
