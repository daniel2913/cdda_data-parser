package main

type Ammo struct {
	Abstract        string                 `json:"abstract"`
	AmmoType        string                 `json:"ammo_type"`
	AsciiPicture    string                 `json:"ascii_picture"`
	Casing          string                 `json:"casing"`
	Category        string                 `json:"category"`
	Color           string                 `json:"color"`
	Container       string                 `json:"container"`
	CopyFrom        string                 `json:"copy-from"`
	Count           float64                `json:"count"`
	Damage          map[string]interface{} `json:"damage"`
	Delete          map[string]interface{} `json:"delete"`
	Description     string                 `json:"description"`
	Dispersion      float64                `json:"dispersion"`
	Drop            string                 `json:"drop"`
	Effects         []interface{}          `json:"effects"`
	ExplodeInFire   bool                   `json:"explode_in_fire"`
	Explosion       map[string]interface{} `json:"explosion"`
	Extend          map[string]interface{} `json:"extend"`
	Flags           StringOrSlice          `json:"flags"`
	Id              string                 `json:"id"`
	LongestSide     string                 `json:"longest_side"`
	LooksLike       string                 `json:"looks_like"`
	Loudness        float64                `json:"loudness"`
	Material        Material               `json:"material"`
	MeleeDamage     map[string]interface{} `json:"melee_damage"`
	Name            NameOrFucingStruct     `json:"name"`
	Phase           string                 `json:"phase"`
	Price           float64                `json:"price"`
	PricePostapoc   float64                `json:"price_postapoc"`
	ProjectileCount float64                `json:"projectile_count"`
	Proportional    map[string]interface{} `json:"proportional"`
	Range           float64                `json:"range"`
	Recoil          float64                `json:"recoil"`
	Relative        map[string]interface{} `json:"relative"`
	Sealed          bool                   `json:"sealed"`
	ShotDamage      map[string]interface{} `json:"shot_damage"`
	ShotSpread      float64                `json:"shot_spread"`
	ShowStats       bool                   `json:"show_stats"`
	StackSize       float64                `json:"stack_size"`
	Symbol          string                 `json:"symbol"`
	ToHit           ToHit                  `json:"to_hit"`
	Type            string                 `json:"type"`
	UseAction       map[string]interface{} `json:"use_action"`
	Volume          string                 `json:"volume"`
	Weight          string                 `json:"weight"`
}
