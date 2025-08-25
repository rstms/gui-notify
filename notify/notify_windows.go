//go:build windows

package notify

import (
	_ "embed"
	"github.com/go-toast/toast"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const ICON_ID = "7d85101a-b502-43f3-b8a9-837a9d85745b"

//go:embed notify.png
var iconData []byte

func Send(message string) error {
	ViperSetDefault("id", ProgramName())
	ViperSetDefault("audio", "silent")
	audioName := ViperGetString("audio")
	audio, err := toast.Audio(audioName)
	if err != nil {
		return Fatal(err)
	}
	durationName := "short"
	loop := ViperGetBool("loop")
	if strings.HasPrefix(audioName, "looping") {
		loop = true
	}
	if ViperGetBool("long") || loop {
		durationName = "long"
	}
	duration, err := toast.Duration(durationName)
	if err != nil {
		return Fatal(err)
	}
	log.Printf("audio: %+v\n", audio)
	n := toast.Notification{
		AppID:    ViperGetString("id"),
		Duration: duration,
		Title:    ViperGetString("title"),
		Message:  message,
		Icon:     ViperGetString("icon"),
		Audio:    audio,
		Loop:     loop,
	}
	if n.Icon == "" {
		n.Icon = filepath.Join(os.TempDir(), ICON_ID+".png")
		err = os.WriteFile(n.Icon, iconData, 0600)
		if err != nil {
			return Fatal(err)
		}
	}
	log.Printf("Sending: %+v\n", n)
	err = n.Push()
	if err != nil {
		return Fatal(err)
	}
	return nil
}
