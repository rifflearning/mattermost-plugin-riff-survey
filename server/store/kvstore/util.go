package kvstore

import (
	"fmt"

	"github.com/Brightscout/mattermost-plugin-survey/server/util"
)

const (
	LatestSurveyKeyPrefix    = "latest_survey_"
	SurveyKeyPrefix          = "survey_"
	SurveyResponseKeyPrefix  = "survey_response_"
	MeetingMetadataKeyPrefix = "meeting_metadata_"
)

func LatestSurveyKey(surveyID string) string {
	key := fmt.Sprintf("%s%s", LatestSurveyKeyPrefix, surveyID)
	return util.GetKeyHash(key)
}

func SurveyKey(surveyID, surveyVersion string) string {
	key := fmt.Sprintf("%s%s_%s", SurveyKeyPrefix, surveyID, surveyVersion)
	return util.GetKeyHash(key)
}

func SurveyResponseKey(userID, meetingID, surveyID, surveyVersion string) string {
	key := fmt.Sprintf("%s%s_%s_%s_%s", SurveyResponseKeyPrefix, userID, meetingID, surveyID, surveyVersion)
	return util.GetKeyHash(key)
}

func MeetingMetadataKey(meetingID string) string {
	key := fmt.Sprintf("%s%s", MeetingMetadataKeyPrefix, meetingID)
	return util.GetKeyHash(key)
}
