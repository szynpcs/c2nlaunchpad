package utils

import (
	"c2n/config"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanHexPrefix(t *testing.T) {
	hexString := "0x1234567890"
	result := CleanHexPrefix(hexString)
	if result != "1234567890" {
		t.Errorf("CleanHexPrefix failed, expected %s, got %s", "1234567890", result)
	}
}

// TODO: 签名和java结果不一致
func TestGetSign(t *testing.T) {
	config.AppConfig.Owner.PrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	hexString := "0x70997970C51812dc3A010C7d01b50e0d17dc79C88acd85898458400f7db866d53fcff6f0d49741ff"
	result := GetSign(hexString)
	fmt.Println(result)
	assert.Equal(t, result, "393231be757e3ce3a7c64a529454346c14241c25598fc60e5a68cbec0f4c60c8500e69fed973c1fc9cf5c5235ac49bced68d0b7ef3614e68284f73c69d6181011b")
}
