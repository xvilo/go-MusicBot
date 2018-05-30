package music

import "time"

type SongType string

const (
	SongTypeSong   SongType = "song"
	SongTypeStream SongType = "stream"
)

type Song struct {
	Name     string
	Artist   string
	Path     string
	SongType SongType
	Duration time.Duration
}
