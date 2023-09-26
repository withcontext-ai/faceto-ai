package biz

import (
	"errors"
)

var (
	ErrInterviewNotFound = errors.New("the interview not found")
	ErrRoomNotFound      = errors.New("the room not found")
	ErrRoomHistoryFound  = errors.New("the room history msg not found")
	ErrAuthFound         = errors.New("the client id of auth not found")
	ErrRoomVodNotFound   = errors.New("the room vod not found")
)
