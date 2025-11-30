# MkDocs Roadmap Plugin

The MkDocs Roadmap plugin allows you to define `roadmap.yml` files in your MkDocs project and automatically generate corresponding `roadmap.md` files during the build process. This plugin uses Jinja2 templates based on the advanced roadmap template to render beautiful, styled roadmap documentation.

## Installation

Install the plugin using pip:

```bash
pip install mkdocs-roadmap
```

Or install from source:

```bash
cd tools/mkdocs-roadmap
pip install .
```

## Usage

### Configuration

Add the plugin to your `mkdocs.yml` configuration file:

```yaml
plugins:
  - roadmap:
      roadmaps:
        - path/to/roadmap.yml
        - another/path/to/roadmap.yml
```

The `roadmaps` configuration option is a list of paths to `roadmap.yml` files relative to your `docs_dir`. The plugin will:

1. Read each `roadmap.yml` file
2. Render it using the Jinja2 template
3. Generate a corresponding `roadmap.md` file in the same location (replacing `.yml` or `.yaml` with `.md`)

### Example

If you have a `roadmap.yml` file at `docs/roadmap.yml`, configure it like this:

```yaml
plugins:
  - roadmap:
      roadmaps:
        - roadmap.yml
```

This will generate `docs/roadmap.md` during the build process, which you can then reference in your `nav` configuration:

```yaml
nav:
  - Home: index.md
  - Roadmap: roadmap.md
```

### Roadmap File Format

The plugin expects roadmap files to follow the roadmap schema. See the main [roadmap documentation](https://roadmap.sierrasoftworks.com/) for details on the schema.

Example `roadmap.yml`:

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

## Development

To develop or modify the plugin:

1. Clone the repository
2. Navigate to `tools/mkdocs-roadmap`
3. Install in development mode: `pip install -e .`
4. Make your changes
5. Test with your MkDocs project
