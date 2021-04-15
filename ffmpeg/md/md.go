package md

import (
	"fmt"
	"time"

	"github.com/jcs/id3-go"
	id3v2 "github.com/jcs/id3-go/v2"
)

// WriteChapters writes ID3v2 chapter information to the MP3 file.
func WriteChapters(filename string, track Track, chapters List) error {
	f, err := id3.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	chapters.addFrames(f, track)
	return nil
}

// Track about the track
type Track struct {
	Title, Album, Artist string
}

// List is a list of Chapters
type List []Chapter

// addFrame writes a table of contents frame, the chapter frames, and a total
// length frame to the ID3 information.
func (l List) addFrames(f *id3.File, track Track) {
	// Remove previous table of contents and chapter frames.
	f.DeleteFrames("CTOC")
	f.DeleteFrames("CHAP")

	f.AddFrames(id3v2.NewTextFrame(
		id3v2.V23FrameTypeMap["TALB"],
		track.Album,
	))
	f.AddFrames(id3v2.NewTextFrame(
		id3v2.V23FrameTypeMap["TPE1"],
		track.Artist,
	))
	f.AddFrames(id3v2.NewTextFrame(
		id3v2.V23FrameTypeMap["TIT2"],
		track.Title,
	))

	var elements []string
	for i := range l {
		elements = append(elements, fmt.Sprintf("chp%d", i+1))
	}

	// Write table of contents.
	f.AddFrames(id3v2.NewTOCFrame(
		id3v2.V23FrameTypeMap["CTOC"],
		"toc",
		true,
		true,
		elements,
	))

	// Add all chapters.
	for i, c := range l {
		c.addFrames(f, i)
	}

	// Book-end with the total length.
	f.AddFrames(id3v2.NewTextFrame(
		id3v2.V23FrameTypeMap["TLEN"],
		fmt.Sprintf("%d", l[len(l)-1].End),
	))

	f.AddFrames(id3v2.NewTextFrame(
		id3v2.V23FrameTypeMap["TYER"],
		time.Now().Format("2006"),
	))
}

// Chapter represents an individual audio file to be concatenated with the
// greater whole.
type Chapter struct {
	Filename string `json:"-"`
	Title    string
	Start    uint32
	End      uint32
}

// addFrames writes a chapter frame to the ID3 information.
func (c Chapter) addFrames(f *id3.File, offset int) {
	f.AddFrames(id3v2.NewChapterFrame(
		id3v2.V23FrameTypeMap["CHAP"],
		fmt.Sprintf("chp%d", offset+1),
		uint32(c.Start),
		uint32(c.End),
		0,
		0,
		true,
		c.Title,
		"",
		"",
	))
}
