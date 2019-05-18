package kvstore

import (
	"github.com/Brightscout/mattermost-plugin-survey/server/store"
)

// Store is an interface to interact with the KV Store.
type Store struct {
	surveyStore SurveyStore
}

// NewStore returns a fresh store and upgrades the db from the given schema version.
func NewStore() store.Store {
	return &Store{
		surveyStore: SurveyStore{},
	}
}

// Survey returns the Survey Store
func (s *Store) Survey() store.SurveyStore { return &s.surveyStore }
