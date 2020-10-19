/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/blocktree/openwallet/v2/openw"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract, nil)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract,
	feeSupportAccount *openwallet.FeesSupportAccount) ([]*openwallet.RawTransactionWithError, error) {

	rawTxArray, err := tm.CreateSummaryRawTransactionWithError(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract, feeSupportAccount)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer(t *testing.T) {

	targets := []string{
		//"18p2zM6CFMmzH98TbE2iYM5FRbquDfXdn2",
		//"1AAGzDgSmcvc7YLf492ZcKhLrdHd7RWGtQ",
		//"1ECc5jUhWgow42mtUXULHJv5f4WpS5JpQ4",
		//"1HCc8DHwvw5QNMWmDcCCp6PA5XQ6D3dCJJ",
		//"1LRbPFZAbwENBeBpXhEqWqHMCXwaHxk3Tk",
		//"1P1PLC4N5oGGnd77dPxyjZMKBR522eP4Ao",
		//"1Q4PafoRpGCHC36vxVFCuowGQHojwSa92r",

		"1TpK5B1dZpR2D4TZqzv7PsUZwnCE31QrY", //fee support address
	}

	tm := testInitWalletManager()
	walletID := "WNJGNWccZqF9TQBzBw3YmYZNWiZSqUT7s1"
	accountID := "BVF5bsaHBLmMY3tPMDdrNPD2Wizc3HhbAzbpom3KGzx"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	for _, to := range targets {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.005", "", nil)
		if err != nil {
			return
		}

		log.Std.Info("rawTx: %+v", rawTx)

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}

}

func TestTransfer_OMNI(t *testing.T) {

	targets := []string{
		//"18p2zM6CFMmzH98TbE2iYM5FRbquDfXdn2",
		//"1AAGzDgSmcvc7YLf492ZcKhLrdHd7RWGtQ",
		//"1ECc5jUhWgow42mtUXULHJv5f4WpS5JpQ4",
		//"1HCc8DHwvw5QNMWmDcCCp6PA5XQ6D3dCJJ",
		//"1LRbPFZAbwENBeBpXhEqWqHMCXwaHxk3Tk",
		//"1P1PLC4N5oGGnd77dPxyjZMKBR522eP4Ao",
		//"1Q4PafoRpGCHC36vxVFCuowGQHojwSa92r",

	}

	tm := testInitWalletManager()
	walletID := "WNJGNWccZqF9TQBzBw3YmYZNWiZSqUT7s1"
	accountID := "BVF5bsaHBLmMY3tPMDdrNPD2Wizc3HhbAzbpom3KGzx"

	contract := openwallet.SmartContract{
		Address:  "31",
		Symbol:   "BTC",
		Name:     "TetherUSD",
		Token:    "USDT",
		Decimals: 8,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	for _, to := range targets {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.005", "", &contract)
		if err != nil {
			return
		}

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}
}

func TestSummary(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WNJGNWccZqF9TQBzBw3YmYZNWiZSqUT7s1"
	accountID := "BVF5bsaHBLmMY3tPMDdrNPD2Wizc3HhbAzbpom3KGzx"
	summaryAddress := "1GKufQGv2F2dcfq4uJBP3nEeRBbrnnoZ3t"

	//walletID := "WAmTnvPKMWpJBqKk6cncFG3mTXz3iPmtzV"
	//accountID := "86uUBCjk4SqEtMGDt92SQfn7YLhCZEcNQGjD5GhNNtSa"
	//summaryAddress := "12kSR8J11Q1d8JiYwZn7DZsPoDoptME35y"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}

func TestSummary_OMNI(t *testing.T) {

	tm := testInitWalletManager()
	walletID := "WNJGNWccZqF9TQBzBw3YmYZNWiZSqUT7s1"
	accountID := "BVF5bsaHBLmMY3tPMDdrNPD2Wizc3HhbAzbpom3KGzx"
	summaryAddress := "1Nfj41RpckiBVY7xXVvgfkwE1f5b3iNLHH"

	//walletID := "WAmTnvPKMWpJBqKk6cncFG3mTXz3iPmtzV"
	//accountID := "86uUBCjk4SqEtMGDt92SQfn7YLhCZEcNQGjD5GhNNtSa"
	//summaryAddress := "12kSR8J11Q1d8JiYwZn7DZsPoDoptME35y"

	//contract := openwallet.SmartContract{
	//	Address:  "2",
	//	Symbol:   "BTC",
	//	Name:     "Test Omni",
	//	Token:    "Omni",
	//	Decimals: 8,
	//}

	contract := openwallet.SmartContract{
		Address:  "31",
		Symbol:   "BTC",
		Name:     "TetherUSD",
		Token:    "USDT",
		Decimals: 8,
	}

	feesSupport := openwallet.FeesSupportAccount{
		AccountID:        "2GVc6ee1RmHhcGD38C7uNZXoatZyrSywtZQ3qE7VjpMj",
		FeesSupportScale: "1",
		//fee support address
		//1TpK5B1dZpR2D4TZqzv7PsUZwnCE31QrY
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 20, &contract, &feesSupport)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTxWithErr := range rawTxArray {

		if rawTxWithErr.Error != nil {
			log.Error(rawTxWithErr.Error.Error())
			continue
		}

		_, err = testSignTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTxWithErr.RawTx)
		if err != nil {
			return
		}
	}

}
