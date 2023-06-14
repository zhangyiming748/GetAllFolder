package GetAllFolder

import "testing"

func TestListFolders(t *testing.T) {
	ret := List("/Users/zen/Downloads")
	for _, d := range ret {
		t.Log(d)
	}
}
