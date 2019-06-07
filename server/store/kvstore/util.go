package kvstore

import (
	"fmt"

	"github.com/rifflearning/mattermost-plugin-survey/server/util"
)

const (
	latestSurveyKeyPrefix        = "latest_survey_"
	surveyKeyPrefix              = "survey_"
	surveyResponseKeyPrefix      = "survey_response_"
	meetingMetadataKeyPrefix     = "meeting_metadata_"
	reminderMetadataKeyPrefix    = "reminder_metadata_"
	userMeetingMetadataKeyPrefix = "user_meeting_metadata_"

	RemindersListKey = "reminders_list"
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

func ReminderMetadataKey(postID string) string {
	key := fmt.Sprintf("%s%s", reminderMetadataKeyPrefix, postID)
	return util.GetKeyHash(key)
}

func UserMeetingMetadataKey(userID, meetingID string) string {
	key := fmt.Sprintf("%s%s_%s", userMeetingMetadataKeyPrefix, userID, meetingID)
	return util.GetKeyHash(key)
}
