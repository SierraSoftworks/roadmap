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
    }
}