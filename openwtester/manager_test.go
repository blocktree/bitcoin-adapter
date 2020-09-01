package openwtester

import (
	"fmt"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
	"github.com/blocktree/openwallet/v2/openwallet"
	"path/filepath"
	"testing"
)

var (
	testApp        = "bitcoin-adapter"
	configFilePath = filepath.Join("conf")
	dbFilePath     = filepath.Join("data", "db")
	dbFileName     = "blockchain.db"
)

func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"BTC",
	}
	return openw.NewWalletManager(tc)
	//tm.Init()
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO BTC", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "WNJGNWccZqF9TQBzBw3YmYZNWiZSqUT7s1")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()
	walletID := "WNJGNWccZqF9TQBzBw3YmYZNWiZSqUT7s1"
	//walletID := "W7tue6SDce38fPwerdKqyebUh6yo2nTQLC"
	account := &openwallet.AssetsAccount{Alias: "mainnetBTC", WalletID: walletID, Required: 1, Symbol: "BTC", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "W7tue6SDce38fPwerdKqyebUh6yo2nTQLC"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	//walletID := "W7tue6SDce38fPwerdKqyebUh6yo2nTQLC"
	//accountID := "FqQBQ8Bn26GogR7UAu6e2ZVhrYYmKUpmBS7CSM1KLTTZ"
	walletID := "WNJGNWccZqF9TQBzBw3YmYZNWiZSqUT7s1"
	//accountID := "BVF5bsaHBLmMY3tPMDdrNPD2Wizc3HhbAzbpom3KGzx"
	//accountID := "J1eD1qtXKW8xk5GFL9jNBWcCjeH2kBFiczQN1W6iAqi3"
	accountID := "AzQr4F4BWheGdm58PnPaYttX2hkF6CEVZ7mWTQbXW4iN"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 300)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WNJGNWccZqF9TQBzBw3YmYZNWiZSqUT7s1"
	//accountID := "BVF5bsaHBLmMY3tPMDdrNPD2Wizc3HhbAzbpom3KGzx"
	accountID := "J1eD1qtXKW8xk5GFL9jNBWcCjeH2kBFiczQN1W6iAqi3"
	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for _, w := range list {
		fmt.Printf("%s \n", w.Address)
	}
	//log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}
