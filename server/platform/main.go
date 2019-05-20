package platform

import (
	"github.com/Brightscout/mattermost-plugin-survey/server/config"
	"github.com/Brightscout/mattermost-plugin-survey/server/model"
)

// GetSurvey returns the survey with a given id and version.
// Returns the survey if found and nil if not.
func GetSurvey(id string, version int) *model.Survey {
	survey, err := config.Store.GetSurvey(id, version)
	if err != nil {
		config.Mattermost.LogError("Unable to get existing survey.", "Error", err.Error())
		return nil
	}
	return survey
}

// SaveSurvey creates a new survey in the DB.
// If this is the first survey for a given ID, the version is set to 1.
// Otherwise, the version for an existing survey is incremented.
func SaveSurvey(survey *model.Survey) error {
	currentSurveyVersion := 0

	if latestSurveyInfo := GetLatestSurveyInfo(survey.ID); latestSurveyInfo != nil {
		currentSurveyVersion = latestSurveyInfo.Version

		// Check for existing survey in DB
		if s := GetSurvey(survey.ID, currentSurveyVersion); s != nil && s.Equals(survey) {
			config.Mattermost.LogInfo("Survey already exists and is the same as the current survey. New version not created.", "SurveyID", s.ID, "SurveyVersion", s.Version)
			return nil
		}
	}

	// Create the first survey or new version of an existing survey
	survey = survey.PreSave(currentSurveyVersion + 1)
	if err := config.Store.SaveSurvey(survey); err != nil {
		config.Mattermost.LogError("Failed to save survey.", "Error", err.Error())
		return err
	}

	if err := SaveLatestSurveyInfo(survey.ID, survey.Version); err != nil {
		config.Mattermost.LogError("Survey saved successfully but latest survey information not updated.", "Error", err.Error())
		return err
	}

	return nil
}

// GetSurvey returns the latest survey information for the survey with a given id.
// Returns the info if found and nil if not.
func GetLatestSurveyInfo(id string) *model.LatestSurveyInfo {
	latestSurveyInfo, err := config.Store.GetLatestSurveyInfo(id)
	if err != nil {
		config.Mattermost.LogError("Unable to get latest survey information.", "Error", err.Error())
		return nil
	}
	return latestSurveyInfo
}

// SaveLatestSurveyInfo saves the latest survey information for a survey with a given id and version.
func SaveLatestSurveyInfo(id string, version int) error {
	info := &model.LatestSurveyInfo{
		ID:      id,
		Version: version,
	}
	info = info.PreSave()
	if err := config.Store.SaveLatestSurveyInfo(info); err != nil {
		config.Mattermost.LogError("Failed to save latest survey information.", "Error", err.Error())
	}
	return nil
}
