package model

import (
	"encoding/json"
)

// UserMeetingMetadata stores the user metadata for a  meeting
type UserMeetingMetadata struct {
	Type         string `json:"type"`
	UserID       string `json:"user_id"`
	MeetingID    string `json:"meeting_id"`
	SurveySentAt int64  `json:"survey_sent_at"`
	RespondedAt  int64  `json:"responded_at"`
}

func (u *UserMeetingMetadata) PreSave() *UserMeetingMetadata {
	u.Type = TypeUserMeetingMetadata
	return u
}

// EncodeToByte returns a user meeting metadata as a byte array
func (u *UserMeetingMetadata) EncodeToByte() []byte {
	b, _ := json.Marshal(u)
	return b
}

// DecodeUserMeetingMetadataFromByte tries to create a user meeting metadata from a byte array
func DecodeUserMeetingMetadataFromByte(b []byte) *UserMeetingMetadata {
	u := UserMeetingMetadata{}
	err := json.Unmarshal(b, &u)
	if err != nil {
		return nil
	}
	return &u
}
