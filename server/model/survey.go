package model

import (
	"encoding/json"

	serverModel "github.com/mattermost/mattermost-server/model"
)

type ValueType string

const (
	TypeSurvey           ValueType = "survey"
	TypeSurveyResponse   ValueType = "survey_response"
	TypeMeetingMetadata  ValueType = "meeting_metadata"
	TypeLatestSurveyInfo ValueType = "latest_survey_info"
)

type QuestionType string

const (
	TypeOpen                 QuestionType = "open"
	TypeFivePointLikertScale QuestionType = "five-point-likert-scale"
)

// Survey stores a survey question
type SurveyQuestion struct {
	ID   string       `json:"id"`
	Type QuestionType `json:"type"`
	Text string       `json:"text"`
}

// Survey stores all needed information for a survey
type Survey struct {
	Type        ValueType         `json:"type"`
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
	Type          ValueType
	UserID        string
	MeetingID     string
	SurveyID      string
	SurveyVersion string
	CreatedAt     int64
	Responses     map[string]string
}

// MeetingMetadata stores the survey metadata for a  meeting
type MeetingMetadata struct {
	Type          ValueType
	MeetingID     string
	SurveyID      string
	SurveyVersion int
	UserResponded map[string]bool
}

// LatestSurveyInfo stores the latest version information for a survey
type LatestSurveyInfo struct {
	Type    ValueType `json:"type"`
	ID      string    `json:"id"`
	Version int       `json:"version"`
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
