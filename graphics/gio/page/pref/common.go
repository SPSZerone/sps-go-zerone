package pref

const (
	TabIdxSettings = iota
	TabIdxPreferences
)

const (
	TabNameSettings    = "Settings"
	TabNamePreferences = "Preferences"
)

const (
	RatioKey   = 0.3
	RatioValue = 0.2
)

const (
	TableColIdxKey = iota
	TableColIdxValue
)

const (
	TableColWidthKey   = 220
	TableColWidthValue = 64
)

const (
	SettingsRowIdxNonModalDrawer = iota
	SettingsRowIdxTabAxis
	SettingsRowIdxBottomBar
	SettingsRowIdxDecorated
	SettingsRowIdxValueInFront
	SettingsRowCount
)

const (
	PreferencesRowIdxTableStyle = iota
	PreferencesRowCount
)
