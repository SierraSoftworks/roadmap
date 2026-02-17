"""
MkDocs plugin for rendering roadmap.yml files to roadmap.md files.
"""

from mkdocs.plugins import BasePlugin
from mkdocs.config import config_options
from mkdocs.config.base import Config
import os
import yaml
import jinja2
import markdown as md_lib
from pathlib import Path
from typing import Dict, Any


def state_color(state: str) -> str:
    """Convert state to color."""
    colors = {
        "TODO": "#aaa",
        "DOING": "#63B2EB",
        "DONE": "#3EAF7C",
        "SKIP": "#F65BD2",
    }
    return colors.get(state, "#aaa")


def requirement_color(requirement: str) -> str:
    """Convert requirement to color."""
    colors = {
        "MUST": "#E06446",
        "SHOULD": "#E0AF2F",
        "MAY": "#3ABDE0",
    }
    return colors.get(requirement, "#888")


def render_markdown(text: str) -> str:
    """Render markdown text to HTML."""
    return md_lib.markdown(text)


def date_format(date_value) -> str:
    """Format date to YYYY-MM-DD format."""
    from datetime import datetime, date

    if isinstance(date_value, (datetime, date)):
        return date_value.strftime("%Y-%m-%d")
    elif isinstance(date_value, str):
        # Try to parse common date formats
        try:
            # Try ISO format first
            dt = datetime.fromisoformat(date_value.replace('Z', '+00:00'))
            return dt.strftime("%Y-%m-%d")
        except ValueError:
            # If that fails, try other common formats
            for fmt in ["%Y-%m-%d", "%Y/%m/%d", "%d/%m/%Y"]:
                try:
                    dt = datetime.strptime(date_value, fmt)
                    return dt.strftime("%Y-%m-%d")
                except ValueError:
                    continue
    return str(date_value)


class RoadmapPlugin(BasePlugin):
    """MkDocs plugin for rendering roadmap.yml files."""

    config_scheme = (
        ('roadmaps', config_options.Type(list, default=[])),
    )

    def __init__(self):
        super().__init__()
        self.template = None

    def on_config(self, config: Config, **kwargs) -> Config:
        """Load the Jinja2 template when config is loaded."""
        template_path = Path(__file__).parent / "roadmap.advanced.md.j2"
        with open(template_path, 'r') as f:
            template_content = f.read()

        # Create Jinja2 environment with custom filters
        env = jinja2.Environment(
            loader=jinja2.DictLoader({'roadmap': template_content}),
            autoescape=False
        )
        env.filters['state_color'] = state_color
        env.filters['requirement_color'] = requirement_color
        env.filters['date_format'] = date_format
        env.filters['markdown'] = render_markdown

        self.template = env.get_template('roadmap')
        return config

    def on_files(self, files, config: Config, **kwargs):
        """Process roadmap.yml files and generate .md files."""
        docs_dir = Path(config['docs_dir'])

        # Get roadmap files from config
        roadmap_files = self.config.get('roadmaps', [])

        for roadmap_file in roadmap_files:
            roadmap_path = docs_dir / roadmap_file

            if not roadmap_path.exists():
                continue

            # Read and parse YAML
            with open(roadmap_path, 'r') as f:
                roadmap_data = yaml.safe_load(f)

            # Handle empty or invalid YAML
            if roadmap_data is None:
                roadmap_data = {}

            # Ensure defaults are set
            roadmap_data.setdefault('authors', [])
            roadmap_data.setdefault('timeline', [])
            roadmap_data.setdefault('objectives', [])
            roadmap_data.setdefault('milestones', [])
            roadmap_data.setdefault('title', 'Roadmap')
            roadmap_data.setdefault('description', '')

            # Set defaults for deliverables
            for milestone in roadmap_data.get('milestones', []):
                for deliverable in milestone.get('deliverables', []):
                    if 'requirement' not in deliverable:
                        deliverable['requirement'] = 'SHOULD'
                    if 'state' not in deliverable:
                        deliverable['state'] = 'TODO'

            # Render template - pass roadmap_data as root context to match Go template structure
            rendered = self.template.render(**roadmap_data)

            # Determine output path (replace .yml/.yaml with .md)
            output_path = roadmap_path.with_suffix('.md')

            # Write output
            with open(output_path, 'w') as f:
                f.write(rendered)

            # Add the generated .md file to MkDocs files if it doesn't exist
            rel_path = str(output_path.relative_to(docs_dir))
            # Normalize path separators for cross-platform compatibility
            rel_path = rel_path.replace('\\', '/')

            # Check if file already exists in files list
            existing_files = [f.src_path for f in files]
            if rel_path not in existing_files:
                from mkdocs.structure.files import File
                file_obj = File(
                    path=rel_path,
                    src_dir=str(docs_dir),
                    dest_dir=config['site_dir'],
                    use_directory_urls=config['use_directory_urls']
                )
                files.append(file_obj)

        return files

