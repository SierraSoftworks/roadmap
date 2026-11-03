using System.Collections.Generic;

namespace Roadmap.DotFX.Models
{
    public class Roadmap
    {
        public string Title { get; set; }

        public string Description { get; set; }

        public IEnumerable<Author> Authors { get; set; }

        public IEnumerable<TimelineEntry> Timeline { get; set; }

        public IEnumerable<Objective> Objectives { get; set; }

        public IEnumerable<Milestone> Milestones { get; set; }

        // The config block is used by the CLI tools and is accepted here so that road maps
        // which include it can still be rendered by DocFX.
        public Dictionary<string, Dictionary<string, object>> Config { get; set; }
    }
}