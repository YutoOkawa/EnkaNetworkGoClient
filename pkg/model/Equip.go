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

const (
	ITEM_TYPE_WEAPON   = "ITEM_WEAPON"
	ITEM_TYPE_ARTIFACT = "ITEM_RELIQUARY"
)

const (
	EQUIP_TYPE_FLOWER  = "EQUIP_BRACER"
	EQUIP_TYPE_FEATHER = "EQUIP_NECKLACE"
	EQUIP_TYPE_SANDS   = "EQUIP_SHOES"
	EQUIP_TYPE_GOBLET  = "EQUIP_RING"
	EQUIP_TYPE_CIRCLET = "EQUIP_DRESS"
)

const (
	APPENDPROP_BASE_ATK              = "FIGHT_PROP_BASE_ATTACK"
	APPENDPROP_HP                    = "FIGHT_PROP_HP"
	APPENDPROP_ATK                   = "FIGHT_PROP_ATTACK"
	APPENDPROP_DEF                   = "FIGHT_PROP_DEFENSE"
	APPENDPROP_HP_PERCENTAGE         = "FIGHT_PROP_HP_PERCENTAGE"
	APPENDPROP_ATK_PERCENTAGE        = "FIGHT_PROP_ATTACK_PERCENTAGE"
	APPENDPROP_DEF_PERCENTAGE        = "FIGHT_PROP_DEFENSE_PERCENTAGE"
	APPENDPROP_CRITICAL_RATE         = "FIGHT_PROP_CRITICAL"
	APPENDPROP_CRITICAL_DAMAGE       = "FIGHT_PROP_CRITICAL_HURT"
	APPENDPROP_ENERGY_RECHARGET      = "FIGHT_PROP_CHARGE_EFFICIENCY"
	APPENDPROP_HEADLING_BONUS        = "FIGHT_PROP_HEAL_ADD"
	APPENDPROP_ELEMENTAL_MASTERY     = "FIGHT_PROP_ELEMENT_MASTERY"
	APPENDPROP_PHYSICAL_DAMAGE_BONUS = "FIGHT_PROP_PHYSICAL_ADD_HURT"
	APPENDPROP_PYRO_DAMAGE_BONUS     = "FIGHT_PROP_FIRE_ADD_HURT"
	APPENDPROP_ELECTRO_DAMAGE_BONUS  = "FIGHT_PROP_ELEC_ADD_HURT"
	APPENDPROP_HYDRO_DAMAGE_BONUS    = "FIGHT_PROP_WATER_ADD_HURT"
	APPENDPROP_ANEMO_DAMAGE_BONUS    = "FIGHT_PROP_WIND_ADD_HURT"
	APPENDPROP_CRYO_DAMAGE_BONUS     = "FIGHT_PROP_ICE_ADD_HURT"
	APPENDPROP_GEO_DAMAGE_BONUS      = "FIGHT_PROP_ROCK_ADD_HURT"
	APPENDPROP_DENDRO_DAMAGE_BONUS   = "FIGHT_PROP_GRASS_ADD_HURT"
)
