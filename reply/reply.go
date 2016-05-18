package reply

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
	"reflect"

	"github.com/m2mtu/facebookbot/MyClassSearch"
	"github.com/m2mtu/facebookbot/MyStudyRoomSearch"
	"github.com/m2mtu/facebookbot/RandomWord"
	"github.com/m2mtu/facebookbot/SearchFreeRoom"
	"github.com/m2mtu/facebookbot/infoSub"
	"github.com/m2mtu/facebookbot/reqCafe"
	"github.com/m2mtu/facebookbot/endpoint"
	"github.com/m2mtu/facebookbot/types"
	"github.com/m2mtu/facebookbot/state"
	"github.com/m2mtu/facebookbot/topic"
)

// Talk method talk with user
func Talk(receivedEvent endpoint.Event) {
	staticState := types.StaticState{}
	staticState.OpponentID = receivedEvent.SenderID
	staticState.EndPointName = endpoint.GetEndPointName()
	staticState.ReceivedContent = receivedEvent.Content
	userID := receivedEvent.SenderID
	lastStaticState, ok := state.Static(userID)
	if ok && len(lastStaticState.PossibleTopics) == 1 {
		staticState.PossibleTopics = lastStaticState.PossibleTopics
	} else if ok {
		tempStaticState := staticState
		tempStaticState.PossibleTopics = lastStaticState.PossibleTopics
		staticState.PossibleTopics = topic.GetCandidates(tempStaticState)
	} else {
		tempStaticState := staticState
		tempStaticState.PossibleTopics = topic.GetAllTopics()
		staticState.PossibleTopics = topic.GetCandidates(tempStaticState)
	}

	if len(staticState.PossibleTopics) == 1 {
		theTopic := staticState.PossibleTopics[0]
		tempState, ok1 := state.Temp(userID)
		permState, ok2 := state.Perm(userID)
		initialPermState := state.InitialPerm()
		talkvalue := reflect.ValueOf(theTopic.Talk)
		var results []reflect.Value
		if !ok1 || !ok2 {
			results = talkvalue.Call(
				[]reflect.Value{
					reflect.ValueOf(staticState),
					reflect.New(talkvalue.Type().In(1)).Elem(),
					reflect.ValueOf(initialPermState),
				},
			)
		} else {
			typedTempState := reflect.ValueOf(tempState).Convert(talkvalue.Type().In(1))
			results = talkvalue.Call([]reflect.Value{reflect.ValueOf(staticState), typedTempState, reflect.ValueOf(permState)})
		}
		newTempStateRfValue := results[0]
		newPermState, ok := results[1].Interface().(types.PermState)
		if !ok {
			panic("Type Error: The second argument type of Talk method must be types.PermState.")
		}
		willTopicContinue := results[2].Bool()
		if willTopicContinue {
			reflect.ValueOf(state.SetTemp).Call([]reflect.Value{reflect.ValueOf(userID), newTempStateRfValue})
		} else {
			reflect.ValueOf(state.SetTemp).Call([]reflect.Value{reflect.ValueOf(userID), reflect.New(newTempStateRfValue.Type()).Elem()})
		}
		state.SetPerm(userID, newPermState)
		if !willTopicContinue {
			staticState.PossibleTopics = topic.GetAllTopics()
		}
	}
	state.SetStatic(userID, staticState)
}
