// go-common local proxy functions

package cmd

import (
	common "github.com/rstms/go-common"
)


func OptionKey(cobraCmd interface{}, key string) string {
	return common.OptionKey(cobraCmd, key)
}

func OptionSwitch(cobraCmd interface{}, name, flag, description string) {
	common.OptionSwitch(cobraCmd, name, flag, description)
}

func OptionString(cobraCmd interface{}, name, flag, defaultValue, description string) {
	common.OptionString(cobraCmd, name, flag, defaultValue, description)
}

func CobraInit(cobraRootCmd interface{}) {
	common.CobraInit(cobraRootCmd)
}

func Init(name, version string) {
	common.Init(name, version)
}

func ProgramName() string {
	return common.ProgramName()
}

func ProgramVersion() string {
	return common.ProgramVersion()
}

func CheckErr(err error) {
	common.CheckErr(err)
}

func OpenLog() {
	common.OpenLog()
}

func CloseLog() {
	common.CloseLog()
}

func FormatJSON(v any) string {
	return common.FormatJSON(v)
}

func ConfigString(header bool) string {
	return common.ConfigString(header)
}

func ConfigInit(allowClobber bool) string {
	return common.ConfigInit(allowClobber)
}

func ConfigEdit() {
	common.ConfigEdit()
}

func Confirm(prompt string) bool {
	return common.Confirm(prompt)
}

func Fatal(err error) error {
	return common.Fatal(err)
}

func Fatalf(format string, args ...interface{}) error {
	return common.Fatalf(format, args...)
}

func Warning(format string, args ...interface{}) {
	common.Warning(format, args...)
}

func HexDump(data []byte) string {
	return common.HexDump(data)
}

func IsDir(path string) bool {
	return common.IsDir(path)
}

func IsFile(pathname string) bool {
	return common.IsFile(pathname)
}

func TildePath(path string) (string, error) {
	return common.TildePath(path)
}

func Sendmail(to, from, subject string, body []byte, smtpServer *SMTPServer) error {
	return common.Sendmail(to, from, subject, body, smtpServer)
}

func Expand(pathname string) string {
	return common.Expand(pathname)
}

func ViperKey(name string) string {
	return common.ViperKey(name)
}

func ViperGetBool(key string) bool {
	return common.ViperGetBool(key)
}

func ViperGetString(key string) string {
	return common.ViperGetString(key)
}

func ViperGetStringSlice(key string) []string {
	return common.ViperGetStringSlice(key)
}

func ViperGetInt(key string) int {
	return common.ViperGetInt(key)
}

func ViperGetInt64(key string) int64 {
	return common.ViperGetInt64(key)
}

func ViperSet(key string, value any) {
	common.ViperSet(key, value)
}

func ViperSetDefault(key string, value any) {
	common.ViperSetDefault(key, value)
}
