package main

type GunAbstract struct {
	*Gun
	Abstract string `json:"abstract"`
}

type Gun struct {
	Ammo                         StringOrSlice          `json:"ammo"`
	AmmoEffects                  []interface{}          `json:"ammo_effects"`
	AmmoModifier                 []interface{}          `json:"ammo_modifier"`
	AmmoToFire                   float64                `json:"ammo_to_fire"`
	AmmoType                     []interface{}          `json:"ammo_type"`
	AoeIncrement                 float64                `json:"aoe_increment"`
	ArmorData                    map[string]interface{} `json:"armor_data"`
	AsciiPicture                 string                 `json:"ascii_picture"`
	BarrelLength                 string                 `json:"barrel_length"`
	BarrelVolume                 string                 `json:"barrel_volume"`
	BaseCastingTime              float64                `json:"base_casting_time"`
	BlackpowderTolerance         float64                `json:"blackpowder_tolerance"`
	BuiltInMods                  []interface{}          `json:"built_in_mods"`
	ChargesPerUse                float64                `json:"charges_per_use"`
	ClipSize                     float64                `json:"clip_size"`
	Color                        string                 `json:"color"`
	CoolingValue                 float64                `json:"cooling_value"`
	CopyFrom                     string                 `json:"copy-from"`
	DamageIncrement              float64                `json:"damage_increment"`
	DamageModifier               map[string]interface{} `json:"damage_modifier"`
	DamageType                   string                 `json:"damage_type"`
	DefaultMods                  []interface{}          `json:"default_mods"`
	Delete                       map[string]interface{} `json:"delete"`
	Description                  string                 `json:"description"`
	Difficulty                   float64                `json:"difficulty"`
	Dispersion                   float64                `json:"dispersion"`
	DispersionModifier           float64                `json:"dispersion_modifier"`
	Durability                   float64                `json:"durability"`
	Effect                       string                 `json:"effect"`
	EnergyDrain                  string                 `json:"energy_drain"`
	EnergySource                 string                 `json:"energy_source"`
	Extend                       map[string]interface{} `json:"extend"`
	Faults                       []interface{}          `json:"faults"`
	Flags                        StringOrSlice          `json:"flags"`
	Handling                     float64                `json:"handling"`
	HandlingModifier             float64                `json:"handling_modifier"`
	HeatPerShot                  float64                `json:"heat_per_shot"`
	Id                           string                 `json:"id"`
	InstallTime                  string                 `json:"install_time"`
	IntegralLongestSide          string                 `json:"integral_longest_side"`
	IntegralVolume               string                 `json:"integral_volume"`
	IntegralWeight               string                 `json:"integral_weight"`
	Location                     string                 `json:"location"`
	LongestSide                  string                 `json:"longest_side"`
	LooksLike                    string                 `json:"looks_like"`
	Loudness                     float64                `json:"loudness"`
	LoudnessModifier             float64                `json:"loudness_modifier"`
	Magazines                    []interface{}          `json:"magazines"`
	Material                     StringOrSlice          `json:"material"`
	MaxAoe                       float64                `json:"max_aoe"`
	MaxDamage                    float64                `json:"max_damage"`
	MaxLevel                     float64                `json:"max_level"`
	MaxRange                     float64                `json:"max_range"`
	MeleeDamage                  map[string]interface{} `json:"melee_damage"`
	Message                      string                 `json:"message"`
	MinAoe                       float64                `json:"min_aoe"`
	MinCycleRecoil               float64                `json:"min_cycle_recoil"`
	MinDamage                    float64                `json:"min_damage"`
	MinRange                     float64                `json:"min_range"`
	MinSkills                    []interface{}          `json:"min_skills"`
	ModTargets                   []interface{}          `json:"mod_targets"`
	ModeModifier                 []interface{}          `json:"mode_modifier"`
	Modes                        []interface{}          `json:"modes"`
	Name                         NameOrFucingStruct     `json:"name"`
	OverheatThreshold            float64                `json:"overheat_threshold"`
	OverwriteMinCycleRecoil      float64                `json:"overwrite_min_cycle_recoil"`
	PocketData                   []interface{}          `json:"pocket_data"`
	Price                        float64                `json:"price"`
	PricePostapoc                float64                `json:"price_postapoc"`
	Proportional                 map[string]interface{} `json:"proportional"`
	Qualities                    []interface{}          `json:"qualities"`
	Range                        float64                `json:"range"`
	RangeIncrement               float64                `json:"range_increment"`
	RangeModifier                float64                `json:"range_modifier"`
	RangedDamage                 map[string]interface{} `json:"ranged_damage"`
	Recoil                       float64                `json:"recoil"`
	Relative                     map[string]interface{} `json:"relative"`
	Reload                       float64                `json:"reload"`
	ReloadNoise                  string                 `json:"reload_noise"`
	ReloadNoiseVolume            float64                `json:"reload_noise_volume"`
	ReloadTime                   float64                `json:"reload_time"`
	Shape                        string                 `json:"shape"`
	ShotSpreadMultiplierModifier float64                `json:"shot_spread_multiplier_modifier"`
	SightDispersion              float64                `json:"sight_dispersion"`
	Skill                        string                 `json:"skill"`
	SpellClass                   string                 `json:"spell_class"`
	Symbol                       string                 `json:"symbol"`
	Techniques                   []interface{}          `json:"techniques"`
	ToHit                        ToHit                  `json:"to_hit"`
	Type                         string                 `json:"type"`
	UseAction                    map[string]interface{} `json:"use_action"`
	ValidModLocations            []interface{}          `json:"valid_mod_locations"`
	ValidTargets                 []interface{}          `json:"valid_targets"`
	VariantType                  string                 `json:"variant_type"`
	Variants                     []interface{}          `json:"variants"`
	Volume                       string                 `json:"volume"`
	WeaponCategory               []interface{}          `json:"weapon_category"`
	Weight                       string                 `json:"weight"`
}
