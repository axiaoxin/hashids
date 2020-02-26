package hashids

import "testing"

func TestAll(t *testing.T) {
	h, err := New("salt", 8)
	if err != nil {
		t.Error(err)
	}

	var id int64 = 1
	str, err := h.Encode(id)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%d encode to %s", id, str)

	id, err = h.Decode(str)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s decode to %d", str, id)

}
