package config

type ConfigStruct struct {
	LogLevel string       `cfg:"LOG_LEVEL"`
	Tts      TtsConfig    ``
	Devices  DeviceConfig ``
}

type TtsConfig struct {
	Lang string `cfg:"TTS_LANG"`
}

type DeviceConfig struct {
	ScanSecond        int      `cfg:"SCAN_SECOND"`
	ExcludeByInstance []string `cfg:"EXCLUDE_INSTANCE_,mergeslice"`
}
