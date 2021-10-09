namespace Roadmap.DotFX.Models
{
    public enum Requirement
    {
        MUST,
        SHOULD,
        MAY
    }

    public enum State
    {
        TODO,
        DOING,
        DONE,
        SKIP
    }

    public class Deliverable
    {
        public string Title { get; set; }

        public string Description { get; set; }

        public string Reference { get; set; }

        public Requirement Requirement { get; set; }
        
        public State State { get; set; }
    }
}