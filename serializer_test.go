package serializer

import (
	"testing"
	"github.com/liquidgecka/testlib"
)

func TestDeserializer(t *testing.T) {
	T := testlib.NewT(t)
	defer T.Finish()
	result, _ := SecureParse("3Ug3lpRAQzWgx5Cw4XHIP-wbVAk=PMfXno5D7681d4e3be7d564ec371698407ec2d6f", "encrypt_key", "validate_key")

	T.Equal(result, "\"test\"", "Not parsed")

}
//
func TestSerializer(t *testing.T) {
	T := testlib.NewT(t)
	defer T.Finish()

	result, _ := SecureStringify("test", "encrypt_key", "validate_key")
	result, _ = SecureParse(result, "encrypt_key", "validate_key")
	T.Equal(result, "test", "not equal")
}