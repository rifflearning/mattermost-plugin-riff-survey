package kvstore

import (
	"fmt"
)

const (
	LatestSurveyKeyPrefix    = "latest_survey_"
	SurveyKeyPrefix          = "survey_"
	SurveyResponseKeyPrefix  = "survey_response_"
	MeetingMetadataKeyPrefix = "meeting_metadata_"
)

func LatestSurveyKey(surveyID string) string {
	return fmt.Sprintf("%s%s", LatestSurveyKeyPrefix, surveyID)
}

func SurveyKey(surveyID, surveyVersion string) string {
	return fmt.Sprintf("%s%s_%s", SurveyKeyPrefix, surveyID, surveyVersion)
}

func SurveyResponseKey(userID, meetingID, surveyID, surveyVersion string) string {
	return fmt.Sprintf("%s%s_%s_%s_%s", SurveyResponseKeyPrefix, userID, meetingID, surveyID, surveyVersion)
}

func MeetingMetadataKey(meetingID string) string {
	return fmt.Sprintf("%s%s", MeetingMetadataKeyPrefix, meetingID)
}
