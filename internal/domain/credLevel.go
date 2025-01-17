package domain

import "github.com/iput-kernel/foundation-account/internal/config"

func GetCredLevel(credit int64, config config.CredConfig) int32 {
	if credit < int64(config.Level5Thereshold) {
		return 5
	}
	if credit < int64(config.Level4Thereshold) {
		return 4
	}
	if credit < int64(config.Level3Thereshold) {
		return 3
	}
	if credit < int64(config.Level2Thereshold) {
		return 2
	}
	if credit < int64(config.Level1Thereshold) {
		return 1
	}
	return 0
}
