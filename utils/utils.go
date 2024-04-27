package utils

import (
    "log"
)

func Error(err error, messages ...string) {
    if err != nil {
        log.Panic(messages)
    }
}
