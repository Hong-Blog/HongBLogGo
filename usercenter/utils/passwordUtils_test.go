package utils

import "testing"

func TestPasswordEncrypt(t *testing.T) {
	user := "root"
	encryptedPassword := "CGUx1FN++xS+4wNDFeN6DA=="

	encrypt := PasswordEncrypt("123456", user)

	if encrypt != encryptedPassword {
		t.Errorf("%s is not %s", encrypt, encryptedPassword)
	}
}
