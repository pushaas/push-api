package models

type (
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

type ErrorCode int

const (
	/*
		channels
	*/
	ErrorChannelCreateInvalidBody     = 10
	ErrorChannelCreateIdAlreadyExists = 11
	ErrorChannelCreateFailed          = 12

	ErrorChannelGetNotFound = 20
	ErrorChannelGetFailed   = 21

	ErrorChannelDeleteNotFound = 30
	ErrorChannelDeleteFailed   = 31

	ErrorChannelGetAllFailed = 40

	/*
		messages
	*/
	ErrorMessageCreateInvalidBody          = 100
	ErrorMessageCreateFailed               = 101
	ErrorMessageCreateInvalidMessageFormat = 102

	/*
		stats
	*/
	ErrorStatsGetNotFound = 200
	ErrorStatsGetFailed   = 201

	/*
		config
	*/
	ErrorConfigGetParsePushStreamUrl = 300
)
