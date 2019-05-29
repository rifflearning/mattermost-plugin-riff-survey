package model

import (
	"encoding/json"
)

// ReminderMetadata stores metadata related to the reminder settings for a survey
type ReminderMetadata struct {
	Type                   string
	MeetingID              string
	UserID                 string
	PostID                 string
	ChannelID              string
	SurveySentAt           int64
	PreviousReminderSentAt int64
	TotalRemindersSent     int
}

func (r *ReminderMetadata) PreSave() *ReminderMetadata {
	r.Type = TypeReminderMetadata
	return r
}

// EncodeToByte returns a LatestSurveyInfo as a byte array
func (r *ReminderMetadata) EncodeToByte() []byte {
	b, _ := json.Marshal(r)
	return b
}

// DecodeReminderMetadataFromByte tries to create a ReminderMetadata from a byte array
func DecodeReminderMetadataFromByte(b []byte) *ReminderMetadata {
	r := ReminderMetadata{}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return nil
	}
	return &r
}
