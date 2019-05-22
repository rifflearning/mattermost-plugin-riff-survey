package kvstore

import (
	"strconv"

	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	"github.com/Brightscout/mattermost-plugin-survey/server/model"
	"github.com/Brightscout/mattermost-plugin-survey/server/store"
)

// Store is the implementation for the interface to interact with the KV Store.
type Store struct {
}

// NewStore returns a fresh store.
func NewStore() store.Store {
	return &Store{}
}

func (s *Store) GetLatestSurveyInfo(id string) (*model.LatestSurveyInfo, error) {
	key := LatestSurveyKey(id)
	b, err := config.Mattermost.KVGet(key)
	if err != nil {
		return nil, err
	}
	info := model.DecodeLatestSurveyInfoFromByte(b)
	return info, nil
}

func (s *Store) SaveLatestSurveyInfo(info *model.LatestSurveyInfo) error {
	key := LatestSurveyKey(info.ID)
	if err := config.Mattermost.KVSet(key, info.EncodeToByte()); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetSurvey(id string, version int) (*model.Survey, error) {
	key := SurveyKey(id, strconv.Itoa(version))
	b, err := config.Mattermost.KVGet(key)
	if err != nil {
		return nil, err
	}
	survey := model.DecodeSurveyFromByte(b)
	return survey, nil
}

func (s *Store) SaveSurvey(survey *model.Survey) error {
	key := SurveyKey(survey.ID, strconv.Itoa(survey.Version))
	if err := config.Mattermost.KVSet(key, survey.EncodeToByte()); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetMeetingMetadata(meetingID string) (*model.MeetingMetadata, error) {
	// TODO: Implement this method
	return nil, nil
}

func (s *Store) SaveMeetingMetadata(data *model.MeetingMetadata) error {
	// TODO: Implement this method
	return nil
}

func (s *Store) SaveSurveyResponse(response *model.SurveyResponse) error {
	key := SurveyResponseKey(response.UserID, response.MeetingID, response.SurveyID, strconv.Itoa(response.SurveyVersion))
	if err := config.Mattermost.KVSet(key, response.EncodeToByte()); err != nil {
		return err
	}
	return nil
}
