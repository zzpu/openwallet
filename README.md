# OpenWallet

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg?style=flat-square)](https://www.gnu.org/licenses/gpl-3.0)
[![Slack](https://img.shields.io/badge/Slack-blocktree-green.svg?style=flat-square)](https://join.slack.com/t/blocktreeworkspace/shared_invite/enQtODE0OTY1ODA5MjM3LTIwNDJjNDA2ZDYyMjE1YTRiMDUyYTg5NDljMTU0NTY4NmE5ZDM2ZGI3NjcxYzMzN2NiZTZjNzIyMjVjYmRiY2E)

## Contents

- [Introduction](#Introduction)
- [Features](#Features)
- [Protocol](./openwallet/docs/README.md)
- [Resources](#Resources)
- [Contributing](#Contributing)
- [Sponsors](#Sponsors)
- [LICENSE](LICENSE)

## Introduction

openwallet框架定义了一套规范化的钱包体系开发模型，应用开发者无需了解区块链底层协议，即可支持多种区块链资产管理功能。
区块链底层协议开发者也可以实现openwallet框架资产适配器协议，开发者便可利用该区块链生态实现更多落地应用。
openwallet框架，未来将进化成为一个去中心化系统开发框架，跟随不断发展的区块链生态，提供更多应用开发组件库。

## Features

### v1

细节上，openwallet框架有以下特性:

- **规范化的钱包体系模型**。我们把钱包体系划分为3层模型：钱包，资产账户，地址。**钱包**：理论上就是一个32字节的种子，可通过用户口令加密保存keystore文件。种子相当重要，它基于BIP32协议衍生子代密钥，丢失将失去所有子代密钥。**资产账户**：一个钱包可以创建不同区块链协议的资产账户。这是因为openwallet的扩展密钥算法已经实现了多种ECC曲线的支持，例如：secp256k1，ed25519等。资产账户主要提供一个非强化版的扩展公钥，用于生成更多的地址。**地址**：真正用于接收区块链资产的是地址模型。业务应用可以根据需求去组织自己的钱包。例如：普通用户创建一个onchain钱包，他分别创建了多种区块链的资产账户，账户默认都只创建一个地址，解决各种支付方式。
- **区块链资产适配器协议**。区块链协议各有不同，如果应用开发者把每种区块链协议都弄明白并实现落地应用，将会是一个漫长艰难复杂的开发过程。经过多时间的研究，openwallet钱包体系模型适用于大量主流的区块链协议。为此，我们定义一个规范化的区块链资产适配器协议，让区块链协议适配者实现：主链基本信息，配置加载过程，地址解析器，交易单解析器，区块链扫描器，智能合约解析器，日志管理工具等等方法。区块链资产适配器实现后，可通过openw包提供的功能进行测试。
- **OWTP通信协议**。openwallet定义自己的加密通信协议，内置SM2协商密码机制，可以在非https环境下开启加密通信。

### v2

兼容v1功能，扩展智能合约解析器，支持与合约进行交互，监听合约的事件日志。

## openwallet V2升级

v2版本完全兼容v1版本，只需在go.mod中依赖v2最新版本

```go script

require (
    github.com/zzpu/openwallet latest
    ....
)

```


已依赖v1库的项目，可以安装[Mod](https://github.com/marwan-at-work/mod/)

```shell script

GO111MODULE=on go get github.com/marwan-at-work/mod/cmd/mod

# 在你的项目的根目录，执行命令，import openwallet库相关的文件都会更新到v2包
mod upgrade --mod-name=github.com/blocktree/openwallet

```

## Resources

### Assets-Adapter 区块链资产适配器

#### v1

- [bitcoin-adapter](https://github.com/blocktree/bitcoin-adapter)
- [litecoin-adapter](https://github.com/blocktree/litecoin-adapter)（继承bitcoin-adapter）
- [ethereum-adapter](https://github.com/blocktree/ethereum-adapter)
- [truechain-adapter](https://github.com/blocktree/truechain-adapter)（继承ethereum-adapter）
- [qtum-adapter](https://github.com/blocktree/qtum-adapter)
- [tron-adapter](https://github.com/blocktree/tron-adapter)
- [more...](https://github.com/blocktree?utf8=%E2%9C%93&q=adapter)

#### v2

- [quorum-adapter](https://github.com/blocktree/quorum-adapter)

### 钱包开发者平台

- [openwallet钱包服务开放平台](https://www.openwallet.cn)
- [go-openw-sdk](https://github.com/blocktree/go-openw-sdk)
- [go-openw-cli](https://github.com/blocktree/go-openw-cli)

## Contributing

如果你对项目有非常感兴趣，可创建[issue](https://github.com/blocktree/openwallet/issues/new) 分享你的想法和创意。

## Sponsors

## License

openwallet is licensed under the GNU General Public License v3.0. See [LICENSE](LICENSE) for the full license text.
