package dota2

import (
	gcmcc "github.com/mellaught/go-dota2/protocol"
	gcm "github.com/mellaught/go-dota2/protocol"
)

// SendChannelMessage attempts to send a message in a channel, text-only.
// Use SendChatMessage for more fine-grained control.
func (d *Dota2) SendChannelMessage(channelID uint64, message string) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCChatMessage), &gcmcc.CMsgDOTAChatMessage{
		ChannelId: &channelID,
		Text:      &message,
	})
}
