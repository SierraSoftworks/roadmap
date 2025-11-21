"""Setup script for mkdocs-roadmap plugin."""

from setuptools import setup, find_packages

with open("README.md", "r", encoding="utf-8") as fh:
    long_description = fh.read()

setup(
    name="mkdocs-roadmap",
    version="0.1.0",
    author="SierraSoftworks",
    description="MkDocs plugin for rendering roadmap.yml files to roadmap.md files",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/SierraSoftworks/roadmap",
    packages=find_packages(),
    include_package_data=True,
    package_data={
        'mkdocs_roadmap': ['roadmap.advanced.md.j2'],
    },
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Intended Audience :: Developers",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
        "Programming Language :: Python :: 3.11",
        "Programming Language :: Python :: 3.12",
        "Programming Language :: Python :: 3.13",
        "Programming Language :: Python :: 3.14",
    ],
    python_requires=">=3.7",
    # Dependencies are defined in pyproject.toml to avoid duplication
    entry_points={
        "mkdocs.plugins": [
            "roadmap = mkdocs_roadmap:RoadmapPlugin",
        ],
    },
)
