//go:generate mockgen -package=mock -destination mock/tagger.go github.com/jcs/id3-go Tagger
package md

import (
	"fmt"
	"testing"

	"github.com/brettbuddin/grift/ffmpeg/md/mock"
	"github.com/golang/mock/gomock"
	id3v2 "github.com/jcs/id3-go/v2"
)

var (
	ftAlbum  = id3v2.V23FrameTypeMap["TALB"]
	ftArtist = id3v2.V23FrameTypeMap["TPE1"]
	ftTitle  = id3v2.V23FrameTypeMap["TIT2"]
	ftLength = id3v2.V23FrameTypeMap["TLEN"]
	ftYear   = id3v2.V23FrameTypeMap["TYER"]
)

func TestAddFrames(t *testing.T) {
	ctrl := gomock.NewController(t)
	tagger := mock.NewMockTagger(ctrl)

	track := Track{
		Title:  "Weird Sounds",
		Album:  "Phonogrifter",
		Artist: "Brett Buddin",
	}

	tagger.EXPECT().DeleteFrames("CTOC")
	tagger.EXPECT().DeleteFrames("CHAP")
	gomock.InOrder(
		tagger.EXPECT().AddFrames(textFrameMatcher{
			frameType: ftAlbum,
			text:      track.Album,
		}),
		tagger.EXPECT().AddFrames(textFrameMatcher{
			frameType: ftArtist,
			text:      track.Artist,
		}),
		tagger.EXPECT().AddFrames(textFrameMatcher{
			frameType: ftTitle,
			text:      track.Title,
		}),
		tagger.EXPECT().AddFrames(tocFrameMatcher{
			chapters: 2,
		}),
		tagger.EXPECT().AddFrames(chapterFrameMatcher{
			chapter: 1,
			start:   0,
			end:     150,
		}),
		tagger.EXPECT().AddFrames(chapterFrameMatcher{
			chapter: 2,
			start:   150,
			end:     300,
		}),
		tagger.EXPECT().AddFrames(textFrameMatcher{
			frameType: ftLength,
			text:      "300",
		}),
		tagger.EXPECT().AddFrames(textFrameMatcher{
			frameType: ftYear,
			text:      "2022",
		}),
	)

	l := List{
		{
			Filename: "chapter1.wav",
			Title:    "Chapter 1",
			Start:    0,
			End:      150,
		},
		{
			Filename: "chapter2.wav",
			Title:    "Chapter 2",
			Start:    150,
			End:      300,
		},
	}
	l.addFrames(tagger, track)
}

type textFrameMatcher struct {
	frameType id3v2.FrameType
	text      string
}

func (m textFrameMatcher) Matches(o interface{}) bool {
	textFrame, ok := o.(*id3v2.TextFrame)
	if !ok {
		return false
	}
	if textFrame.FrameType.Id() != m.frameType.Id() {
		return false
	}
	if textFrame.Text() != m.text {
		return false
	}
	return true
}

func (m textFrameMatcher) String() string {
	return fmt.Sprintf("FrameType=%s Text=%s", m.frameType.Id(), m.text)
}

type tocFrameMatcher struct {
	*id3v2.TOCFrame
	chapters int
}

func (m tocFrameMatcher) Matches(o interface{}) bool {
	tocFrame, ok := o.(*id3v2.TOCFrame)
	if !ok {
		return false
	}

	if tocFrame.FrameType.Id() != id3v2.V23FrameTypeMap["CTOC"].Id() {
		return false
	}

	for i, c := range tocFrame.ChildElements {
		if c != fmt.Sprintf("chp%d", i+1) {
			return false
		}
	}

	return true
}

func (m tocFrameMatcher) String() string {
	return fmt.Sprintf("FrameType=TOC Chapters=%d", m.chapters)
}

type chapterFrameMatcher struct {
	chapter    int
	start, end uint32
}

func (m chapterFrameMatcher) Matches(o interface{}) bool {
	chapterFrame, ok := o.(*id3v2.ChapterFrame)
	if chapterFrame.Element != fmt.Sprintf("chp%d", m.chapter) {
		return false
	}
	if chapterFrame.Id() != id3v2.V23FrameTypeMap["CHAP"].Id() {
		return false
	}
	if chapterFrame.StartTime != m.start {
		return false
	}
	if chapterFrame.EndTime != m.end {
		return false
	}
	return ok
}

func (m chapterFrameMatcher) String() string {
	return fmt.Sprint("FrameType=")
}
