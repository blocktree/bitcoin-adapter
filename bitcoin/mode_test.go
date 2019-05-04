package bitcoin

import (
	"encoding/hex"
	"github.com/blocktree/openwallet/log"
	"testing"
)

func TestDecodeScript_P2PK(t *testing.T) {
	scriptPubKeyEqual := "76a914d20b7f40aa0b18ccd008ce75b9abc980db1f37d788ac"
	script := "OP_DUP OP_HASH160 d20b7f40aa0b18ccd008ce75b9abc980db1f37d7 OP_EQUALVERIFY OP_CHECKSIG"
	scriptPubKey, err := DecodeScript(script)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
		return
	}
	log.Infof("scriptPubKey: %s", hex.EncodeToString(scriptPubKey))
	if scriptPubKeyEqual != hex.EncodeToString(scriptPubKey) {
		t.Errorf("scriptPubKey is not equal: %s", scriptPubKeyEqual)
		return
	}
}

func TestDecodeScript_P2PSW(t *testing.T) {
	scriptPubKeyEqual := "002079db247b3da5d5e33e036005911b9341a8d136768a001e9f7b86c5211315e3e1"
	script := "0 79db247b3da5d5e33e036005911b9341a8d136768a001e9f7b86c5211315e3e1"
	scriptPubKey, err := DecodeScript(script)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
		return
	}
	log.Infof("scriptPubKey: %s", hex.EncodeToString(scriptPubKey))
	if scriptPubKeyEqual != hex.EncodeToString(scriptPubKey) {
		t.Errorf("scriptPubKey is not equal: %s", scriptPubKeyEqual)
		return
	}
}
