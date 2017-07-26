package windows

import (
	"syscall"
	"github.com/spf13/cobra"
)


func SetupLockCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "lock",
		Short: "鎖定螢幕",
		Run: func(cmd *cobra.Command, args []string) {
            var mod = syscall.NewLazyDLL("user32.dll")
            var proc = mod.NewProc("LockWorkStation")
            proc.Call()
		},
	}

	rootCmd.AddCommand(cmd)

}
