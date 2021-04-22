package roadmap_test

import (
	"testing"
	"time"

	"github.com/SierraSoftworks/roadmap"
	"github.com/stretchr/testify/assert"
)

func TestAuthorID(t *testing.T) {
	a := &roadmap.Author{
		Name:    "Benjamin Pannell",
		Contact: "contact@sierrasoftworks.com",
	}

	assert.Contains(t, a.ID(), "Benjamin_Pannell_", "the ID should include the normalized name")
}

func TestTimelineID(t *testing.T) {
	te := &roadmap.TimelineEntry{
		Date:        time.Date(2021, 04, 22, 00, 0, 0, 0, time.UTC),
		Title:       "Timeline Entry Title",
		Description: "Timeline Entry Description",
	}

	assert.Contains(t, te.ID(), "Timeline_Entry_Title_", "the ID should include the normalized title")
}

func TestMilestoneID(t *testing.T) {
	m := &roadmap.Milestone{
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
	}

	assert.Contains(t, m.ID(), "Milestone_Title_", "the ID should include the normalized title")
}

func TestDeliverableID(t *testing.T) {
	d := &roadmap.Deliverable{
		Title:       "Deliverable Title",
		Description: "Deliverable Description",
		State:       "TODO",
		Requirement: "SHOULD",
		Reference:   "https://reference.example.com",
	}

	assert.Contains(t, d.ID(), "Deliverable_Title_", "the ID should include the normalized title")
}
