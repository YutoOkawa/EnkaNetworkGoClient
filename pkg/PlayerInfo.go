package pkg

type PlayerInfo struct {
	NickName             string           `json:"nickName"`
	Level                int              `json:"level"`
	Signature            string           `json:"signature"`
	WorldLevel           int              `json:"worldLevel"`
	NameCardId           int              `json:"nameCardId"`
	FinishArchivementNum int              `json:"finishArchivementNum"`
	TowerFloorIndex      int              `json:"towerFloorIndex"`
	TowerLevelIndex      int              `json:"towerLevelIndex"`
	ShowAvatarInfoList   []ShowAvatarInfo `json:"showAvatarInfoList"`
	ShowNameCardIdList   []int            `json:"showNameCardIdList"`
	ProfilePicture       ProfilePicture   `json:"profilePicture"`
	IsShowAvatarTalent   bool             `json:"isShowAvatarTalent"`
	FetterCount          int              `json:"fetterCount"`
	TowerStarIndex       int              `json:"towerStarIndex"`
}
