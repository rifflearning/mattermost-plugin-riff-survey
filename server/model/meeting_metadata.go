package model

import (
	"encoding/json"
)

// MeetingMetadata stores the survey metadata for a  meeting
type MeetingMetadata struct {
	Type          string `json:"type"`
	MeetingID     string `json:"meeting_id"`
	SurveyID      string `json:"survey_id"`
	SurveyVersion int    `json:"survey_version"`
}

func (m *MeetingMetadata) PreSave() *MeetingMetadata {
	m.Type = TypeMeetingMetadata
	return m
}

// EncodeToByte returns a meeting metadata as a byte array
func (m *MeetingMetadata) EncodeToByte() []byte {
	b, _ := json.Marshal(m)
	return b
}

// DecodeMeetingMetadataFromByte tries to create a meeting metadata from a byte array
func DecodeMeetingMetadataFromByte(b []byte) *MeetingMetadata {
	m := MeetingMetadata{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil
	}
	return &m
}
