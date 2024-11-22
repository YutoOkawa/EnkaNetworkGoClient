package pkg

type AvatarInfo struct {
	AvatarID              int             `json:"avatarId"`
	TalentIdList          []int           `json:"talentIdList,omitempty"`
	PropMap               map[string]Prop `json:"propMap"`
	FightProp             map[string]int  `json:"fightProp"`
	SkillDepotId          int             `json:"skillDepotId"`
	InheritProudSkillList []int           `json:"inheritProudSkillList"`
	SkillLevelMap         map[string]int  `json:"skillLevelMap"`
	EquipList             []Equip         `json:"equipList"`
	FetterInfo            map[string]int  `json:"fetterInfo"`
}

type Prop struct {
	Type int    `json:"type"`
	Ival string `json:"ival"`
	Val  string `json:"val"`
}
