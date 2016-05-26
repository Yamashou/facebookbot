package topic

import (
	"github.com/m2mtu/facebookbot/learnword"
	"github.com/m2mtu/facebookbot/shiritori"
	"github.com/m2mtu/facebookbot/types"
	"github.com/m2mtu/facebookbot/reqCafe"
)

type empty struct {}

var topics []types.Topic

func init() {
	topics = []types.Topic{}
	regist(types.Topic{
		IsProper:         shiritori.IsProper,
		Talk:             shiritori.Talk,
		InitialTempState: shiritori.InitialTempState,
	})
	regist(types.Topic{
		IsProper:         learnword.IsProper,
		Talk:             learnword.Talk,
		InitialTempState: learnword.InitialTempState,
	})
	regist(types.Topic{
		IsProper: reqCafe.IsProper,
		Talk: reqCafe.Talk,
		InitialTempState: reqCafe.InitialTempState,
	})
}

// regist add new topic to topicModules. This function called from any topic packages.
func regist(_topic types.Topic) {
	topics = append(topics, _topic)
}

// GetCandidates returns topics possible to be demanded by user
func GetCandidates(content types.StaticState) []types.Topic {
	candidates := []types.Topic{}
	for _, _topic := range topics {
		if _topic.IsProper(content) {
			candidates = append(candidates, _topic)
		}
	}
	return candidates
}

// GetAllTopics return all topics registed
func GetAllTopics() []types.Topic {
	return topics
}
