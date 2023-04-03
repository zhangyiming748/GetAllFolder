package GetAllFolder

import "testing"

func TestListFolders(t *testing.T) {
	ret := ListFolders("/Users/zen/Downloads")
	for _, d := range ret {
		t.Log(d)
	}
}
