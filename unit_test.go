package GetAllFolder

import "testing"

func TestListFolders(t *testing.T) {
	ret := ListFolders("F:\\整理")
	for _, d := range ret {
		t.Log(d)
	}
}
