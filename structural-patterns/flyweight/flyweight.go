package flyweight

import (
	"time"
)

type Team struct {
	ID uint64
	Name uint64
	Shield []byte
	Players []Player
	HistoricalData []HistoricalData
}

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name string
	Surname string
	PreviousTeam uint64
	Photo []byte
}

type HistoricalData struct {
	Year uint8
	LeagueResults []Match
}

type Match struct {
	Date time.Time
	VisitorID uint64
	LocalID uint64
	LocalScore byte
	VisitorScore byte
	LocalShoots uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct {
	createdTeams map[uint64]*Team
}

func (t *teamFlyweightFactory) GetTeam(id uint64) *Team {
	if t.createdTeams == nil {
		t.createdTeams = make(map[uint64]*Team)
	}
	if t.createdTeams[id] != nil {
		return t.createdTeams[id]
	}

	team := getTeamFactory(id)
	t.createdTeams[id] = &team
	return t.createdTeams[id]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(team uint64) Team {
	switch team {
	case TEAM_B:
	return Team{
		ID: 2,
		Name: TEAM_B,
	}
	default:
	return Team{
		ID: 1,
		Name: TEAM_A,
	}
	}
}