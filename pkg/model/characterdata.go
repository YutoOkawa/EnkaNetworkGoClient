package model

type CharacterData struct {
	Element         string             `json:"element"`
	CharacterImage  []string           `json:"Consts"`
	SkillOrder      []int              `json:"SkillOrder"`
	Skills          map[string]string  `json:"Skills"`
	ProudMap        map[string]int     `json:"ProudMap"`
	NameTextMapHash int                `json:"NameTextMapHash"`
	SideIconName    string             `json:"SideIconName"`
	QualityType     string             `json:"QualityType"`
	WeaponType      string             `json:"WeaponType"`
	Costumes        map[string]Costume `json:"Costumes,omitempty"`
}

type Costume struct {
	SideIconName string `json:"sideIconName"`
	Icon         string `json:"icon"`
	Art          string `json:"art"`
	AvatarId     int    `json:"avatarId"`
}
