package modules

import (
	"Deadcord/core"
	"Deadcord/requests"
	"Deadcord/util"
	"strconv"
	"strings"
	"sync"

	"github.com/enescakir/emoji"
)

func StartReactThreads(channel_id string, message_id string, emoji string, suffix bool) {
	var wg sync.WaitGroup
	wg.Add(len(core.RawTokensLoaded))
	for _, token := range core.RawTokensLoaded {
		go func(token string, channel_id string, message_id string, emoji string, suffix bool) {
			reactWorker(token, channel_id, message_id, emoji, suffix)
		}(token, channel_id, message_id, emoji, suffix)
	}

	wg.Done()
}

func reactWorker(token string, channel_id string, message_id string, emoji_string string, suffix bool) {
	react_emoji := ""

	if suffix {
		react_emoji = strings.TrimSuffix(emoji.Parse(":"+emoji_string+":"), " ")
	} else {
		react_emoji = emoji_string
	}

	status, status_code, _ := requests.SendDiscordRequest("channels/"+channel_id+"/messages/"+message_id+"/reactions/"+react_emoji+"/@me", "PUT", token, map[string]interface{}{})

	if status {
		switch status_code {
		case 204:
			util.WriteToConsole("Token reacted with: [ "+react_emoji+" ].", 2)
		case 429:
			util.WriteToConsole("Reaction request was rate limited.", 1)
		default:
			util.WriteToConsole("Token could not react, request failed. Code: "+strconv.Itoa(status_code), 3)
		}
	}
}
