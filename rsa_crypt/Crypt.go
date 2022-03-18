package rsa_crypt

type ICrypt interface {
	Encrypt(data string) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}
