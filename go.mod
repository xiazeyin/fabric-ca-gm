module github.com/xiazeyin/fabric-ca-gm

go 1.16

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.5.2
	github.com/hyperledger/fabric-amcl v0.0.0-20210603140002-2670f91851c8
	github.com/jmoiron/sqlx v1.3.5
	github.com/kisielk/sqlstruct v0.0.0-20210630145711-dae28ed37023
	github.com/lib/pq v1.10.5
	github.com/mattn/go-sqlite3 v1.14.12
	github.com/mitchellh/mapstructure v1.4.3
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.18.1
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.6.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.1
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d
)

require (
	cloud.google.com/go v0.99.0 // indirect
	github.com/certifi/gocertifi v0.0.0-20210507211836-431795d63e8d // indirect
	github.com/cncf/xds/go v0.0.0-20211130200136-a8f946100490 // indirect
	github.com/envoyproxy/go-control-plane v0.10.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.2 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli v1.22.5 // indirect
	github.com/xiazeyin/cfssl-gm v0.0.0-20220524125831-51406acbda75
	github.com/xiazeyin/fabric-config-gm v0.0.0-20220524082816-d8436d264ce6
	github.com/xiazeyin/fabric-gm v0.0.0-20220524125017-2f4d2b4d2089
	github.com/xiazeyin/gmgo v0.0.0-20220524080332-30d9ca7379fe
	github.com/xiazeyin/zcgolog v0.0.0-20220524064633-5ef357e6e1b9
	go.etcd.io/etcd/client/v2 v2.305.1 // indirect
)

replace (
	github.com/go-kit/kit => github.com/go-kit/kit v0.7.0
	// TODO:需要确认1.3.2版本的ReadInConfig内部将map key转为小写字母是否有影响。
	// 如果有影响，需要将viper版本改回 v0.0.0-20150908122457-1967d93db724 ，但会有一些测试案例编译不过。
	github.com/spf13/viper => github.com/spf13/viper v1.3.2
	// zlint与zcrypto版本必须匹配，否则zlint编译出错
	github.com/zmap/zcrypto => github.com/zmap/zcrypto v0.0.0-20190729165852-9051775e6a2e
	github.com/zmap/zlint => github.com/zmap/zlint v0.0.0-20190806154020-fd021b4cfbeb
)
