package reply

import (
	"reflect"

	"github.com/Yamashou/facebookbot/state"
	"github.com/Yamashou/facebookbot/talk"
	"github.com/Yamashou/facebookbot/topic"
	"github.com/Yamashou/facebookbot/types"
)

// Talk method talk with user
func Talk(receivedEvent talk.Event) {
	staticState := types.StaticState{}
	staticState.OpponentID = receivedEvent.SenderID
	staticState.EndPointName = talk.GetEndPointName()
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
		var tempStateRfValue reflect.Value
		var permStateRfValue reflect.Value
		theTopic := staticState.PossibleTopics[0]
		initialTempStateValue := reflect.ValueOf(theTopic.InitialTempState)
		tempState, ok := state.Temp(userID)
		if ok {
			tempStateRfValue = reflect.ValueOf(tempState)
		} else {
			tempStateRfValue = initialTempStateValue.Call([]reflect.Value{})[0]
		}
		permState, ok := state.Perm(userID)
		if ok {
			permStateRfValue = reflect.ValueOf(permState)
		} else {
			permStateRfValue = reflect.ValueOf(state.InitialPerm())
		}
		talkvalue := reflect.ValueOf(theTopic.Talk)
		var results []reflect.Value
		typedTempState := tempStateRfValue.Convert(talkvalue.Type().In(1))
		results = talkvalue.Call([]reflect.Value{reflect.ValueOf(staticState), typedTempState, permStateRfValue})
		newTempStateRfValue := results[0]
		newPermState, ok := results[1].Interface().(types.PermState)
		if !ok {
			panic("Type Error: The second argument type of Talk method must be types.PermState.")
		}
		willTopicContinue := results[2].Bool()
		if willTopicContinue {
			reflect.ValueOf(state.SetTemp).Call([]reflect.Value{reflect.ValueOf(userID), newTempStateRfValue})
		} else {
			state.UnsetTemp(userID)
		}
		state.SetPerm(userID, newPermState)
		if !willTopicContinue {
			staticState.PossibleTopics = topic.GetAllTopics()
		}
	}
	state.SetStatic(userID, staticState)
}
