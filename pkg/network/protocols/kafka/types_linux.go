// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -I ../../ebpf/c -I ../../../ebpf/c -fsigned-char types.go

package kafka

const (
	TopicNameBuckets = 0xa
)

type ConnTuple struct {
	Saddr_h  uint64
	Saddr_l  uint64
	Daddr_h  uint64
	Daddr_l  uint64
	Sport    uint16
	Dport    uint16
	Netns    uint32
	Pid      uint32
	Metadata uint32
}

type EbpfTx struct {
	Tup         ConnTuple
	Transaction KafkaTransaction
}

type KafkaTransactionKey struct {
	Tuple     ConnTuple
	Id        int32
	Pad_cgo_0 [4]byte
}
type KafkaTransaction struct {
	Request_started     uint64
	Records_count       uint32
	Request_api_key     uint8
	Request_api_version uint8
	Topic_name_size     uint8
	Topic_name          [80]byte
	Pad_cgo_0           [1]byte
}

type KafkaResponseContext struct {
	Transaction                 KafkaTransaction
	State                       uint8
	Remainder                   uint8
	Varint_position             uint8
	Partition_state             uint8
	Remainder_buf               [4]int8
	Record_batches_num_bytes    int32
	Record_batch_length         int32
	Expected_tcp_seq            uint32
	Carry_over_offset           int32
	Partitions_count            uint32
	Varint_value                uint32
	Record_batches_arrays_idx   uint32
	Record_batches_arrays_count uint32
}

type RawKernelTelemetry struct {
	Name_size_buckets [10]uint64
}
