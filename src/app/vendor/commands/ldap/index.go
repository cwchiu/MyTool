package ldap

import (
	// "fmt"
	"github.com/spf13/cobra"
	"gopkg.in/ldap.v2"
)

func SetupSearchCommand(rootCmd *cobra.Command) {
	var basedn string
	var addr string

	cmd := &cobra.Command{
		Use:   "search <filter>",
		Short: "LDAP Search",
        Long: "example: MyTool -a 127.0.0.1:389 -b dc=htc \"(&(cn=張花*)(|(objectclass=person)))\"",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <filter>")
			}

			if addr == "" {
				panic("required -a <LDAP Server Address> ")
			}

			l, err := ldap.Dial("tcp", addr)
			if err != nil {
				panic(err)
			}
			defer l.Close()

			searchRequest := ldap.NewSearchRequest(
				basedn,
				ldap.ScopeWholeSubtree,
				ldap.NeverDerefAliases,
				0,
				10,
				false,
				args[0],
				[]string{},
				nil,
			)

			sr, err := l.Search(searchRequest)
			if err != nil {
				panic(err)
			}
			for _, entry := range sr.Entries {
				// fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
				// fmt.Println( entry.GetAttributeValue("mail"))
				// fmt.Println( len (entry.Attributes))
				entry.PrettyPrint(4)
			}

		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:389", "LDAP Server Address")
	cmd.Flags().StringVarP(&basedn, "base-dn", "b", "", "LDAP Base DN")
	rootCmd.AddCommand(cmd)
}
