package types

import (
	cabocha "github.com/ledyba/go-cabocha"
)

// Topic express each topic.
type Topic struct {
	IsProper         func(StaticState) bool
	Talk             interface{}
	InitialTempState interface{}
}

// TopicModule express each topic. This is written to avoid import cycle between topic and state.
type TopicModule interface {
	IsProper(EndPointContent) bool
}

// State contains all states.
type State struct {
	Static StaticState
	Temp   TempState
	Perm   PermState
}

// StaticState contains various short-term data about user and conversation running.
type StaticState struct {
	PossibleTopics  []Topic
	OpponentID      UserID
	EndPointName    string
	ReceivedContent EndPointContent
}

// TempState is temporaly state.
type TempState interface{}

// Noun express each noun.
type Noun struct {
	Token       cabocha.Token
	Description string
}

// PermState is permutable state.
type PermState struct {
	LearnedNouns []Noun
}

// EndPointContent express each message content.
type EndPointContent interface{}

// UserID express User ID
type UserID string

func (userid UserID) String() string {
	return string(userid)
}
