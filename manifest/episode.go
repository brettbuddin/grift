package manifest

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

func Parse(filename string) (Episode, error) {
	var m Episode
	if err := hclsimple.DecodeFile(filename, nil, &m); err != nil {
		return Episode{}, err
	}
	if len(m.Chapters) == 0 {
		return Episode{}, fmt.Errorf("chapters are required")
	}
	if m.Samplerate == 0 {
		m.Samplerate = 44100
	}
	if m.Bitrate == 0 {
		m.Bitrate = 128
	}
	for i, ch := range m.Chapters {
		if ch.Filename == "" {
			return Episode{}, fmt.Errorf("chapter %d: filename is empty", i)
		}
		if ch.Title == "" {
			return Episode{}, fmt.Errorf("chapter %d: title is empty", i)
		}
		if ch.Start < 0 {
			return Episode{}, fmt.Errorf("chapter %d: start is negative", i)
		}
		if ch.Stop < 0 {
			return Episode{}, fmt.Errorf("chapter %d: stop is negative", i)
		}
		if ch.Stop != 0 && ch.Start >= ch.Stop {
			return Episode{}, fmt.Errorf("chapter %d: start is after stop", i)
		}
	}
	return m, nil
}

type Episode struct {
	Title      string     `hcl:"title"`
	Album      string     `hcl:"album"`
	Authors    []string   `hcl:"authors"`
	Marker     string     `hcl:"marker,optional"`
	Chapters   []Chapter  `hcl:"chapter,block"`
	Samplerate uint       `hcl:"samplerate,optional"`
	Bitrate    uint       `hcl:"bitrate,optional"`
	Normalize  *Normalize `hcl:"normalize,block"`
}

type Normalize struct {
	LoudnessTarget float64 `hcl:"loudness_target"`
	LoudnessRange  float64 `hcl:"loudness_range"`
	TruePeak       float64 `hcl:"true_peak"`
}

type Chapter struct {
	Filename string  `hcl:",label"`
	Title    string  `hcl:"title"`
	Gain     float64 `hcl:"gain,optional"`
	Start    float64 `hcl:"start,optional"`
	Stop     float64 `hcl:"stop,optional"`
}
