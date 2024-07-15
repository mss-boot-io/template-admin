package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mss-boot-io/mss-boot-admin/cmd/migrate"
	"github.com/mss-boot-io/mss-boot-admin/cmd/server"
	"github.com/mss-boot-io/mss-boot-admin/pkg"
	"github.com/spf13/cobra"

	_ "admin/cmd/migrate/migration/custom"
	"admin/utils"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2024/1/25 10:57:11
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2024/1/25 10:57:11
 */

var rootCmd = &cobra.Command{
	Use:          "chainide-admin",
	Short:        "chainide-admin",
	SilenceUsage: true,
	Long:         "chainide-admin is a background management system developed by the chainide",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`chainide-admin `+utils.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
