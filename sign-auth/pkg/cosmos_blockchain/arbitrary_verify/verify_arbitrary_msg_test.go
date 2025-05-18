package arbitraryverify

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/assert"
)

func Test_VerifyArbitraryMsg(t *testing.T) {
	bytePubKey, err := hex.DecodeString("038318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed75")
	if err != nil {
		log.Fatal("failed to get byte pubkey from hex: ", err)
	}

	hexSign := "18686c968bf55832a0a97e08e05d81e859ea2b85208dfe57a7bc09d075cc293b16ec21f5593ff97cdb14ee204b325496c69f5282dc16af3b267dc0ec1371431e"
	byteSignature, err := hex.DecodeString(hexSign)
	if err != nil {
		log.Fatal("failed to get byte sign from hex: ", err)
	}
	pubKey := secp256k1.PubKey{Key: bytePubKey}
	stat, err := VerifyArbitraryMsg("sah17w0adeg64ky0daxwd2ugyuneellmjgnx5awphz", "asdf", byteSignature, pubKey)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, stat, "verify result should be true")
}
