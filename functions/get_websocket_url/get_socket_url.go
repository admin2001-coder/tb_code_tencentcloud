package lib

import (
	"github.com/taubyte/go-sdk/event"
	"github.com/taubyte/go-sdk/pubsub"
)

//export getsocketurl
func getsocketurl(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	channel, err := pubsub.Channel("channelToMake")
	if err != nil {
		return 1
	}

	url, err := channel.WebSocket().Url()
	if err != nil {
		return 1
	}

	h.Write([]byte(url.Path))

	return 0
}
