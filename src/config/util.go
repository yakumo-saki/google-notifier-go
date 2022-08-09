package config

import "strings"

func IsIgnoredInstance(instance string) bool {
	for _, v := range Config.Devices.ExcludeByInstance {
		if strings.EqualFold(instance, v) {
			return true
		}
	}

	return false
}
