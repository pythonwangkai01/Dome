package utils

import (
	"math/rand"
	"time"
)

//生成Guid字串
func UniqueId() int64 {

	return rand.New(rand.NewSource(time.Now().Unix())).Int63n(time.Now().Unix())
}
