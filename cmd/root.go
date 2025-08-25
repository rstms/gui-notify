/*
Copyright © 2024 Matt Krueger <mkrueger@rstms.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"fmt"
	"github.com/rstms/gui-notify/notify"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: "0.0.9",
	Use:     "notify",
	Short:   "display a notification on the user desktop",
	Long: `
generate a user notification message using the appropriate OS interface
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if ViperGetString("audio") == "help" {
			for _, name := range notify.AudioNames {
				fmt.Println(name)
			}
			return
		}
		message := "notification"
		if len(args) > 0 {
			message = args[0]
		}
		err := notify.Send(message)
		cobra.CheckErr(err)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	CobraInit(rootCmd)
	OptionSwitch(rootCmd, "no-wait", "", "fire and forget")
	OptionSwitch(rootCmd, "force", "", "bypass confirmation prompts")
	OptionString(rootCmd, "id", "i", rootCmd.Name(), "notification AppID")
	OptionString(rootCmd, "title", "t", "", "notificaton title")
	OptionSwitch(rootCmd, "long", "L", "request long duration")
	OptionString(rootCmd, "audio", "a", "silent", "audio name (help for list)")
	OptionString(rootCmd, "icon", "I", "", "icon file (PNG format)")
	OptionSwitch(rootCmd, "loop", "", "loop audio")
}
