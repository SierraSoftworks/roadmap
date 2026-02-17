using System.Collections.Generic;

namespace Roadmap.DotFX.Models
{
    public class Milestone
    {
        public string Title { get; set; }

        public string Description { get; set; }

        public string Reference { get; set; }

        public IEnumerable<Deliverable> Deliverables { get; set; }
    }
}