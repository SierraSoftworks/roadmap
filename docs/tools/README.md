# Overview
Road map is, at its heart, a standardized data structure. The tools built to consume it are
where the real value comes from.

## Visualizations
One of the best things to do with your road map is show it to people and what better way than
through a colourful diagram? These tools are designed to make understanding your road map so
easy that your team will wonder what all the fuss was about.

#### [GraphViz](/tools/graphviz/README.md)
Converts your `roadmap.yml` file into a GraphViz DOT format which can be easily rendered into
SVG, PNG and a range of other formats.

```sh
go run github.com/SierraSoftworks/roadmap/tools/graphviz --in roadmap.yml
```

## Documentation
Sometimes you want to be able to convert your road map into a human readable document. This might
be so that you can display it on a website, drop it into a planning system or simply to make it a
bit easier to read. These tools are designed to help with that.

## Ticketing
Once you've finalized your road map, you will probably want to execute on it. Manually entering
road map items into your team's ticketing system is enough to drive even the most hardened planner
crazy, so why not automate that by using one of our ingestion tools?