package types

const (
	// ModuleName defines the module name
	ModuleName = "socialapp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_socialapp"

	// ProfileKey defines the unique ID of a profile
	ProfileKey = "Profile/value/"

	// Posts defines the prefix of all Posts related values
	Posts = "Posts/"
	// PostKey defines the unique ID of a post owned by a profile
	PostKey = "value"

	//PostKeyCount defines the number of created posts by a profile
	PostCountKey = "count"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
