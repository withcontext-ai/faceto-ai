// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

const (
	RoomVodStatusReady = iota
	RoomVodStatusStarting
	RoomVodStatusComplete
	RoomVodStatusFail
)

const (
	VodPlatFormAzure = iota + 1
	VodPlatFormGCP
	VodPlatFormS3
)

const (
	VodTypeFile = iota + 1
	VodTypeStreamData
)

type RoomVod struct {
	ent.Schema
}

func (RoomVod) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Default(""),
		field.String("sid").Default(""),
		field.String("egress_id").Default(""),
		field.Uint8("status").Default(0).Optional(),     // vod status 0.init 1.start 2.complete 3.fail
		field.Uint8("platform").Default(1).Optional(),   // vod storage platform 1.azure 2.gcp 3.s3
		field.Uint8("vod_type").Default(1).Optional(),   // vod type 1.file 2.steam data
		field.String("vod_path").Default("").Optional(), // room vod path
		field.String("vod_url").Default("").Optional(),  // room vod url
		field.Time("start_time").Optional(),
		field.Time("complete_time").Optional(),
		field.Uint64("duration").Default(0).Optional(), // vod duration. unit:s
	}
}

func (RoomVod) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (RoomVod) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "room_vod"}}
}

func (RoomVod) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UUIDMixin{},
		TimeMixin{},
	}
}
