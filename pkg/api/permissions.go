package api

const (
	PermissionAdministrator       Permission = 8
	PermissionViewAuditLog        Permission = 128
	PermissionViewServerInsights  Permission = 524288
	PermissionManageServer        Permission = 32
	PermissionManageRoles         Permission = 268435456
	PermissionManageChannels      Permission = 16
	PermissionKickMembers         Permission = 2
	PermissionBanMembers          Permission = 4
	PermissionCreateInstantInvite Permission = 1
	PermissionChangeNickname      Permission = 67108864
	PermissionManageNicknames     Permission = 134217728
	PermissionManageEmojis        Permission = 1073741824
	PermissionManageWebhooks      Permission = 536870912
	PermissionViewChannels        Permission = 1024

	PermissionMessagesSend              Permission = 2048
	PermissionMessagesSendTTS           Permission = 4096
	PermissionMessagesManage            Permission = 8192
	PermissionMessagesEmbedLinks        Permission = 16384
	PermissionMessagesAttachFiles       Permission = 32768
	PermissionMessagesReadHistory       Permission = 65536
	PermissionMessagesMentionEveryone   Permission = 131072
	PermissionMessagesUseExternalEmojis Permission = 262144
	PermissionMessagesAddReactions      Permission = 64

	PermissionVoiceConnect         Permission = 1048576
	PermissionVoiceSpeak           Permission = 2097152
	PermissionVoiceMuteMembers     Permission = 4194304
	PermissionVoiceDeafenMembers   Permission = 8388608
	PermissionVoiceUseMembers      Permission = 16777216
	PermissionVoiceUseActivity     Permission = 33554432
	PermissionVoicePrioritySpeaker Permission = 256
)