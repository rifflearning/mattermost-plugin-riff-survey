package model

import (
	"encoding/json"
)

// LatestSurveyInfo stores the latest version information for a survey
type LatestSurveyInfo struct {
	Type          string `json:"type"`
	SurveyID      string `json:"survey_id"`
	SurveyVersion int    `json:"survey_version"`
}

func (info *LatestSurveyInfo) PreSave() *LatestSurveyInfo {
	info.Type = TypeLatestSurveyInfo
	return info
}

// EncodeToByte returns a LatestSurveyInfo as a byte array
func (info *LatestSurveyInfo) EncodeToByte() []byte {
	b, _ := json.Marshal(info)
	return b
}

// DecodeLatestSurveyInfoFromByte tries to create a LatestSurveyInfo from a byte array
func DecodeLatestSurveyInfoFromByte(b []byte) *LatestSurveyInfo {
	info := LatestSurveyInfo{}
	err := json.Unmarshal(b, &info)
	if err != nil {
		return nil
	}
	return &info
}
