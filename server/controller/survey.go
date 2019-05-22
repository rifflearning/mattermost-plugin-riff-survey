package controller

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/Brightscout/mattermost-plugin-survey/server/platform"
)

var getSurvey = &Endpoint{
	Path:         "/survey",
	Method:       http.MethodGet,
	Execute:      executeGetSurvey,
	RequiresAuth: true,
}

func executeGetSurvey(w http.ResponseWriter, r *http.Request) error {
	surveyID := r.URL.Query().Get("survey_id")
	surveyVersion := r.URL.Query().Get("survey_version")
	surveyVersionInt, err := strconv.Atoi(surveyVersion)
	if err != nil {
		http.Error(w, "Invalid Survey Version.", http.StatusBadRequest)
		return errors.Wrap(err, "invalid survey version")
	}

	// TODO: Get survey questions using meetingID instead
	survey := platform.GetSurvey(surveyID, surveyVersionInt)
	if survey == nil {
		http.Error(w, "Unable to get survey for requested id and version.", http.StatusBadRequest)
		return errors.Wrap(err, "unable to get survey for requested id and version")
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(survey.EncodeToByte()); err != nil {
		return errors.Wrap(err, "failed to write data to HTTP response")
	}

	return nil
}
