package store

import (
	"github.com/Brightscout/mattermost-plugin-survey/server/model"
)

// Store allows the interaction with some kind of store.
type Store interface {
	Survey() SurveyStore
}

// SurveyStore allows the access to surveys in the store.
type SurveyStore interface {
	GetLatestSurveyVersion(id string) (*model.LatestSurveyVersion, error)
	SaveLatestSurveyVersion(id, version string) error
	GetSurvey(id, version string) (*model.Survey, error)
	SaveSurvey(survey *model.Survey) error
	GetMeetingMetadata(meetingID string) (*model.MeetingMetadata, error)
	SaveMeetingMetadata(data *model.MeetingMetadata) error
	SaveSurveyResponse(response *model.SurveyResponse) error
}
