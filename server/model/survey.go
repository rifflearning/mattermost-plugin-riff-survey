package model

type ValueType string

const (
	TypeSurvey              ValueType = "survey"
	TypeSurveyResponse      ValueType = "survey_response"
	TypeMeetingMetadata     ValueType = "meeting_metadata"
	TypeLatestSurveyVersion ValueType = "latest_survey_version"
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
	CreatedAt   int               `json:"created_at"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Questions   []*SurveyQuestion `json:"questions"`
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

// LatestSurveyVersion stores the latest version information for a survey
type LatestSurveyVersion struct {
	Type    ValueType `json:"type"`
	ID      string    `json:"id"`
	Version int       `json:"version"`
}
