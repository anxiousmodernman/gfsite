package models

import (
	"errors"
)

type Episode struct {
	episodeNumber int
	summary       string
	title         string
	notes         string
}

func (this *Episode) EpisodeNumber() int {
	return this.episodeNumber
}
func (this *Episode) Summary() string {
	return this.summary
}
func (this *Episode) Title() string {
	return this.title
}
func (this *Episode) Notes() string {
	return this.notes
}

func GetEpisodes() (episodes []Episode, err error) {
	db, err := getDBConnection()
	// episodes = 
	if err == nil {
		defer db.Close()


	} else {
		episodes, errors.New("Unable to connect to the database")
	}
}