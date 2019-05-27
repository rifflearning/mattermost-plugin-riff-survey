package kvstore

import (
	"fmt"

	"github.com/Brightscout/mattermost-plugin-survey/server/util"
)

const (
	latestSurveyKeyPrefix        = "latest_survey_"
	surveyKeyPrefix              = "survey_"
	surveyResponseKeyPrefix      = "survey_response_"
	meetingMetadataKeyPrefix     = "meeting_metadata_"
	userMeetingMetadataKeyPrefix = "user_meeting_metadata_"
)

func LatestSurveyKey(surveyID string) string {
	key := fmt.Sprintf("%s%s", latestSurveyKeyPrefix, surveyID)
	return util.GetKeyHash(key)
}

func SurveyKey(surveyID, surveyVersion string) string {
	key := fmt.Sprintf("%s%s_%s", surveyKeyPrefix, surveyID, surveyVersion)
	return util.GetKeyHash(key)
}

func SurveyResponseKey(userID, meetingID, surveyID, surveyVersion string) string {
	key := fmt.Sprintf("%s%s_%s_%s_%s", surveyResponseKeyPrefix, userID, meetingID, surveyID, surveyVersion)
	return util.GetKeyHash(key)
}

func MeetingMetadataKey(meetingID string) string {
	key := fmt.Sprintf("%s%s", meetingMetadataKeyPrefix, meetingID)
	return util.GetKeyHash(key)
}

func UserMeetingMetadataKey(userID, meetingID string) string {
	key := fmt.Sprintf("%s%s_%s", meetingMetadataKeyPrefix, userID, meetingID)
	return util.GetKeyHash(key)
}
