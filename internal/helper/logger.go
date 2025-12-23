package helper

import (
	"runtime"

	"github.com/rs/zerolog/log"
)

func LogError(err error) {
	if err == nil {
		return
	}

	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc).Name()

	log.Error().
		Err(err).
		Str("func", fn).
		Int("line", line).
		Send()
}

func LogInfo(message string) {
	log.Info().Msg(message)
}

func LogWarn(message string) {
	log.Warn().Msg(message)
}
