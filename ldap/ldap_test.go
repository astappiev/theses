package ldap

import (
	"log"
	"os"
	"testing"
)

func TestFindUser(t *testing.T) {
	user, err := FindUser(os.Getenv("LDAP_TEST_USERNAME"), os.Getenv("LDAP_TEST_PASSWORD"))
	if err != nil {
		t.Error(err)
	}

	log.Print(user)
}
