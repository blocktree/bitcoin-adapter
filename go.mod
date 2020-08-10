module github.com/blocktree/bitcoin-adapter

go 1.12

require (
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet/v2 v2.0.6
	github.com/bndr/gotabulate v1.1.2
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v0.0.0-20191219182022-e17c9730c422
	github.com/codeskyblue/go-sh v0.0.0-20190412065543-76bd3d59ff27
	github.com/ethereum/go-ethereum v1.9.9
	github.com/graarh/golang-socketio v0.0.0-20170510162725-2c44953b9b5f
	github.com/imroc/req v0.2.4
	github.com/pborman/uuid v1.2.0
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/tidwall/gjson v1.3.5
)

//replace github.com/blocktree/openwallet => ../../openwallet
//replace github.com/blocktree/go-owcdrivers => ../../go-owcdrivers
