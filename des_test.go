package bosskg

import "testing"

func TestDESBase64_RoundTrip(t *testing.T) {
	t.Parallel()

	key, err := DeriveDESKey("524538F6")
	if err != nil {
		t.Fatal(err)
	}
	plain := []byte("重瞳本是无敌路，何须再借他人骨！")

	ct, err := EncryptDESBase64(plain, key)
	if err != nil {
		t.Fatal(err)
	}
	got, err := DecryptDESBase64(ct, key)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(plain) {
		t.Fatalf("got=%q want=%q", got, plain)
	}
}
