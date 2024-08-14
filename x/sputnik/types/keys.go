package types

const (
	// ModuleName defines the module name
	ModuleName = "sputnik"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sputnik"
)

var (
	ParamsKey = []byte("p_sputnik")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
