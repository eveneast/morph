package db

import "encoding/binary"

var (
	syncedL1HeightKey      = []byte("LastSyncedL1Height")
	L1MessagePrefix        = []byte("l1")
	BatchBlockNumberPrefix = []byte("batch")

	derivationL1HeightKey = []byte("LastDerivationL1Height")
)

// encodeEnqueueIndex encodes an L1 enqueue index as big endian uint64
func encodeEnqueueIndex(index uint64) []byte {
	enc := make([]byte, 8)
	binary.BigEndian.PutUint64(enc, index)
	return enc
}

// L1MessageKey = L1MessagePrefix + enqueueIndex (uint64 big endian)
func L1MessageKey(enqueueIndex uint64) []byte {
	return append(L1MessagePrefix, encodeEnqueueIndex(enqueueIndex)...)
}

// BatchBlockNumberKey = BatchBlockNumberKeyPrefix + batchIndex (uint64 big endian)
func BatchBlockNumberKey(batchIndex uint64) []byte {
	return append(BatchBlockNumberPrefix, encodeEnqueueIndex(batchIndex)...)
}
