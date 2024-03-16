package main

type Gunmod struct {
	ShotSpreadMultiplierModifier float64                `json:"shot_spread_multiplier_modifier"`
	Description                  string                 `json:"description"`
	Weight                       string                 `json:"weight"`
	Type                         string                 `json:"type"`
	WeightMultiplier             float64                `json:"weight_multiplier"`
	PricePostapoc                float64                `json:"price_postapoc"`
	BlacklistMod                 []interface{}          `json:"blacklist_mod"`
	Symbol                       string                 `json:"symbol"`
	AmmoEffects                  []interface{}          `json:"ammo_effects"`
	ToHit                        float64                `json:"to_hit"`
	Extend                       map[string]interface{} `json:"extend"`
	Color                        string                 `json:"color"`
	Id                           string                 `json:"id"`
	IntegralWeight               string                 `json:"integral_weight"`
	BarrelLength                 string                 `json:"barrel_length"`
	BlacklistSlot                []interface{}          `json:"blacklist_slot"`
	Name                         map[string]interface{} `json:"name"`
	DispersionModifier           float64                `json:"dispersion_modifier"`
	MagazineAdaptor              []interface{}          `json:"magazine_adaptor"`
	AimSpeedModifier             float64                `json:"aim_speed_modifier"`
	UseAction                    map[string]interface{} `json:"use_action"`
	LongestSide                  string                 `json:"longest_side"`
	GunmodData                   map[string]interface{} `json:"gunmod_data"`
	Volume                       string                 `json:"volume"`
	IntegralVolume               string                 `json:"integral_volume"`
	GunData                      map[string]interface{} `json:"gun_data"`
	FieldOfView                  float64                `json:"field_of_view"`
	Price                        float64                `json:"price"`
	MinSkills                    []interface{}          `json:"min_skills"`
	EnergyDrainMultiplier        float64                `json:"energy_drain_multiplier"`
	ConsumeDivisor               float64                `json:"consume_divisor"`
	DamageModifier               map[string]interface{} `json:"damage_modifier"`
	ArmorData                    map[string]interface{} `json:"armor_data"`
	Faults                       []interface{}          `json:"faults"`
	PocketMods                   []interface{}          `json:"pocket_mods"`
	HandlingModifier             float64                `json:"handling_modifier"`
	ConsumeChance                float64                `json:"consume_chance"`
	AddMod                       []interface{}          `json:"add_mod"`
	Proportional                 map[string]interface{} `json:"proportional"`
	MeleeDamage                  map[string]interface{} `json:"melee_damage"`
	Delete                       map[string]interface{} `json:"delete"`
	IntegralLongestSide          string                 `json:"integral_longest_side"`
	InstallTime                  string                 `json:"install_time"`
	Material                     []interface{}          `json:"material"`
	CopyFrom                     string                 `json:"copy-from"`
	SightDispersion              float64                `json:"sight_dispersion"`
	AmmoModifier                 []interface{}          `json:"ammo_modifier"`
	Flags                        []interface{}          `json:"flags"`
	Location                     string                 `json:"location"`
	LoudnessModifier             float64                `json:"loudness_modifier"`
	ModeModifier                 []interface{}          `json:"mode_modifier"`
	Qualities                    []interface{}          `json:"qualities"`
	PocketData                   []interface{}          `json:"pocket_data"`
	RangeModifier                float64                `json:"range_modifier"`
	RangeMultiplier              float64                `json:"range_multiplier"`
	ModTargets                   []interface{}          `json:"mod_targets"`
	AcceptableAmmo               []interface{}          `json:"acceptable_ammo"`
}

type ModLocations struct {
	barrel        int
	bore          int
	brass_catcher int
	grip          int
	mechanism     int
	magazine      int
	muzzle        int
	rail          int
	sights        int
	stock         int
	underbarrel   int
}
