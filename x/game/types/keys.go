package types

const (
	// ModuleName defines the module name
	ModuleName = "game"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_game"
)

var (
	ParamsKey = []byte("p_game")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
