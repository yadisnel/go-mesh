module github.com/yadisnel/go-ms/v2plugins/broker/segmentio/v2

go 1.13

require (
	github.com/google/uuid v1.1.1
	github.com/yadisnel/go-ms/v2 v2.0.0-alpha.3
	github.com/yadisnel/go-ms/v2plugins/broker/kafka/v2 v2.3.0
	github.com/yadisnel/go-ms/v2plugins/codec/segmentio/v2 v2.3.0
	github.com/segmentio/kafka-go v0.3.7
)

replace github.com/yadisnel/go-ms/v2plugins/codec/segmentio/v2 => ../../codec/segmentio
