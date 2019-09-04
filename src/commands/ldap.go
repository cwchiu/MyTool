package commands

import (
	"github.com/spf13/cobra"
    "github.com/cwchiu/MyTool/libs/ldap"
)


func setupSearchCommand(rootCmd *cobra.Command) {
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
            sr, err := ldap.Search(addr, basedn, args[0])
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

func init() {
	cmd := &cobra.Command{Use: "ldap", Short: "ldap api"}

	setupSearchCommand(cmd)

	rootCmd.AddCommand(cmd)
}
