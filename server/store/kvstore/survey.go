package kvstore

import (
	"strconv"

	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	"github.com/Brightscout/mattermost-plugin-survey/server/model"
	"github.com/Brightscout/mattermost-plugin-survey/server/store"
)

// KVStore is the implementation for the SurveyStore interface using the plugin KV Store.
type KVStore struct {
}

// NewStore returns a fresh store.
func NewStore() store.SurveyStore {
	return &KVStore{}
}

func (s *KVStore) GetLatestSurveyInfo(id string) (*model.LatestSurveyInfo, error) {
	key := LatestSurveyKey(id)
	b, err := config.Mattermost.KVGet(key)
	if err != nil {
		return nil, err
	}
	info := model.DecodeLatestSurveyInfoFromByte(b)
	return info, nil
}

func (s *KVStore) SaveLatestSurveyInfo(info *model.LatestSurveyInfo) error {
	key := LatestSurveyKey(info.SurveyID)
	if err := config.Mattermost.KVSet(key, info.EncodeToByte()); err != nil {
		return err
	}
	return nil
}

func (s *KVStore) GetSurvey(id string, version int) (*model.Survey, error) {
	key := SurveyKey(id, strconv.Itoa(version))
	b, err := config.Mattermost.KVGet(key)
	if err != nil {
		return nil, err
	}
	survey := model.DecodeSurveyFromByte(b)
	return survey, nil
}

func (s *KVStore) SaveSurvey(survey *model.Survey) error {
	key := SurveyKey(survey.ID, strconv.Itoa(survey.Version))
	if err := config.Mattermost.KVSet(key, survey.EncodeToByte()); err != nil {
		return err
	}
	return nil
}

func (s *KVStore) GetMeetingMetadata(meetingID string) (*model.MeetingMetadata, error) {
	key := MeetingMetadataKey(meetingID)
	b, err := config.Mattermost.KVGet(key)
	if err != nil {
		return nil, err
	}
	m := model.DecodeMeetingMetadataFromByte(b)
	return m, nil
}

func (s *KVStore) SaveMeetingMetadata(data *model.MeetingMetadata) error {
	key := MeetingMetadataKey(data.MeetingID)
	if err := config.Mattermost.KVSet(key, data.EncodeToByte()); err != nil {
		return err
	}
	return nil
}

func (s *KVStore) SaveSurveyResponse(response *model.SurveyResponse) error {
	key := SurveyResponseKey(response.UserID, response.MeetingID, response.SurveyID, strconv.Itoa(response.SurveyVersion))
	if err := config.Mattermost.KVSet(key, response.EncodeToByte()); err != nil {
		return err
	}
	return nil
}
