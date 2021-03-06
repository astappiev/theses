package user

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

var (
	searchDN = "uid=%s,ou=users"
	filter   = "(objectClass=*)"
)

func bindAndSearch(l *ldap.Conn, username, password string) (*ldap.Entry, error) {
	userDN := fmt.Sprintf(searchDN+","+os.Getenv("LDAP_BASE"), ldap.EscapeFilter(username))
	err := l.Bind(userDN, password)
	if err != nil {
		return nil, fmt.Errorf("bind error: %s", err)
	}

	searchReq := ldap.NewSearchRequest(
		userDN,
		ldap.ScopeBaseObject, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		1,
		0,
		false,
		filter,
		[]string{"uidNumber", "mail", "cn", "givenName", "sn"},
		nil,
	)
	result, err := l.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("search error: %s", err)
	}

	if len(result.Entries) > 0 {
		return result.Entries[0], nil
	} else {
		return nil, fmt.Errorf("couldn't fetch search entries")
	}
}

func FindUser(username, password string) (*User, error) {
	if strings.Contains(username, "@") {
		username = strings.Split(username, "@")[0]
	}

	l, err := ldap.DialURL(os.Getenv("LDAP_URL"))
	if err != nil {
		return nil, err
	}
	defer l.Close()

	result, err := bindAndSearch(l, username, password)
	if err != nil {
		return nil, err
	}

	user := User{}
	if id, err := strconv.Atoi(result.GetAttributeValue("uidNumber")); err == nil {
		user.Id = uint(id)
	}
	user.Email = result.GetAttributeValue("mail")
	user.FirstName = result.GetAttributeValue("givenName")
	user.LastName = result.GetAttributeValue("sn")
	user.FullName = result.GetAttributeValue("cn")
	return &user, nil
}
