package manifest

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	cases := []struct {
		filename    string
		expected    Episode
		expectedErr error
	}{
		{
			filename: "testdata/ok_minimum.hcl",
			expected: Episode{
				Title:      "Weird Sounds",
				Album:      "Phonogrifter",
				Authors:    []string{"Brett Buddin"},
				Samplerate: 44100,
				Bitrate:    128,
				Chapters: []Chapter{
					{
						Filename: "ZOOM0012_TrLR.WAV",
						Title:    "Chapter 1",
					},
					{

						Filename: "ZOOM0013_TrLR.WAV",
						Title:    "Chapter 2",
					},
				},
			},
		},
		{
			filename: "testdata/ok_optionals_present.hcl",
			expected: Episode{
				Title:      "Weird Sounds",
				Album:      "Phonogrifter",
				Authors:    []string{"Brett Buddin"},
				Marker:     "beep.wav",
				Samplerate: 48000,
				Bitrate:    256,
				Normalize: &Normalize{
					LoudnessTarget: -24,
					LoudnessRange:  7,
					TruePeak:       -2,
				},
				Chapters: []Chapter{
					{
						Filename: "ZOOM0012_TrLR.WAV",
						Title:    "Chapter 1",
					},
					{

						Filename: "ZOOM0013_TrLR.WAV",
						Title:    "Chapter 2",
						Gain:     -2,
						Start:    3,
						Stop:     8,
					},
				},
			},
		},
		{
			filename:    "testdata/error_episode_stop_before_start.hcl",
			expectedErr: fmt.Errorf("chapter 0: start is after stop"),
		},
		{
			filename:    "testdata/error_episode_negative_start.hcl",
			expectedErr: fmt.Errorf("chapter 0: start is negative"),
		},
		{
			filename:    "testdata/error_episode_negative_stop.hcl",
			expectedErr: fmt.Errorf("chapter 0: stop is negative"),
		},
		{
			filename:    "testdata/error_episode_empty_filename.hcl",
			expectedErr: fmt.Errorf("chapter 0: filename is empty"),
		},
		{
			filename:    "testdata/error_episode_empty_title.hcl",
			expectedErr: fmt.Errorf("chapter 0: title is empty"),
		},
	}

	for _, c := range cases {
		t.Run(c.filename, func(t *testing.T) {
			ep, err := Parse(c.filename)
			if c.expectedErr != nil && err != nil {
				if !cmp.Equal(c.expectedErr.Error(), err.Error()) {
					t.Error(cmp.Diff(c.expectedErr.Error(), err.Error()))
				}
			} else if err != nil {
				t.Errorf("expected no error, but received: %v", err)
			}
			if !cmp.Equal(c.expected, ep) {
				t.Error(cmp.Diff(c.expected, ep))
			}
		})
	}
}
