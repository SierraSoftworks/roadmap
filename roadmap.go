package roadmap

import (
	"crypto/sha256"
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

//go:embed roadmap.schema.json
var roamdapSchema string

// Schema will return the raw JSON Schema used to validate a road map file
// for this version of the road map library.
func Schema() string {
	return roamdapSchema
}

// Parse will convert a provided input buffer into a roadmap structure.
func Parse(in []byte) (*Roadmap, error) {
	r := Roadmap{}

	if err := yaml.Unmarshal(in, &r); err != nil {
		return nil, errors.Wrap(err, "roadmap: unable to parse roadmap ")
	}

	r.ensureDefaults()

	return &r, nil
}

// A Roadmap describes a series of milestones and important dates which combine to
// outline a planned sequence of execution for a given project or team.
type Roadmap struct {
	// The Title is used to briefly introduce the road map to a reader
	Title string `json:"title"`

	// The Description is a markdown formatted body of text explaining the road map
	// and any context that a reader should have when consuming it.
	Description string `json:"description,omitempty"`

	// The Authors list provides the names and contact details for each of the individuals
	// or teams involved in defining the road map. They act as points of contact should questions
	// about the road map arise in future.
	Authors []*Author `json:"authors,omitempty"`

	// The Objectives list contains the high level goals that the team is working towards
	// and should inform both the content and prioritization of deliverables in each milestone.
	Objectives []*Objective `json:"objectives,omitempty"`

	// The Timeline lists important dates which are of relevance to the execution of this
	// road map. They are intentionally separated from milestones as we expect that milestones
	// are executed in sequence and might be delayed or accelerated based on external factors.
	Timeline []*TimelineEntry `json:"timeline,omitempty"`

	// The Milestones act as a series of high-level progress indicators. These allow a team
	// to visualize where they are on the road map without being overly constrained to the
	// specific execution.
	Milestones []*Milestone `json:"milestones,omitempty"`
}

// Validate will check whether the provided road map is valid according to the embedded
// JSON Schema.
func (r *Roadmap) Validate() error {
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020

	compiler.AddResource("roadmap.schema.json", strings.NewReader(Schema()))

	schema, err := compiler.Compile("roadmap.schema.json")
	if err != nil {
		return err
	}

	j, err := json.Marshal(r)
	if err != nil {
		return err
	}

	var raw map[string]interface{}
	err = json.Unmarshal(j, &raw)
	if err != nil {
		return err
	}

	err = schema.Validate(raw)
	if err != nil {
		return err
	}

	return nil
}

func (r *Roadmap) ensureDefaults() {
	for _, a := range r.Authors {
		a.ensureDefaults()
	}

	for _, t := range r.Timeline {
		t.ensureDefaults()
	}

	for _, m := range r.Milestones {
		m.ensureDefaults()
	}
}

// An Author is an individual or team which has contributed to the road map and which
// may be able to provide additional context if required.
type Author struct {
	// The Name of an individual or team which has contributed to this road map.
	Name string `json:"name"`

	// The Contact details for the individual or team, should a reader wish to contact them.
	// This may be an email address, IM handle or other method of contact.
	Contact string `json:"contact,omitempty"`
}

// ID will return a deterministic identifier for this author which may be used to uniquely identify them.
func (a *Author) ID() string {
	return id(a.Name, a.Contact)
}

func (a *Author) ensureDefaults() {

}

// A TimelineEntry describes an important date which is relevant to the road map. This
// may be a delivery date, quarterly milestone or any other point-in-time reference.
type TimelineEntry struct {
	// The Date associated with this timeline entry.
	Date time.Time `json:"date"`

	// The Title provides a brief name for the timeline entry to succinctly convey meaning to a reader.
	Title string `json:"title"`

	// The Description is a markdown formatted body of text providing additional context
	// on this timeline entry to any reader who needs it.
	Description string `json:"description,omitempty"`
}

// ID will return a deterministic identifier for this timeline entry which is based on its title and date.
func (t *TimelineEntry) ID() string {
	return id(t.Title, t.Date)
}

func (t *TimelineEntry) ensureDefaults() {

}

// A Milestone is used to describe a high-level progress indicator for a team or project.
// It is composed of a series of deliverables and represents a shift in the delivered
// value at a level which can be understood by the business and customers.
type Milestone struct {
	// The Title provides a brief name for this milestone.
	Title string `json:"title"`

	// The Description is a markdown formatted body of text which provides additional
	// information about the milestone.
	Description string `json:"description,omitempty"`

	// The Deliverables are concrete tasks which combine to achieve a given milestone.
	Deliverables []*Deliverable `json:"deliverables,omitempty"`
}

// ID will return a deterministic identifier for this milestone which is based on its title and deliverables.
func (m *Milestone) ID() string {
	dids := make([]interface{}, len(m.Deliverables))
	for i, d := range m.Deliverables {
		dids[i] = d.ID()
	}

	return id(m.Title, dids...)
}

func (m *Milestone) ensureDefaults() {
	for _, d := range m.Deliverables {
		d.ensureDefaults()
	}
}

// An Objective describes a high level goal for the team. It is usually something that
// will be worked towards over several milestones and might not have a clear definition of done.
type Objective struct {
	// The Title provides a brief name for this objective.
	Title string `json:"title"`

	// The Description is a markdown formatted body of text that which provides additional
	// context on the objective.
	Description string `json:"description,omitempty"`
}

// ID returns a deterministic identifier for this deliverable which is based on its title and reference
func (o *Objective) ID() string {
	return id(o.Title)
}

// A Deliverable is a concrete task which may be delivered as part of a broader milestone.
// Deliverables can state their RFC2119 requirement level to indicate how their delivery
// impacts the milestone.
type Deliverable struct {
	// The Title provides a brief name for this deliverable.
	Title string `json:"title"`

	// The Description is a markdown formatted body of text which provides additional
	// context on this deliverable.
	Description string `json:"description,omitempty"`

	// The Reference is a URL at which additional information about the deliverable can
	// be found. This may be documentation, a tracking ticket or a PR.
	Reference string `json:"reference,omitempty"`

	// The Requirement specifies, using RFC2119 form, the impact that delivery has on
	// the milestone (MUST, SHOULD, MAY).
	Requirement string `json:"requirement,omitempty"`

	// The State of the deliverable may be one of TODO, DOING, DONE, SKIP.
	State string `json:"state,omitempty"`
}

// ID returns a deterministic identifier for this deliverable which is based on its title and reference
func (d *Deliverable) ID() string {
	return id(d.Title, d.Reference)
}

func (d *Deliverable) ensureDefaults() {
	if d.Requirement == "" {
		d.Requirement = "SHOULD"
	}

	if d.State == "" {
		d.State = "TODO"
	}
}

// The id function generates a deterministic identifier based on the provided seeds.
func id(key string, hash ...interface{}) string {
	prefix := string(regexp.MustCompile("[^a-zA-Z0-9_]").ReplaceAll([]byte(key), []byte("_")))

	h := sha256.New()
	h.Write([]byte(fmt.Sprint(hash...)))
	suffix := fmt.Sprintf("%x", h.Sum(nil))

	return prefix + "_" + suffix[:8]
}
