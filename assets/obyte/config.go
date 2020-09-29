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

package obyte

import (
	"fmt"
	"github.com/blocktree/go-owcrypt"
	"github.com/shopspring/decimal"
	"github.com/zzpu/openwallet/common/file"
	"path/filepath"
	"strings"
	"time"
)

const (
	//币种
	Symbol    = "GBYTE"
	MasterKey = "masterkey seed"
	CurveType = owcrypt.ECC_CURVE_SECP256K1
)

type WalletConfig struct {

	//币种
	Symbol    string
	MasterKey string

	addressDir string
	//配置文件路径
	configFilePath string
	//配置文件名
	configFileName string
	//数据路径
	//dbPath string
	//备份路径
	backupDir string
	//钱包服务API
	ServerAPI string
	//钱包数据文件目录
	WalletDataPath string
	//汇总阀值
	Threshold decimal.Decimal
	//汇总地址
	SumAddress string
	//汇总执行间隔时间
	CycleSeconds time.Duration
	//默认配置内容
	DefaultConfig string
	//曲线类型
	CurveType uint32
	//小数位长度
	CoinDecimals int32
	//最小矿工费
	MinFees string
}

func NewConfig(symbol string) *WalletConfig {

	c := WalletConfig{}

	//币种
	c.Symbol = symbol
	//c.MasterKey = masterKey
	c.CurveType = CurveType
	//地址导出路径
	c.addressDir = filepath.Join("data", strings.ToLower(c.Symbol), "address")
	//配置文件路径
	c.configFilePath = filepath.Join("conf")
	//配置文件名
	c.configFileName = c.Symbol + ".ini"
	//本地数据库文件路径
	//c.dbPath = filepath.Join("data", strings.ToLower(c.Symbol), "db")
	//备份路径
	c.backupDir = filepath.Join("data", strings.ToLower(c.Symbol), "backup")
	//钱包服务API
	c.ServerAPI = "http://127.0.0.1:10000"
	//钱包数据文件目录
	c.WalletDataPath = ""
	//汇总阀值
	c.Threshold = decimal.NewFromFloat(5)
	//汇总地址
	c.SumAddress = ""
	//汇总执行间隔时间
	c.CycleSeconds = time.Second * 10
	//小数位长度
	c.CoinDecimals = 6
	//核心钱包密码，配置有值用于自动解锁钱包
	//c.WalletPassword = ""

	//默认配置内容
	c.DefaultConfig = `

# wallet data store path
walletDataPath = ""
# RPC api url
serverAPI = ""
# the safe address that wallet send money to.
sumAddress = ""
# when wallet's balance is over this value, the wallet will send money to [sumAddress]
threshold = ""
# summary task timer cycle time, sample: 1m , 30s, 3m20s etc
cycleSeconds = ""
# Coin Decimals
coinDecimals = 
# Minimum fee for summary wallet 
minFees = ""
`

	//创建目录
	//file.MkdirAll(c.dbPath)
	file.MkdirAll(c.backupDir)
	//file.MkdirAll(c.keyDir)

	return &c
}

//printConfig Print config information
func (wc *WalletConfig) PrintConfig() error {

	wc.InitConfig()
	//读取配置
	absFile := filepath.Join(wc.configFilePath, wc.configFileName)
	fmt.Printf("-----------------------------------------------------------\n")
	file.PrintFile(absFile)
	fmt.Printf("-----------------------------------------------------------\n")

	return nil

}

//initConfig 初始化配置文件
func (wc *WalletConfig) InitConfig() {

	//读取配置
	absFile := filepath.Join(wc.configFilePath, wc.configFileName)
	if !file.Exists(absFile) {
		file.MkdirAll(wc.configFilePath)
		file.WriteFile(absFile, []byte(wc.DefaultConfig), false)
	}

}
