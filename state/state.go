package state

// Topic express each topic function.
type Topic struct {
	ID int64
}

// State contains various short-term data about user and conversation running.
type State struct {
	PossibleTopics []Topic
	UserID string
	EndPointName string
	ReceivedContents []interface{}
	Memory interface{}
}
