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
	Filename string   `hcl:",label"`
	Title    string   `hcl:"title"`
	Gain     float64  `hcl:"gain,optional"`
	Start    *float64 `hcl:"start,optional"`
	Stop     *float64 `hcl:"stop,optional"`
}
