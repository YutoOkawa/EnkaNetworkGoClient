package model

type Equip struct {
	ItemID    int        `json:"itemId"`
	Weapon    *Weapon    `json:"weapon,omitempty"`
	Reliquary *Reliquary `json:"reliquary,omitempty"`
	Flat      Flat       `json:"flat"`
}

type Weapon struct {
	Level        int            `json:"level"`
	PromoteLevel int            `json:"promoteLevel"`
	AffixMap     map[string]int `json:"affixMap"`
}

type Reliquary struct {
	Level            int   `json:"level"`
	MainPropId       int   `json:"mainPropId"`
	AppendPropIdList []int `json:"appendPropIdList"`
}

type Flat struct {
	NameTextMapHash    string              `json:"nameTextMapHash"`
	SetNameTextMapHash string              `json:"setNameTextMapHash,omitempty"`
	RankLevel          int                 `json:"rankLevel"`
	ReliquaryMainstat  *ReliquaryMainstat  `json:"reliquaryMainstat,omitempty"`
	ReliquarySubstats  []*ReliquarySubstat `json:"reliquarySubstats,omitempty"`
	WeaponStats        []*WeaponStat       `json:"weaponStats,omitempty"`
	ItemType           string              `json:"itemType"`
	Icon               string              `json:"icon"`
	EquipType          string              `json:"equipType,omitempty"`
}

type ReliquaryMainstat struct {
	MainPropId string  `json:"mainPropId"`
	StatValue  float64 `json:"statValue"`
}

type ReliquarySubstat struct {
	AppendPropId string  `json:"appendPropId"`
	StatValue    float64 `json:"statValue"`
}

type WeaponStat struct {
	AppendPropId string  `json:"appendPropId"`
	StatValue    float64 `json:"statValue"`
}
