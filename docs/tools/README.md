# Overview
Road map is, at its heart, a standardized data structure. The tools built to consume it are
where the real value comes from.

## Editors
Putting together your road map is something we'd like to make as easy as possible. To achieve
that, there are several editors and editor integrations available. Their goal is to provide
you with guidance and (in some cases) realtime feedback on your road map.


## Visualizations
One of the best things to do with your road map is show it to people and what better way than
through a colourful diagram? These tools are designed to make understanding your road map so
easy that your team will wonder what all the fuss was about.

#### [GraphViz](/tools/visualizations/graphviz/README.md)
Converts your `roadmap.yml` file into a GraphViz DOT format which can be easily rendered into
SVG, PNG and a range of other formats.

```sh
go install github.com/SierraSoftworks/roadmap/tools/roadmap@latest
roadmap render graphviz --in roadmap.yml
```

## Documentation
Sometimes you want to be able to convert your road map into a human readable document. This might
be so that you can display it on a website, drop it into a planning system or simply to make it a
bit easier to read. These tools are designed to help with that.

#### [DocFX](/tools/documentation/docfx/README.md)
Allows you to render a `roadmap.yml` file directly within DocFX, if that's your poison.

```sh
nuget install DocFX.Plugins.Roadmap -ExcludeVersion -OutputDirectory .
```

#### [HTML](/tools/documentation/html/README.md)
Converts your `roadmap.yml` file into a static HTML file which can be easily viewed in most
web browsers.

```sh
go install github.com/SierraSoftworks/roadmap/tools/roadmap@latest
roadmap render html --in roadmap.yml --out roadmap.html
```

#### [Markdown](/tools/documentation/markdown/README.md)
Converts your `roadmap.yml` file into a Markdown file which can be rendered by most Markdown
rendering programs.

```sh
go install github.com/SierraSoftworks/roadmap/tools/roadmap@latest

# Generate a full Markdown file
roadmap render md --in roadmap.yml --out roadmap.md

# Generate a simplified Markdown file (for GitHub comments/issues etc)
roadmap render md --int roadmap.yml --out roadmap.md --simple
```

#### [MkDocs](/tools/documentation/mkdocs/README.md)
Allows you to render a `roadmap.yml` file directly within MkDocs projects through an automated plugin.

```sh
pip install mkdocs-roadmap
```

#### [Web Viewer](/tools/documentation/web-viewer/README.md)
The Road Map website includes a viewer which allows you to visualize your road map directly
from a public repository. It will load the `roadmap.yml` file (if present) in the root of
your repository and render that in your browser.

```
https://roadmap.sierrasoftworks.com/viewer/github.com#USER/REPO
```

## Ticketing
Once you've finalized your road map, you will probably want to execute on it. Manually entering
road map items into your team's ticketing system is enough to drive even the most hardened planner
crazy, so why not automate that by using one of our ingestion tools?
