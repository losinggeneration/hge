package resource

type Pointer uintptr

type Resource struct {
	Pointer
}

// Loads a resource into memory from disk.
func New(filename string) (*Resource, uint32) {
	return &Resource{}, 0
}

// Deletes a previously loaded resource from memory.
func (r *Resource) Free() {
}

// Loads a resource, puts the loaded data into a byte array, and frees the data.
func LoadBytes(filename string) []byte {
	return []byte{}
}

// Loads a resource, puts the data into a string, and frees the data.
func LoadString(filename string) string {
	return ""
}

// Attaches a resource pack.
func AttachPack(filename string, a ...interface{}) bool {
	return false
}

// Removes a resource pack.
func RemovePack(filename string) {
}

// Removes all resource packs previously attached.
func RemoveAllPacks() {
}

// Builds absolute file path.
func MakePath(a ...interface{}) string {
	return ""
}

// Enumerates files by given wildcard.
func EnumFiles(a ...interface{}) string {
	return ""
}

// Enumerates folders by given wildcard.
func EnumFolders(a ...interface{}) string {
	return ""
}
