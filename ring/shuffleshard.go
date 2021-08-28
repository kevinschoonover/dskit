package ring

import (
	"crypto/md5"
	"encoding/binary"
	"math"
)

var (
	seedSeparator = []byte{0}
)

// shuffleShardSeed returns seed for random number generator, computed from provided identifier.
func shuffleShardSeed(identifier, zone string) int64 {
	// Use the identifier to compute an hash we'll use to seed the random.
	hasher := md5.New()
	hasher.Write(YoloBuf(identifier)) // nolint:errcheck
	if zone != "" {
		hasher.Write(seedSeparator) // nolint:errcheck
		hasher.Write(YoloBuf(zone)) // nolint:errcheck
	}
	checksum := hasher.Sum(nil)

	// Generate the seed based on the first 64 bits of the checksum.
	return int64(binary.BigEndian.Uint64(checksum))
}

// shuffleShardExpectedInstancesPerZone returns the number of instances that should be selected for each
// zone when zone-aware replication is enabled. The algorithm expects the shard size to be divisible
// by the number of zones, in order to have nodes balanced across zones. If it's not, we do round up.
func shuffleShardExpectedInstancesPerZone(shardSize, numZones int) int {
	return int(math.Ceil(float64(shardSize) / float64(numZones)))
}
