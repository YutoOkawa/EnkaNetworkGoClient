package pkg

type PlayerInfo struct {
	NickName             string `json:"nickName"`
	Level                int    `json:"level"`
	Signature            string `json:"signature"`
	WorldLevel           int    `json:"worldLevel"`
	NameCardId           int    `json:"nameCardId"`
	FinishArchivementNum int    `json:"finishArchivementNum"`
	TowerFloorIndex      int    `json:"towerFloorIndex"`
	TowerLevelIndex      int    `json:"towerLevelIndex"`
}
