package mi

import (
	"encoding/json"
	"fmt"
	"github.com/ApheXOK/mi/convert"
	"path/filepath"
	"strings"
)

func Parse(data []byte) (string, error) {
	var mediaInfo MediaInfo
	if err := json.Unmarshal(data, &mediaInfo); err != nil {
		return "", err
	}
	if len(mediaInfo.Media.Track) == 0 {
		return "", fmt.Errorf("invalid")
	}
	var mediaInfoText strings.Builder

	ref := mediaInfo.Media.Ref
	mediaInfoText.WriteString("RELEASE NAME..: " + filepath.Base(ref) + "\n")
	for _, r := range mediaInfo.Media.Track {
		if r.Type == "General" {
			mediaInfoText.WriteString("RELEASE DATE..: " + r.EncodedDate + "\n")
			size, err := convert.Size(r.FileSize)
			if err != nil {
				return "", err
			}
			mediaInfoText.WriteString("RELEASE SIZE..: " + size + "\n")
			duration, err := convert.Duration(r.Duration)
			if err != nil {
				return "", err
			}
			mediaInfoText.WriteString("RUNTIME.......: " + duration + "\n")
		}
		if r.Type == "Video" {
			mediaInfoText.WriteString("VIDEO CODEC...: " + r.Format + ", " +
				r.FormatProfile + "@L" + r.FormatLevel + "\n")
			mediaInfoText.WriteString("FRAMERATE.....: " + r.FrameRate + " FPS" + "\n")
			mediaInfoText.WriteString("BITRATE.......: " + convert.BitRate(r.BitRate) + "\n")
			mediaInfoText.WriteString("RESOLUTION....: " + r.Width + "x" + r.Height + "\n")
		}
		if r.Type == "Audio" {
			mediaInfoText.WriteString("AUDIO.........: " + r.Format + " " +
				r.Channels + " Channels " + "@" + convert.BitRate(r.BitRate) + " " +
				convert.AudioLang(r.Language) + "\n")
		}
		if r.Type == "Text" {
			mediaInfoText.WriteString("SUBTITLES.....: " + convert.SubtitleLang(r.Language, r.Title) + "\n")
		}

	}
	return mediaInfoText.String(), nil
}
