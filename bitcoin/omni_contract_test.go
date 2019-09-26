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

package bitcoin

import (
	"testing"
)

func TestWalletManager_GetOmniBalance(t *testing.T) {
	propertyID := uint64(2)
	address := "mkJrhf8Bp3RWfL5eyatcfqPBDDUUXgHQYm"
	balance, err := tw.GetOmniBalance(propertyID, address)
	if err != nil {
		t.Errorf("GetOmniBalance failed unexpected error: %v\n", err)
		return
	}
	t.Logf("balance: %v\n", balance)
}

func TestWalletManager_IsHaveOmniAssets(t *testing.T) {
	address := "mwawxdBn9w4CPxic961vPnyj9HqDVGnkth"
	//address := "mi9qsHKMqtrgnbxg7ifdPMk1LsFmen4xNn"
	bool := tw.IsHaveOmniAssets(address)
	t.Logf("IsHaveOmniAssets: %v\n", bool)
}

func TestWalletManager_GetOmniTransaction(t *testing.T) {
	//txid := "9bceadcd1f043b5888eaff6ec3656717a8baeaf67d04a3c78db2aedaf8cb477e"
	txid := "924da0633c28ad4d61ff6f68014e8af50020a71338aa067cec6f3469b74a1121"
	transaction, err := tw.GetOmniTransaction(txid)
	if err != nil {
		t.Errorf("GetOmniBalance failed unexpected error: %v\n", err)
		return
	}
	t.Logf("transaction: %+v", transaction)
}

func TestWalletManager_GetOmniInfo(t *testing.T) {
	result, err := tw.GetOmniInfo()
	if err != nil {
		t.Errorf("TestWalletManager_GetOmniInfo failed unexpected error: %v\n", err)
		return
	}
	t.Logf("OmniInfo: %+v", result)
}

func TestWalletManager_GetOmniProperty(t *testing.T) {
	propertyID := uint64(2)
	result, err := tw.GetOmniProperty(propertyID)
	if err != nil {
		t.Errorf("GetOmniProperty failed unexpected error: %v\n", err)
		return
	}
	t.Logf("GetOmniProperty: %+v", result)
}


func TestWalletManager_GetOmniBestBlockHash(t *testing.T) {
	blockhash, err := tw.GetOmniBestBlockHash()
	if err != nil {
		t.Errorf("GetOmniBestBlockHash failed unexpected error: %v\n", err)
		return
	}
	t.Logf("blockhash: %+v", blockhash)
}

func TestWalletManager_GetOmniBlockHeight(t *testing.T) {
	blockheight, err := tw.GetOmniBlockHeight()
	if err != nil {
		t.Errorf("GetOmniBlockHeight failed unexpected error: %v\n", err)
		return
	}
	t.Logf("blockheight: %+v", blockheight)
}

func TestWalletManager_GetOmniBlockHash(t *testing.T) {
	blockheight, err := tw.GetOmniBlockHash(596442)
	if err != nil {
		t.Errorf("GetOmniBlockHeight failed unexpected error: %v\n", err)
		return
	}
	t.Logf("blockheight: %+v", blockheight)
}