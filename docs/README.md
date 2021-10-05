---
home: true

actions:
    - text: Editor
      link: /editor/
    - text: Get Started
      link: /guide/
      type: secondary

features:
    - title: Structured
      details: |
        Describe your team and project road-maps in a logical, easy to understand, format which can then
        be transformed into diagrams, documents and work items.

    - title: Accessible
      details: |
        Planning can be challenging enough on its own, the tooling you use to manage it shouldn't make that
        any harder and we've tried hard to live up to that.

    - title: Extensible
      details: |
        Road map's schema is designed to be easy to parse and interpret, making the process of adding your own
        output formats or using the data to drive your workflows as easy as pie.
---


Road map is designed to make visualizing and tweaking your road maps as easy
and straightforward as possible. There are a number of good options out there,
including PowerPoint, Visio, LucidChart, Draw.io, GraphViz and just plain old
lists - however none of them have made me particularly excited about drawing
up our team's road maps and many of them make future edits difficult.

The goal here is to change that, by defining your road map in everyone's least
loved markup language: YAML. An example road map might look something like the
following.

## Example

```yaml
title: Example Road Map
description: |
    This is an example of what a road map might look like.

authors:
  - name: Benjamin Pannell
    email: contact@sierrasoftworks.com

timeline:
  - date: 2021-04-21
    title: Project Start
    description: This is when we will start working on the project, get the team ready!

objectives:
  - title: Market Dominance
    description: |
      Provide actionable analytics drawn from big data to improve our brand identity in
      the advertainment sector, maximizing clickbait and establishing ourselves as a disruptor
      in this industry.

milestones:
  - id: team
    title: Build the Team
    description: We don't yet have anyone, that's not gonna work...
    deliverables:
      - title: Team Lead
        state: TODO
        requirement: MUST
        description: This person needs to know enough about this domain to be able to run with the project.

      - title: Senior Engineer 1
      - title: Intern 1..50
        description: This should be cheaper than hiring a proper team (right?).

  - id: done
    title: Finish the Project
    description: We don't need other milestones, do we?
    dependencies:
      - team
    deliverables:
      - title: MVP
        description: Who needs a polished product? Let's just ship the MVP and call it done.
      - title: Marketing
      - title: VC Funding
      - title: Yacht
        reference: https://lmgtfy.app/?q=yacht&t=i
```


<ClientOnly>
    <Contributors repo="SierraSoftworks/roadmap" />
</ClientOnly>