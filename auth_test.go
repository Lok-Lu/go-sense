package sense

import "testing"

func TestCheckJwtToken(t *testing.T) {
	tokenString, err := EncodeJwtToken("", "")
	if err != nil {
		t.Error(err)
	}
	t.Log(tokenString)
	t.Log(CheckJwtToken(tokenString, ""))
}
