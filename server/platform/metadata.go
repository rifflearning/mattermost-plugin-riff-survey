package platform

import (
	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	"github.com/Brightscout/mattermost-plugin-survey/server/model"
)

// TODO: Remove this file and update all usages

// GetMeetingMetadata returns the survey metadata for a meeting with a given meetingID.
// Returns the meetingMetadata if found and nil if not.
func GetMeetingMetadata(meetingID string) *model.MeetingMetadata {
	meetingMetadata, err := config.Store.GetMeetingMetadata(meetingID)
	if err != nil {
		config.Mattermost.LogError("Unable to get meeting metadata.", "MeetingID", meetingID, "Error", err.Error())
		return nil
	}
	return meetingMetadata
}

// SaveMeetingMetadata saves the survey metadata for a meeting.
func SaveMeetingMetadata(meetingID, surveyID string, surveyVersion int) error {
	m := &model.MeetingMetadata{
		MeetingID:     meetingID,
		SurveyID:      surveyID,
		SurveyVersion: surveyVersion,
	}
	m = m.PreSave()
	if err := config.Store.SaveMeetingMetadata(m); err != nil {
		config.Mattermost.LogError("Failed to save the meeting metadata.", "MeetingID", meetingID, "SurveyID", surveyID, "SurveyVersion", surveyVersion, "Error", err.Error())
		return err
	}
	return nil
}

// GetUserMeetingMetadata returns the user metadata for a meeting with a given userID and meetingID.
// Returns the userMeetingMetadata if found and nil if not.
func GetUserMeetingMetadata(userID, meetingID string) *model.UserMeetingMetadata {
	userMeetingMetadata, err := config.Store.GetUserMeetingMetadata(userID, meetingID)
	if err != nil {
		config.Mattermost.LogError("Unable to get meeting metadata.", "UserID", userID, "MeetingID", meetingID, "Error", err.Error())
		return nil
	}
	return userMeetingMetadata
}

// SaveUserMeetingMetadata saves the user metadata for a meeting.
func SaveUserMeetingMetadata(u *model.UserMeetingMetadata) error {
	u = u.PreSave()
	if err := config.Store.SaveUserMeetingMetadata(u); err != nil {
		config.Mattermost.LogError("Failed to save the user meeting metadata.", "Error", err.Error())
	}
	return nil
}
