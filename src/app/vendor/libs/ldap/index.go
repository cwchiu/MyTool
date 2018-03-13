package ldap

import (
	lib "gopkg.in/ldap.v2"
)

func Search(addr, dn, filter string) (*lib.SearchResult, error) {
	l, err := lib.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer l.Close()

	searchRequest := lib.NewSearchRequest(
		dn,
		lib.ScopeWholeSubtree,
		lib.NeverDerefAliases,
		0,
		10,
		false,
		filter,
		[]string{},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}
    
    return sr, nil
}
