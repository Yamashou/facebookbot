package types

// Topic express each topic. This is written to avoid import cycle between topic and state.
type Topic interface {
	IsProper(EndPointContent) bool
	Talk(StaticState, TempState, PermState) (TempState, PermState)
	InitialTempState() TempState
	InitialPermState() PermState
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

// PermState is permutation state.
type PermState interface{}

// EndPointContent express each message content.
type EndPointContent interface{}

// UserID express User ID
type UserID string

func (userid UserID) String() string {
	return string(userid)
}
