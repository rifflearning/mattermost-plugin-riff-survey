package model

import (
	"encoding/json"

	serverModel "github.com/mattermost/mattermost-server/model"
)

const (
	TypeSurvey              = "survey"
	TypeSurveyResponse      = "survey_response"
	TypeMeetingMetadata     = "meeting_metadata"
	TypeLatestSurveyInfo    = "latest_survey_info"
	TypeReminderMetadata    = "reminder_metadata"
	TypeUserMeetingMetadata = "user_meeting_metadata"

	QuestionTypeOpen                 = "open"
	QuestionTypeFivePointLikertScale = "five-point-likert-scale"
)

// Survey stores a survey question
type SurveyQuestion struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

// Survey stores all needed information for a survey
type Survey struct {
	Type        string            `json:"type"`
	ID          string            `json:"id"`
	Version     int               `json:"version"`
	CreatedAt   int64             `json:"created_at"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Questions   []*SurveyQuestion `json:"questions"`
}

func (s *Survey) PreSave(version int) *Survey {
	s.Type = TypeSurvey
	s.Version = version
	s.CreatedAt = serverModel.GetMillis()
	for _, question := range s.Questions {
		question.ID = serverModel.NewId()
	}
	return s
}

// Equals checks the equality of deterministic fields of two surveys.
// Checks for the equality of Title, Description, Number of Questions and
// Type and Text for the Questions
func (s *Survey) Equals(survey *Survey) bool {
	if s.Title != survey.Title {
		return false
	}

	if s.Description != survey.Description {
		return false
	}

	if len(s.Questions) != len(survey.Questions) {
		return false
	}

	for i := range s.Questions {
		if s.Questions[i].Text != survey.Questions[i].Text {
			return false
		}
		if s.Questions[i].Type != survey.Questions[i].Type {
			return false
		}
	}

	return true
}

// EncodeToByte returns a survey as a byte array
func (s *Survey) EncodeToByte() []byte {
	b, _ := json.Marshal(s)
	return b
}

// DecodeSurveyFromByte tries to create a survey from a byte array
func DecodeSurveyFromByte(b []byte) *Survey {
	s := Survey{}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return nil
	}
	return &s
}

// SurveyResponse records the users responses to the survey questions
type SurveyResponse struct {
	Type          string            `json:"type"`
	UserID        string            `json:"user_id"`
	MeetingID     string            `json:"meeting_id"`
	SurveyID      string            `json:"survey_id"`
	SurveyVersion int               `json:"survey_version"`
	CreatedAt     int64             `json:"created_at"`
	Responses     map[string]string `json:"responses"`
}

func (r *SurveyResponse) PreSave() *SurveyResponse {
	r.Type = TypeSurveyResponse
	r.CreatedAt = serverModel.GetMillis()
	return r
}

// EncodeToByte returns a survey response as a byte array
func (r *SurveyResponse) EncodeToByte() []byte {
	b, _ := json.Marshal(r)
	return b
}

// DecodeSurveyResponseFromByte tries to create a survey response from a byte array
func DecodeSurveyResponseFromByte(b []byte) *SurveyResponse {
	r := SurveyResponse{}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return nil
	}
	return &r
}

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

// LatestSurveyInfo stores the latest version information for a survey
type LatestSurveyInfo struct {
	Type          string `json:"type"`
	SurveyID      string `json:"survey_id"`
	SurveyVersion int    `json:"survey_version"`
}

func (info *LatestSurveyInfo) PreSave() *LatestSurveyInfo {
	info.Type = TypeLatestSurveyInfo
	return info
}

// EncodeToByte returns a LatestSurveyInfo as a byte array
func (info *LatestSurveyInfo) EncodeToByte() []byte {
	b, _ := json.Marshal(info)
	return b
}

// DecodeLatestSurveyInfoFromByte tries to create a LatestSurveyInfo from a byte array
func DecodeLatestSurveyInfoFromByte(b []byte) *LatestSurveyInfo {
	info := LatestSurveyInfo{}
	err := json.Unmarshal(b, &info)
	if err != nil {
		return nil
	}
	return &info
}

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
