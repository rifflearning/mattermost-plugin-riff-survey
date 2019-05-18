package kvstore

import "github.com/Brightscout/mattermost-plugin-survey/server/model"

// import (
// 	"errors"

// 	"github.com/Brightscout/mattermost-plugin-survey/server/store"
// )

// SurveyStore allows to access surveys in the KV Store.
type SurveyStore struct {
}

func (s *SurveyStore) GetLatestSurveyVersion(id string) (*model.LatestSurveyVersion, error) {
	// TODO: Implement this method
	return nil, nil
}

func (s *SurveyStore) SaveLatestSurveyVersion(id, version string) error {
	// TODO: Implement this method
	return nil
}

func (s *SurveyStore) GetSurvey(id, version string) (*model.Survey, error) {
	// TODO: Implement this method
	return nil, nil
}

func (s *SurveyStore) SaveSurvey(survey *model.Survey) error {
	// TODO: Implement this method
	return nil
}

func (s *SurveyStore) GetMeetingMetadata(meetingID string) (*model.MeetingMetadata, error) {
	// TODO: Implement this method
	return nil, nil
}

func (s *SurveyStore) SaveMeetingMetadata(data *model.MeetingMetadata) error {
	// TODO: Implement this method
	return nil
}

func (s *SurveyStore) SaveSurveyResponse(response *model.SurveyResponse) error {
	// TODO: Implement this method
	return nil
}
