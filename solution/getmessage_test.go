package getmessage_test

import (
	"testing"


)

func TestWorld(t *testing.T) {
	if getmessage.GetMessage() != "Hello 🗺️ !" {
		t.Fatal ("Wrong message :((")
		getmessage.GetMessage ()
		
	}
}