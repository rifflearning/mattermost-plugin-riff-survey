package store

import (
	"github.com/Brightscout/mattermost-plugin-survey/server/model"
)

// Store allows the interaction with some kind of store.
type Store interface {
	GetLatestSurveyInfo(id string) (*model.LatestSurveyInfo, error)
	SaveLatestSurveyInfo(l *model.LatestSurveyInfo) error
	GetSurvey(id string, version int) (*model.Survey, error)
	SaveSurvey(survey *model.Survey) error
	GetMeetingMetadata(meetingID string) (*model.MeetingMetadata, error)
	SaveMeetingMetadata(data *model.MeetingMetadata) error
	SaveSurveyResponse(response *model.SurveyResponse) error
}
