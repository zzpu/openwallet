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

package ontology

import (
	"fmt"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcdrivers/ontologyTransaction"
	owcrypt "github.com/blocktree/go-owcrypt"
)

type addressDecoder struct {
	wm *WalletManager //钱包管理者
}

//NewAddressDecoder 地址解析器
func NewAddressDecoder(wm *WalletManager) *addressDecoder {
	decoder := addressDecoder{}
	decoder.wm = wm
	return &decoder
}

//PublicKeyToAddress 公钥转地址
func (decoder *addressDecoder) PublicKeyToAddress(pub []byte, isTestnet bool) (string, error) {
	cfg := addressEncoder.ONT_Address

	pub = append([]byte{byte(len(pub))}, pub...)
	pub = append(pub, ontologyTransaction.OpCodeCheckSig)

	pkHash := owcrypt.Hash(pub, 0, owcrypt.HASH_ALG_HASH160)

	address := addressEncoder.AddressEncode(pkHash, cfg)

	return address, nil
}

//ScriptPubKeyToBech32Address scriptPubKey转Bech32地址
func ScriptPubKeyToBech32Address(scriptPubKey []byte, isTestnet bool) (string, error) {
	var (
		hash []byte
	)

	cfg := addressEncoder.BTC_mainnetAddressBech32V0
	if isTestnet {
		cfg = addressEncoder.BTC_testnetAddressBech32V0
	}

	if len(scriptPubKey) == 22 || len(scriptPubKey) == 34 {

		hash = scriptPubKey[2:]

		address := addressEncoder.AddressEncode(hash, cfg)

		return address, nil

	} else {
		return "", fmt.Errorf("scriptPubKey length is invalid")
	}

}
