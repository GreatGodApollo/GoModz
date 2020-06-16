package api

import "github.com/bwmarrin/discordgo"

const (
	PluginsDir		= "./plugins"
	ModSymbolName 	= "Mod"
)

func CheckSessionPermissions(s *discordgo.Session, guildid, memberid string, required Permission) bool {
	if required == 0 {
		return true
	}

	guild, err := s.State.Guild(guildid)
	if err != nil {
		return false
	}

	if guild.OwnerID == memberid {
		return true
	}

	member, err := s.State.Member(guildid, memberid)
	if err != nil {
		return false
	}

	var perms int
	for _, roleid := range member.Roles {
		role, err := s.State.Role(guildid, roleid)
		if err != nil {
			return false
		}

		if perms & (role.Permissions) == 0 {
			perms = perms | role.Permissions
		}

		if role.Permissions & int(PermissionAdministrator) != 0 {
			return true
		}
	}

	if perms & int(required) == int(required) {
		return true
	}

	return false
}