/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package tezos

import (
	"testing"
	"github.com/blocktree/OpenWallet/openwallet"
)

var wm *WalletManager

func init() {
	wm = NewWalletManager()
	wm.InitConfigFlow()
	wm.Config.ServerAPI = "https://rpc.tezrpc.me"
	wm.WalletClient = NewClient(wm.Config.ServerAPI, false)
}

func TestApi(t *testing.T) {
	ret := wm.WalletClient.callGetHeader()
	t.Logf(string(ret))

	ret = wm.WalletClient.callGetbalance("tz1Neor2KRu3zp5FdMox98sxYLvFqtUs4fCJ")
	t.Logf(string(ret))
}

func TestCreateNewWallet(t *testing.T) {
	return
	w, keyfile, err := wm.CreateNewWallet("12", "jinxin")
	if err != nil {
		t.Error("create new wallet fail")
		return
	}

	t.Logf(w.WalletID)
	t.Logf(keyfile)

	ret, err := wm.GetWalletByID(w.WalletID)
	if err != nil {
		t.Error("get wallet by id err")
		t.Logf(err.Error())
		return
	}

	t.Logf(ret.Alias)
}

func TestLoadConfig(t *testing.T) {
	err := wm.LoadConfig()
	if err != nil {
		t.Error("load config error")
		t.Logf(err.Error())
	}
}

func TestWalletManager_GetWallets(t *testing.T) {
	err := wm.GetWalletList()
	if err != nil {
		t.Error("get wallet list error")
		t.Logf(err.Error())
	}
}

func TestWalletConfig_PrintConfig(t *testing.T) {
	err := wm.Config.PrintConfig()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestWalletManager_CreateBatchAddress(t *testing.T) {
	return
	var addrs []*openwallet.Address
	fpath, addrs, err := wm.CreateBatchAddress("WEY5DDuXbvHrBUa5UBKmVpwLCwP69bieeB", "jinxin", 10)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf(fpath)
	for _, a := range addrs {
		t.Logf(a.Address)
	}
}

func TestWalletManager_TransferFlow(t *testing.T) {
	w, err := wm.GetWalletByID("WEY5DDuXbvHrBUa5UBKmVpwLCwP69bieeB")
	if err != nil {
		t.Error("get wallet by id error")
		return
	}

	keystore, _ := w.HDKey("jinxin")

	db, err := w.OpenDB()
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer db.Close()

	var addrs []*openwallet.Address
	db.All(&addrs)

	var sender *openwallet.Address
	//key, _ := wm.getKeys(keystore, addrs[0])
	for _, a := range addrs {
		if a.Address == "tz1XSuiDu5Fevzwv3CG26dxmsinejRsUviC2" {
			sender = a
			t.Logf(a.Address)
			break
		}
	}

	key, err := wm.getKeys(keystore, sender)
	if err != nil {
		t.Error("get key error")
	}

	t.Logf("len:%d", len(key.PrivateKey))
	dst := "tz1Neor2KRu3zp5FdMox98sxYLvFqtUs4fCJ"

	inj, pre := wm.Transfer(*key, dst, "100", "100", "100", "1001")
	t.Logf("inj: %s, pre: %s", inj, pre)
}