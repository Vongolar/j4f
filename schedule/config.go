package schedule

type moduleConfig struct {
	AutoRestart bool `json:"autoRestart" toml:"autoRestart" yaml:"autoRestart"`
	Buffer      int  `json:"buffer" toml:"buffer" yaml:"buffer"`
}
