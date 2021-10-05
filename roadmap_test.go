package roadmap_test

import (
	"testing"
	"time"

	"github.com/SierraSoftworks/roadmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getTestSchema() string {
	return `
title: Roadmap Title
description: Roadmap Description

authors:
  - name: Author Name
    contact: Author Contact

timeline:
  - date: 2021-01-01
    title: Timeline Entry Title
    description: TImeline Entry Description

objectives:
  - title: Objective Title
    description: Objective Description

milestones:
  - title: Milestone Title
    description: Milestone Description
    deliverables:
      - title: Deliverable Title
        state: TODO
        requirement: MUST
        description: Deliverable Description
`
}

func TestParse(t *testing.T) {
	in := getTestSchema()

	r, err := roadmap.Parse([]byte(in))
	require.NoError(t, err, "the parser should not return an error")
	require.NotNil(t, r, "the road map should not be nil")

	assert.Equal(t, "Roadmap Title", r.Title)
}

func TestValidate(t *testing.T) {
	in := getTestSchema()

	r, err := roadmap.Parse([]byte(in))
	require.NoError(t, err, "the parser should not return an error")
	require.NotNil(t, r, "the road map should not be nil")

	require.NoError(t, r.Validate(), "the validation should not return an error")
}

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

func TestObjectiveID(t *testing.T) {
	te := &roadmap.Objective{
		Title:       "Objective Title",
		Description: "Objective Description",
	}

	assert.Contains(t, te.ID(), "Objective_Title_", "the ID should include the normalized title")
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
