# MkDocs
[MkDocs](https://www.mkdocs.org/) is a fast, simple and downright gorgeous static site generator
that's geared towards building project documentation. It is commonly used to host documentation for
teams and projects, making it the perfect platform to showcase your team's road map. Of course, needing
to manually render your road map into Markdown and then update the file isn't the best workflow, so
we've built an MkDocs plugin which will render your `roadmap.yml` file into beautiful documentation
without any extra work on your part.

## Usage
### Prerequisites
You'll need MkDocs installed on your machine and a prepared project. The best way to do this is by
following their [Getting Started Guide](https://www.mkdocs.org/getting-started/).

### Installation
The simplest way to install the plugin is using pip.

::: code-tabs

@tab pip

```sh
pip install mkdocs-roadmap
```

@tab from source

```sh
cd tools/mkdocs-roadmap
pip install .
```

:::

### Configuration
Once you have installed the plugin, you will need to configure MkDocs to use it. This is done by adding
the plugin to your `mkdocs.yml` configuration file and specifying which `roadmap.yml` files you'd like
to render. The plugin will automatically generate corresponding Markdown files during the build process.

```yaml{2-5}
# mkdocs.yml
plugins:
  - roadmap:
      roadmaps:
        - roadmap.yml
```

The `roadmaps` configuration option is a list of paths to `roadmap.yml` files relative to your `docs_dir`.
For each roadmap file, the plugin will:

1. Read the `roadmap.yml` file
2. Render it using a Jinja2 template based on the advanced roadmap format
3. Generate a corresponding `roadmap.md` file in the same location

### Creating a Road map
Once you have installed and configured the plugin, you can render a road map by placing a `roadmap.yml`
file anywhere within your MkDocs `docs_dir`. The plugin will convert this into a `roadmap.md` file which
you can then reference in your `nav` configuration or link to from other pages.

For example, if you have a `roadmap.yml` file at `docs/roadmap.yml`, you can add it to your navigation:

```yaml
nav:
  - Home: index.md
  - Roadmap: roadmap.md
```

::: tip
The plugin generates the Markdown files during the MkDocs build process, so you don't need to manually
create or maintain the `.md` files. Simply update your `roadmap.yml` files and rebuild your site.
:::

## Output
The plugin generates a Markdown file which leverages your existing MkDocs theme. The content is rendered
using Jinja2 templates that produce beautiful, styled roadmap documentation with support for timeline
entries, objectives, milestones, and deliverables. The generated Markdown is compatible with all standard
MkDocs themes.
