package dota2

import (
	"github.com/mellaught/go-steam/steamid"
	bgcm "github.com/mellaught/go-dota2/protocol"
)

// InviteLobbyMember attempts to invite a player to the current lobby.
func (d *Dota2) InviteLobbyMember(playerID steamid.SteamId) {
	steamID := playerID.ToUint64()
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCInviteToLobby), &bgcm.CMsgInviteToLobby{
		SteamId: &steamID,
	})
}
