package modpacksch

import "testing"

func TestClient_ErrorHandling(t *testing.T) {
	client := NewClient(nil)

	_, err := client.Packs.GetPack(99)
	if err == nil {
		t.Error("client.Packs.GetPack(99) should error, perhaps a pack with that ID exists now?")
		return
	}
	t.Logf("Error: %s", err)
}
