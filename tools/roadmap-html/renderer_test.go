package main

import (
	"testing"
	"time"

	"github.com/SierraSoftworks/roadmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	r := &roadmap.Roadmap{
		Title:       "Road Map Title",
		Description: "Road Map Description",
		Authors: []*roadmap.Author{
			{
				Name:    "Author",
				Contact: "contact@author",
			},
		},
		Timeline: []*roadmap.TimelineEntry{
			{
				Date:        time.Date(2021, 04, 22, 00, 0, 0, 0, time.UTC),
				Title:       "Timeline Entry Title",
				Description: "Timeline Entry Description",
			},
		},
		Milestones: []*roadmap.Milestone{
			{
				Title:       "Milestone Title",
				Description: "Milestone Description",
				Deliverables: []*roadmap.Deliverable{
					{
						Title:       "Deliverable Title",
						Description: "Deliverable Description",
						State:       "TODO",
						Requirement: "SHOULD",
						Reference:   "https://reference.example.com",
					},
				},
			},
		},
	}

	dot, err := render(r)
	require.NoError(t, err, "no error should have been returned")

	t.Log(dot)

	assert.Contains(t, dot, "label=\"Road Map Title\";", "the roadmap title should be set")
	assert.Contains(t, dot, "tooltip=\"Road Map Description\";", "the roadmap description should be set")
}
