package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RandIntn(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func NewUUID() string {
	vUUID, err := uuid.NewRandom()
	if err == nil {
		return strings.Replace(vUUID.String(), "-", "", -1)
	}

	vTime := fmt.Sprintf("%v", time.Now().UnixNano())
	var vSuffix []byte
	for i := len(vTime); i < 32; i++ {
		vSuffix = append(vSuffix, "abcdefghijklmnopqrstuvwxyz"[RandIntn(26)])
	}
	vFakeUUID := vTime + string(vSuffix)
	return vFakeUUID
}
