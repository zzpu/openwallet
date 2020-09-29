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

package openw

import "path/filepath"

var (
	defaultDataDir = filepath.Join(".", "openw_data")
)

type Config struct {
	KeyDir          string   //钥匙备份路径
	DBPath          string   //本地数据库文件路径
	BackupDir       string   //备份路径
	SupportAssets   []string //支持的资产类型
	EnableBlockScan bool
	ConfigDir       string
}

func NewConfig() *Config {

	c := Config{}

	//钥匙备份路径
	c.KeyDir = filepath.Join(defaultDataDir, "key")
	//本地数据库文件路径
	c.DBPath = filepath.Join(defaultDataDir, "db")
	//备份路径
	c.BackupDir = filepath.Join(defaultDataDir, "backup")
	//支持资产
	c.SupportAssets = []string{"BTC", "ETH", "QTUM", "NAS", "TRX"}
	//开启区块扫描
	c.EnableBlockScan = true

	return &c
}

//loadConfig 加载配置文件
//@param path 配置文件路径
func loadConfig(path string) *Config {
	return nil
}
