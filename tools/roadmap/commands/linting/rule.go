package linting

import "github.com/SierraSoftworks/roadmap"

var rules = []Rule{}

func addRule(rule Rule) {
	rules = append(rules, rule)
}

type Rule interface {
	Validate(*roadmap.Roadmap) []Issue
}

type rule struct {
	fn func(*roadmap.Roadmap) []Issue
}

func (r *rule) Validate(rm *roadmap.Roadmap) []Issue {
	return r.fn(rm)
}

/// RuleFromFunction creates a new rule implementation based on the provided function.
func RuleFromFunction(fn func(*roadmap.Roadmap) []Issue) Rule {
	return &rule{fn}
}

/// Validate will run all of the registered rules against the provided road map and return
/// a list of issues that were found.
func Validate(rm *roadmap.Roadmap) []Issue {
	var issues []Issue

	for _, rule := range rules {
		issues = append(issues, rule.Validate(rm)...)
	}

	return issues
}
