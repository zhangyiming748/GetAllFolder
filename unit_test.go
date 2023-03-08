package GetAllFolder

import "testing"

func TestListFolders(t *testing.T) {
	ret := ListFolders("D:\\甄嬛传")
	for _, d := range ret {
		t.Log(d)
	}
}
