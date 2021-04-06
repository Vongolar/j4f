package mhttp

type config struct {
	Address string `toml:"address" json:"address" yaml:"address"`
	Cert    string `toml:"certFile" json:"certFile" yaml:"certFile"`
	Key     string `toml:"keyFile" json:"keyFile" yaml:"keyFile"`
}
