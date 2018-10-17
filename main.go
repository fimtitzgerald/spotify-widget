package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	getSongCmd := `tell application "Spotify" to name of current track as string`
	getArtistCmd := `tell application "Spotify" to artist of current track as string`
	getAlbumCmd := `tell application "Spotify" to album of current track as string`

	textView := tview.NewTextView()
	textView.SetBorder(true).SetTitle(" Spotify - Now Playing ")
	textView.SetTextAlign(1).SetChangedFunc(func() {
		app.Draw()
	})

	go func() {
		for {
			track := refresh(getSongCmd)
			artist := refresh(getArtistCmd)
			album := refresh(getAlbumCmd)
			songInfo := fmt.Sprintf("Arist: %s \nTrack: %s \nAlbum: %s", artist, track, album)
			display(textView, songInfo)
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	if err := app.SetRoot(textView, true).Run(); err != nil {
		panic(err)
	}
}

func currentSong(cmd string) []byte {
	out, err := exec.Command("osascript", "-e", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func display(view *tview.TextView, songInfo string) {
	text := fmt.Sprintf("%s", songInfo)
	view.SetText(text)
}

func refresh(cmd string) []byte {
	songInfo := currentSong(cmd)
	return songInfo
}
