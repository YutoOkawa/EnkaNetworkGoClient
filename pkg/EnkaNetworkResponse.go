package pkg

type EnkaNetworkResponse struct {
	PlayerInfo     PlayerInfo   `json:"playerInfo"`
	AvatarInfoList []AvatarInfo `json:"avatarInfoList"`
}
