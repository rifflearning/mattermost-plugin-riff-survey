package model

import (
	"encoding/json"
)

// ReminderMetadata stores metadata related to the reminder settings for a survey
type ReminderMetadata struct {
	Type                   string `json:"type"`
	MeetingID              string `json:"meeting_id"`
	UserID                 string `json:"user_id"`
	PostID                 string `json:"post_id"`
	ChannelID              string `json:"channel_id"`
	SurveySentAt           int64  `json:"survey_sent_at"`
	PreviousReminderSentAt int64  `json:"previous_reminder_sent_at"`
	TotalRemindersSent     int    `json:"total_reminders_sent"`
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
