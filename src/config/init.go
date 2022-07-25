package config

var LOG_LEVEL string

// 環境変数からconfigをセット
func Initialize() {
	LOG_LEVEL = getenv("LOG_LEVEL", "DEBUG")
}
