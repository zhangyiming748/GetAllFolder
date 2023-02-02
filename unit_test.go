package GetAllFolder

import "testing"

func TestListFolders(t *testing.T) {
	ret := ListFolders("/Users/zen/Downloads")
	t.Log(ret)
}
