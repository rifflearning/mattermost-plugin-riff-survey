package platform

type SurveyApp interface {
	GetSurvey()
	GetDashboardLink()
	GetMeetingParticipants()
	SaveSurveyResponse()
	CreateSurveyPost()
	SelectSurvey()
}
