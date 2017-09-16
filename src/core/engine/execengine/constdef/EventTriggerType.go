package constdef

const (
	Engine_Unknown = iota
	Engine_Initialised
	Engine_BeforeStarted
	Engine_AfterStarted
	Engine_BeforeRun
	Engine_AfterRun
	Engine_Changed
)

var EventTriggerType = map[int]string{
	Engine_Unknown:       "Engine_Unknown",
	Engine_Initialised:   "Engine_Initialised",
	Engine_BeforeStarted: "Engine_BeforeStarted",
	Engine_AfterStarted:  "Engine_AfterStarted",
	Engine_BeforeRun:     "Engine_BeforeRun",
	Engine_AfterRun:      "Engine_AfterRun",
	Engine_Changed:       "Engine_Changed",
}
