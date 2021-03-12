package moduleconfig

type ModuleConfig struct {
	Buffer      int  `toml:"buffer" json:"buffer" yaml:"buffer"`
	AutoRestart bool `toml:"restart" json:"restart" yaml:"restart"`
}
