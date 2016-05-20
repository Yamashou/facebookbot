package state

import (
	"github.com/m2mtu/facebookbot/types"
)

var statics map[types.UserID]types.StaticState
var perms map[types.UserID]types.PermState
var temps map[types.UserID]types.TempState

func init() {
	statics = make(map[types.UserID]types.StaticState)
	perms = make(map[types.UserID]types.PermState)
	temps = make(map[types.UserID]types.TempState)
}

// Static return state by user
func Static(userid types.UserID) (types.StaticState, bool) {
	_state, ok := statics[userid]
	return _state, ok
}

// SetStatic set state
func SetStatic(userid types.UserID, _state types.StaticState) {
	statics[userid] = _state
}

// Temp return state by user
func Temp(userid types.UserID) (types.TempState, bool) {
	_state, ok := temps[userid]
	return _state, ok
}

// SetTemp set state
func SetTemp(userid types.UserID, _state types.TempState) {
	temps[userid] = _state
}

// Perm return state by user
func Perm(userid types.UserID) (types.PermState, bool) {
	_state, ok := perms[userid]
	return _state, ok
}

// SetPerm set state
func SetPerm(userid types.UserID, _state types.PermState) {
	perms[userid] = _state
}

// InitialPerm method returns initial object of PermState.
func InitialPerm() types.PermState {
	return types.PermState{
		LearnedNouns: []types.Noun{
			types.Noun{
				Text: "ごりら",
				Description: "動物",
			},
		},
	}
}
