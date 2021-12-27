module cart

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.0.0-20211224114320-e0de23c54615
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211224114320-e0de23c54615
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4 v4.0.0-20211224114320-e0de23c54615
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4 v4.0.0-20211224114320-e0de23c54615
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go-micro.dev/v4 v4.2.1
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
