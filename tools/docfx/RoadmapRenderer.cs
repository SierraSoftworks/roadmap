using System.Collections.Generic;
using System.Collections.Immutable;
using System.Composition;
using System.IO;
using Microsoft.DocAsCode.Common;
using Microsoft.DocAsCode.Plugins;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;

namespace Roadmap.DotFX
{
    [Export(typeof(IDocumentProcessor))]
    public class RoadmapRenderer : IDocumentProcessor
    {
        public string Name => nameof(RoadmapRenderer);

        public IEnumerable<IDocumentBuildStep> BuildSteps => new [] {
            new RoadmapToHtmlBuildStep()
        };

        public ProcessingPriority GetProcessingPriority(FileAndType file)
        {
            if (file.Type == DocumentType.Article && file.File.EndsWith("roadmap.yml"))
            {
                return ProcessingPriority.High;
            }

            return ProcessingPriority.NotSupported;
        }

        public FileModel Load(FileAndType file, ImmutableDictionary<string, object> metadata)
        {
            var deserializer = new DeserializerBuilder()
                .WithNamingConvention(new CamelCaseNamingConvention())
                .Build();

            var roadmap = deserializer.Deserialize<Models.Roadmap>(File.ReadAllText(EnvironmentContext.FileAbstractLayer.GetPhysicalPath(file.File)));
            var localPathFromRoot = PathUtility.MakeRelativePath(EnvironmentContext.BaseDirectory, EnvironmentContext.FileAbstractLayer.GetPhysicalPath(file.File));

            return new FileModel(file, new Dictionary<string, object>{
                ["conceptual"] = roadmap
            })
            {
                LocalPathFromRoot = localPathFromRoot,

            };
        }

        public SaveResult Save(FileModel model)
        {
            var result = new SaveResult
            {
                DocumentType = "Conceptual",
                FileWithoutExtension = Path.ChangeExtension(model.File, null),
                LinkToFiles = model.LinkToFiles.ToImmutableArray(),
                LinkToUids = model.LinkToUids,
                FileLinkSources = model.FileLinkSources,
                UidLinkSources = model.UidLinkSources,
            };

            return result;
        }

        public void UpdateHref(FileModel model, IDocumentBuildContext context)
        {
        }
    }
}