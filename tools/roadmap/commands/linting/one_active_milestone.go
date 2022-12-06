package linting

import (
	"fmt"

	"github.com/SierraSoftworks/roadmap"
)

func init() {
	addRule(RuleFromFunction(func(rm *roadmap.Roadmap) []Issue {
		active := map[int]bool{}

		for i, m := range rm.Milestones {
			for _, d := range m.Deliverables {
				if d.State == "DOING" {
					active[i] = true
				}
			}
		}

		if len(active) > 2 {
			return []Issue{}
		}

		return []Issue{
			{
				Level:   LEVEL_WARNING,
				Message: fmt.Sprintf("You have %d milestones with ongoing tasks (out of a total of %d) which probably means you're using them to track workstreams instead of progress towards your overarching goals.", len(active), len(rm.Milestones)),
				Rule:    "milestones::one_active",
			},
		}
	}))
}
