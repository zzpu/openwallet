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

package walletnode

import (
	"errors"
	"log"
	"strings"
)

const (
	RPCUser       = "walletUser"         //RPC默认的认证账户名
	RPCPassword   = "walletPassword2017" //RPC默认的认证账户密码
	RPCDockerPort = "9360/tcp"           //Docker中默认的RPC端口

	// DockerAllowed = "127.0.0.1" // ?500 For productive environment
	DockerAllowed = "0.0.0.0"

	MainNetDataPath = "/data" //容器中目录，实则在物理机："/openwallet/<Symbol>/data"
	TestNetDataPath = "/data" //容器中目录，实则在物理机："/openwallet/<Symbol>/testdata"

	MountSrcPrefix = "/openwallet/data" // The prefix to mounted source directory
	MountDstDir    = "/data"            // Which directory will be mounted in container
)

// Node setup 节点配置
type WalletnodeConfig struct {

	//钱包全节点服务
	walletnodePrefix          string // container prefix for name to run same within one server
	walletnodeServerType      string // "service"/"localdocker"/"remotedocker"
	walletnodeServerAddr      string // type:remotedocker required
	walletnodeServerPort      string // type:remotedocker required
	walletnodeStartNodeCMD    string // type:local required (from old: startNodeCMD)
	walletnodeStopNodeCMD     string // type:local required (from old: stopNodeCMD)
	walletnodeMainNetDataPath string
	walletnodeTestNetDataPath string
	walletnodeIsEncrypted     string // true/false
	// walletnodeServerSocket string "/var/run/docker.sock" // type:localdocker required
	// walletnodePubAPIs      string ""                     // walletnode returns API to rpc client, etc.

	isTestNet   string // 是否测试网络，default in TestNet
	RPCUser     string // RPC认证账户名
	RPCPassword string // RPC认证账户密码
	WalletURL   string // Fullnode API URL

	//------------------------------------------------------------------------------
	//默认配置内容
	defaultConfig string
}

func (wc WalletnodeConfig) getDataDir() (string, error) {
	if wc.isTestNet == "true" {
		if wc.walletnodeTestNetDataPath == "" {
			return "", errors.New("TestNetDataPath not config!")
		} else {
			return wc.walletnodeTestNetDataPath, nil
		}
	} else {
		if wc.walletnodeMainNetDataPath == "" {
			return "", errors.New("")
		} else {
			return wc.walletnodeMainNetDataPath, nil
		}
	}
}

func (wc WalletnodeConfig) isTestNetCheck() bool {
	if wc.isTestNet == "true" {
		return true
	}
	return false
}

type FullnodeContainerConfig struct {
	NAME string
	//CMD      [2][]string // Commands to run fullnode wallet ex: {{"/bin/sh", "mainnet"}, {"/bin/sh", "testnet"}}
	PORT      [][3]string // Which ports need to be mapped, ex: {{innerPort, mainNetPort, testNetPort}, ...}
	APIPORT   []string    // Port of default fullnode API(within container), from PORT
	IMAGE     string      // Image that container run from
	ENCRYPT   []string    // Encrypt wallet fullnode as an option
	STOPCMD   []string    // Use CMD to stop service
	NOTESTNET bool        // Is to support Testnet?
	LOGFIELS  [2]string   // [2]string{mainnet, testnet}
}

func (nc *FullnodeContainerConfig) isEncrypted() bool {
	if nc.ENCRYPT != nil {
		return true
	} else {
		return false
	}
}

func (nc *FullnodeContainerConfig) getLogFile() string {
	// is testnet?

	// have configs of LOG?
	if nc.LOGFIELS == [2]string{} {
		return ""
	} else if nc.LOGFIELS[0] == "" || nc.LOGFIELS[1] == "" {
		log.Println("FullnodeContainerConfig:LOGFILES invalid!")
		return ""
	}

	if WNConfig.isTestNetCheck() {
		return nc.LOGFIELS[1]
	} else {
		return nc.LOGFIELS[0]
	}
}

func getFullnodeConfig(symbol string) *FullnodeContainerConfig {
	if v, exist := FullnodeContainerConfigs[strings.ToLower(symbol)]; exist {
		return v
	} else {
		return nil
	}
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	WNConfig = &WalletnodeConfig{

		//钱包全节点服务
		walletnodePrefix:          "",
		walletnodeServerType:      "",
		walletnodeServerAddr:      "",
		walletnodeServerPort:      "",
		walletnodeStartNodeCMD:    "",
		walletnodeStopNodeCMD:     "",
		walletnodeMainNetDataPath: MainNetDataPath,
		walletnodeTestNetDataPath: TestNetDataPath,
		// walletnodeServerSocket string "/var/run/docker.sock" // type:localdocker required
		// walletnodePubAPIs      string ""                     // walletnode returns API to rpc client, etc.

		isTestNet:   "",
		RPCUser:     "walletUser",
		RPCPassword: "walletPassword2017",
		WalletURL:   "",

		//------------------------------------------------------------------------------
		//默认配置内容
		defaultConfig: `
# Is network test?
isTestNet = true
# node api url
WalletURL = ""
# RPC Authentication Username
rpcUser = "walletUser"
# RPC Authentication Password
rpcPassword = "walletPassword2017"
		
[walletnode]
# walletnode server type: local/docker
servertype = "docker"
# remote docker master server addr
serveraddr = "192.168.2.194"
# remote docker master server port
serverport = "2375"
# local docker master server socket if serveraddr=="" and servertype=="docker"
serversocket = "/var/run/docker.socket"

# prefix for container name
prefix = "openw_",
# wallet fullnode is crypted?
isEnCrypted = ""

# start node command if servertype==local
startNodeCMD = "",
# stop node command if servertype==local
stopNodeCMD = "",

# mainnet data path
mainNetDataPath = "/data"
# testnet data path
testNetDataPath = "/data"
`,
	}

	FullnodeContainerConfigs = map[string]*FullnodeContainerConfig{
		"btc": &FullnodeContainerConfig{ // Release0929
			PORT:     [][3]string{{"9360/tcp", "10001", "20001"}},
			APIPORT:  []string{"9360/tcp"},
			IMAGE:    string("openw/btc:v0.15.1"),
			ENCRYPT:  []string{"bitcoin-cli", "-datadir=/data", "-conf=/etc/bitcoin.conf", "encryptwallet 1234qwer"},
			STOPCMD:  []string{"bitcoin-cli", "-datadir=/data", "-conf=/etc/bitcoin.conf", "stop"},
			LOGFIELS: [2]string{"debug.log", "testnet3/debug.log"},
		},
		"eth": &FullnodeContainerConfig{ // Release0929
			PORT:     [][3]string{{"8545/tcp", "18545", "28545"}},
			APIPORT:  []string{"8545/tcp"},
			IMAGE:    string("openw/eth:geth-v1.8.15"),
			LOGFIELS: [2]string{"run.log", "run.log"},
		},
		"eos": &FullnodeContainerConfig{
			PORT:      [][3]string{{"8888/tcp", "18888", "28888"}},
			APIPORT:   []string{"8888/tcp"},
			IMAGE:     string("openw/eos:v1.2.5"),
			NOTESTNET: true,
		},
		// "iota": &FullnodeContainerConfig{
		// 	PORT:    [][3]string{{"18265/tcp", "10004", "20004"}},
		// 	APIPORT: []string{"18265/tcp"},
		// 	IMAGE:   string("openwallet/iota:latest"),
		// },

		// "bch": &FullnodeContainerConfig{
		// 	PORT:    [][3]string{{"18335/tcp", "10011", "20011"}},
		// 	APIPORT: []string{"18335/tcp"},
		// 	IMAGE:   string("openwallet/bch:0.17.1"),
		// },
		"bopo": &FullnodeContainerConfig{ // Release0929
			PORT:      [][3]string{{"9360/tcp", "10021", "20021"}},
			APIPORT:   []string{"9360/tcp"},
			IMAGE:     string("openw/bopo:latest"),
			NOTESTNET: true,
			LOGFIELS:  [2]string{"run.log", "run.log"},
		},
		"sc": &FullnodeContainerConfig{ // Siadcoin
			PORT:    [][3]string{{"9980/tcp", "19980", "29980"}, {"9981/tcp", "19981", "29981"}},
			APIPORT: []string{"9980/tcp"},
			IMAGE:   string("openw/siacoin:v1.3.4"),
		},
		// "hc": &FullnodeContainerConfig{
		// 	CMD: [2][]string{
		// 		{"/bin/bash", "-c", "/usr/local/hypercash/bin/hcd --datadir=/openwallet/data --logdir=/openwallet/data --rpcuser=wallet --rpcpass=walletPassword2017 --txindex --rpclisten=0.0.0.0:14009 && /usr/local/hypercash/bin/hcwallet --rpcconnect=127.0.0.1:14009 --username=wallet --password=walletPassword2017 --rpclisten=0.0.0.0:12010"},
		// 		{"/bin/bash", "-c", "/usr/local/hypercash/bin/hcd --datadir=/openwallet/testdata --logdir=/openwallet/testdata --rpcuser=wallet --rpcpass=walletPassword2017 --txindex --rpclisten=0.0.0.0:14009 --testnet && /usr/local/hypercash/bin/hcwallet --testnet --rpcconnect=127.0.0.1:14009 --username=wallet --password=walletPassword2017 --rpclisten=0.0.0.0:12010"}},
		// 	PORT:    [][3]string{{"12010/tcp", "10051", "20051"}, {"14009/tcp", "14009", "24009"}},
		// 	APIPORT: []string{"12010/tcp"},
		// 	IMAGE:   string("openwallet/hc:2.0.3dev"),
		// },
		"ltc": &FullnodeContainerConfig{ // Release0929
			NAME:     "Litecoin",
			PORT:     [][3]string{{"9360/tcp", "10061", "20061"}},
			APIPORT:  []string{"9360/tcp"},
			IMAGE:    string("openw/litecoin:v0.16.0"),
			ENCRYPT:  []string{"litecoin-cli", "-datadir=/data", "-conf=/etc/litecoin.conf", "encryptwallet 1234qwer"},
			STOPCMD:  []string{"litecoin-cli", "-datadir=/data", "-conf=/etc/litecoin.conf", "stop"},
			LOGFIELS: [2]string{"debug.log", "testnet4/debug.log"},
		},
		"qtum": &FullnodeContainerConfig{ // Release0929
			PORT:     [][3]string{{"8332/tcp", "18332", "28332"}},
			APIPORT:  []string{"8332/tcp"},
			IMAGE:    string("openw/qtum:v0.16.1"),
			ENCRYPT:  []string{"qtum-cli", "-datadir=/data", "-conf=/etc/qtum.conf", "encryptwallet 1234qwer"},
			STOPCMD:  []string{"qtum-cli", "-datadir=/data", "-conf=/etc/qtum.conf", "stop"},
			LOGFIELS: [2]string{"debug.log", "testnet3/debug.log"},
		},
		"trx": &FullnodeContainerConfig{ // Release0929
			NAME:     "Tron Network",
			PORT:     [][3]string{{"8090/tcp", "18090", "28090"}},
			APIPORT:  []string{"8090/tcp"},
			IMAGE:    string("openw/tron:v3.1.2"),
			STOPCMD:  []string{"bash", "-c", "kill -15 $(ps -ef | grep java-tron.jar | grep -v grep | awk '{print $2}') > /data/xxxx.txt"},
			LOGFIELS: [2]string{"logs/tron.log", "logs/tron.log"},
		},
		"nas": &FullnodeContainerConfig{ // Release0929
			NAME:     "NebulasIO",
			PORT:     [][3]string{{"8685/tcp", "18685", "28685"}},
			APIPORT:  []string{"8685/tcp"},
			IMAGE:    string("openw/nebulasio:v1.0.8"),
			LOGFIELS: [2]string{"logs/neb.log", "logs/neb.log"},
		},
		"ont": &FullnodeContainerConfig{ // Release0929
			NAME:     "Ontology",
			PORT:     [][3]string{{"20336/tcp", "20336", "30336"}},
			APIPORT:  []string{"20336/tcp"},
			IMAGE:    string("openw/ontology:v1.5.1"),
			LOGFIELS: [2]string{"run.log", "run.log"},
		},
	}
}
