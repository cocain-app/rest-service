package main

type tomlConfig struct {
	Title   string
	Weights weights
}

//Weights config - is imported by calcTransScore() in calc.go
type weights struct {
	BPM float64 `toml:"bpm"`
	Key float64 `toml:"key"`
	R   float64 `toml:"reputation"`
	En  float64 `toml:"energy"`
	I   float64 `toml:"instrumental"`
	Da  float64 `toml:"danceability"`
	Lo  float64 `toml:"loudness"`
	V   float64 `toml:"valence"`
	T   float64 `toml:"timeSignature"`
	G   float64 `toml:"genre"`
	A   float64 `toml:"artist"`
	F   float64 `toml:"festival"`
	Du  float64 `toml:"duration"`
	Ex  float64 `toml:"existing"`
	Bl  float64 `toml:"blacklist"`
	Li  float64 `toml:"libary"`
}
