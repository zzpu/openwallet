module github.com/zzpu/openwallet

go 1.13

require (
	docker.io/go-docker v1.0.0
	github.com/NebulousLabs/entropy-mnemonics v0.0.0-20181203154559-bc7e13c5ccd8
	github.com/asdine/storm v2.1.2+incompatible
	github.com/astaxie/beego v1.12.1
	github.com/blocktree/ethereum-adapter v1.6.2
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.2
	github.com/bndr/gotabulate v1.1.2
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/btcsuite/btcutil v0.0.0-20191219182022-e17c9730c422
	github.com/bwmarrin/snowflake v0.3.0
	github.com/codeskyblue/go-sh v0.0.0-20190412065543-76bd3d59ff27
	github.com/couchbase/go-couchbase v0.0.0-20200519150804-63f3cdb75e0d
	github.com/docker/go-connections v0.4.0
	github.com/ethereum/go-ethereum v1.9.9
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gorilla/websocket v1.4.1
	github.com/imroc/req v0.3.0
	github.com/lib/pq v1.3.0
	github.com/mr-tron/base58 v1.1.3
	github.com/pborman/uuid v1.2.0
	github.com/peterh/liner v1.1.1-0.20190123174540-a2c9a5303de7
	github.com/pkg/errors v0.8.1
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/siddontang/ledisdb v0.0.0-20190202134119-8ceb77e66a92
	github.com/ssdb/gossdb v0.0.0-20180723034631-88f6b59b84ec
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/tidwall/gjson v1.5.0
	github.com/tyler-smith/go-bip39 v1.0.2
	go.etcd.io/bbolt v1.3.3
	golang.org/x/crypto v0.0.0-20191227163750-53104e6ec876
	gopkg.in/urfave/cli.v1 v1.20.0
)

//replace github.com/blocktree/go-owcdrivers v1.12.0 => /opt/code/abc/blocktree/go-owcdrivers/

//replace github.com/zzpu/openwallet => /opt/code/abc/blocktree/openwallet
replace github.com/astaxie/beego v1.11.1 => github.com/astaxie/beego v1.12.2

replace github.com/blocktree/ethereum-adapter v1.1.10 => /opt/code/abc/blocktree/openwallet
