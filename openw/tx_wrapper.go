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

package openw

import (
	"fmt"
	"github.com/asdine/storm/q"
	"github.com/shopspring/decimal"
	"github.com/zzpu/openwallet/common"
	"github.com/zzpu/openwallet/openwallet"
)

// TransactionWrapper 交易包装器，扩展钱包交易单相关功能
type TransactionWrapper struct {
	*WalletWrapper
}

func NewTransactionWrapper(args ...interface{}) *TransactionWrapper {

	wrapper := NewWalletWrapper(args...)

	txWrapper := TransactionWrapper{WalletWrapper: wrapper}

	for _, arg := range args {
		switch obj := arg.(type) {
		case *WalletWrapper:
			txWrapper.WalletWrapper = obj
		}
	}

	return &txWrapper
}

//GetTxInputs 获取钱包的出账记录
func (wrapper *WalletWrapper) GetTxInputs(offset, limit int, cols ...interface{}) ([]*openwallet.TxInput, error) {

	//打开数据库
	db, err := wrapper.OpenStormDB()
	if err != nil {
		return nil, err
	}
	defer wrapper.CloseDB()

	var txs []*openwallet.TxInput

	query := make([]q.Matcher, 0)

	if len(cols)%2 != 0 {
		return nil, fmt.Errorf("condition param is not pair")
	}

	for i := 0; i < len(cols); i = i + 2 {
		field := common.NewString(cols[i])
		val := cols[i+1]
		query = append(query, q.Eq(field.String(), val))
	}

	if limit > 0 {

		err = db.Select(q.And(
			query...,
		)).Limit(limit).Skip(offset).Find(&txs)

	} else {

		err = db.Select(q.And(
			query...,
		)).Skip(offset).Find(&txs)

	}

	if err != nil {
		return nil, fmt.Errorf("can not find txInputs")
	}

	return txs, nil
}

//GetTxOutputs 获取钱包的入账记录
func (wrapper *WalletWrapper) GetTxOutputs(offset, limit int, cols ...interface{}) ([]*openwallet.TxOutPut, error) {

	//打开数据库
	db, err := wrapper.OpenStormDB()
	if err != nil {
		return nil, err
	}
	defer wrapper.CloseDB()

	var txs []*openwallet.TxOutPut

	query := make([]q.Matcher, 0)

	if len(cols)%2 != 0 {
		return nil, fmt.Errorf("condition param is not pair")
	}

	for i := 0; i < len(cols); i = i + 2 {
		field := common.NewString(cols[i])
		val := cols[i+1]
		query = append(query, q.Eq(field.String(), val))
	}

	if limit > 0 {

		err = db.Select(q.And(
			query...,
		)).Limit(limit).Skip(offset).Find(&txs)

	} else {

		err = db.Select(q.And(
			query...,
		)).Skip(offset).Find(&txs)

	}

	if err != nil {
		return nil, fmt.Errorf("can not find txoutputs")
	}

	return txs, nil
}

//GetTransactions 获取钱包的交易记录
func (wrapper *WalletWrapper) GetTransactions(offset, limit int, cols ...interface{}) ([]*openwallet.Transaction, error) {

	//打开数据库
	db, err := wrapper.OpenStormDB()
	if err != nil {
		return nil, err
	}
	defer wrapper.CloseDB()

	var txs []*openwallet.Transaction

	query := make([]q.Matcher, 0)

	if len(cols)%2 != 0 {
		return nil, fmt.Errorf("condition param is not pair")
	}

	for i := 0; i < len(cols); i = i + 2 {
		field := common.NewString(cols[i])
		val := cols[i+1]
		query = append(query, q.Eq(field.String(), val))
	}

	if limit > 0 {

		err = db.Select(q.And(
			query...,
		)).Limit(limit).Skip(offset).Find(&txs)

	} else {

		err = db.Select(q.And(
			query...,
		)).Skip(offset).Find(&txs)

	}

	if err != nil {
		return nil, fmt.Errorf("can not find transactions")
	}

	return txs, nil
}

//SaveBlockExtractData 保存区块提取数据
func (wrapper *TransactionWrapper) SaveBlockExtractData(accountID string, data *openwallet.TxExtractData) error {

	var (
		accountSpent    = decimal.Zero
		accountReceived = decimal.Zero
	)

	//打开数据库
	db, err := wrapper.OpenStormDB()
	if err != nil {
		return err
	}
	defer wrapper.CloseDB()

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	//保存出账的记录
	for _, input := range data.TxInputs {
		a, err := wrapper.GetAddress(input.Address)
		if err != nil {
			continue
		}
		input.AccountID = a.AccountID
		err = tx.Save(input)
		if err != nil {
			return fmt.Errorf("wallet save TxInputs failed, unexpected error: %v", err)
		}

		//统计该交易单下的各个资产账户的支出总数
		if a.AccountID == accountID {
			amount, _ := decimal.NewFromString(input.Amount)
			accountSpent = accountSpent.Add(amount)
		}
	}

	//保存入账的记录
	for _, output := range data.TxOutputs {
		a, err := wrapper.GetAddress(output.Address)
		if err != nil {
			continue
		}
		output.AccountID = a.AccountID
		err = tx.Save(output)
		if err != nil {
			return fmt.Errorf("wallet save TxOutputs failed, unexpected error: %v", err)
		}

		//统计该交易单下的各个资产账户的收入总数
		if a.AccountID == accountID {
			amount, _ := decimal.NewFromString(output.Amount)
			accountReceived = accountReceived.Add(amount)
		}
	}

	//计算该交易单下的各个资产账户实际总收支，记录为账单数据
	trx := data.Transaction
	trx.AccountID = accountID
	trx.Amount = accountReceived.Sub(accountSpent).StringFixed(trx.Decimal)

	//保存账户相关的记录
	err = tx.Save(trx)
	if err != nil {
		return fmt.Errorf("wallet save Transactions failed, unexpected error: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("wallet save TxExtractData failed, unexpected error: %v", err)
	}

	return nil
}

//DeleteBlockDataByHeight 删除钱包中指定区块高度相关的交易记录
func (wrapper *TransactionWrapper) DeleteBlockDataByHeight(height uint64) error {

	//打开数据库
	db, err := wrapper.OpenStormDB()
	if err != nil {
		return err
	}
	defer wrapper.CloseDB()

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	trxs, err := wrapper.GetTransactions(0, -1, "BlockHeight", height)
	if err != nil {
		return err
	}

	for _, obj := range trxs {
		err = db.DeleteStruct(&obj)
		if err != nil {
			return err
		}
	}

	inputs, err := wrapper.GetTxInputs(0, -1, "BlockHeight", height)
	if err != nil {
		return err
	}

	for _, obj := range inputs {
		err = db.DeleteStruct(&obj)
		if err != nil {
			return err
		}
	}

	outputs, err := wrapper.GetTxOutputs(0, -1, "BlockHeight", height)
	if err != nil {
		return err
	}

	for _, obj := range outputs {
		err = db.DeleteStruct(&obj)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
