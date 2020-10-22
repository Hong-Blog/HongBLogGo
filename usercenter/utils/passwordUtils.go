package utils

const security_key = "929123f8f17944e8b0a531045453e1f1"

func PasswordEncrypt(password string, salt string) string {
	hash := Md5(salt + security_key)
	encrypt := AesEncrypt(password, hash)
	return encrypt
}

func PasswordDecrypt(password string, salt string) string {
	return ""
}
