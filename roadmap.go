package roadmap

import (
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Roadmap struct {
	Title       string          `json:"title"`
	Description string          `json:"description,omitempty"`
	Authors     []Author        `json:"authors,omitempty"`
	Timeline    []TimelineEntry `json:"timeline,omitempty"`
	Milestones  []Milestone     `json:"milestones,omitempty"`
}

// Parse will convert a provided input buffer into a roadmap structure.
func Parse(in []byte) (*Roadmap, error) {
	r := Roadmap{}

	if err := yaml.Unmarshal(in, &r); err != nil {
		return nil, errors.Wrap(err, "roadmap: unable to parse roadmap ")
	}

	return &r, nil
}

type Author struct {
	Name    string `json:"name"`
	Contact string `json:"contact,omitempty"`
}

type TimelineEntry struct {
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
}

type Milestone struct {
	Title        string        `json:"title"`
	Description  string        `json:"description,omitempty"`
	Deliverables []Deliverable `json:"deliverables,omitempty"`
}

type Deliverable struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Reference   string `json:"reference,omitempty"`
	Requirement string `json:"requirement,omitempty"`
	State       string `json:"state,omitempty"`
}
