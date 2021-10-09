using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.IO;
using System.Text;
using System.Web;
using System.Xml.Linq;
using HandlebarsDotNet;
using HandlebarsDotNet.IO;
using Markdig;
using Markdig.Renderers;
using Markdig.Renderers.Html.Inlines;
using Markdig.Syntax.Inlines;
using Microsoft.DocAsCode.Common;
using Microsoft.DocAsCode.Plugins;

namespace Roadmap.DotFX
{
    public class RoadmapToHtmlBuildStep : IDocumentBuildStep
    {
        public string Name => nameof(RoadmapToHtmlBuildStep);

        public int BuildOrder => 1;

        public IEnumerable<FileModel> Prebuild(ImmutableList<FileModel> models, IHostService host)
        {
            var markdownPipeline = new MarkdownPipelineBuilder()
                .UseAdvancedExtensions();

            Handlebars.RegisterHelper("markdown", (writer, context, parameters) =>
            {
                var model = parameters.At<FileModel>(1);
                if (model is null)
                {
                    throw new InvalidOperationException($"The _model state is missing from the markdown evaluation context ({context.ToJsonString()})");
                }

                var linkToFiles = new HashSet<string>(model.LinkToFiles);

                var builder = new StringBuilder();
                var textWriter = new StringWriter(builder);
                var renderer = new HtmlRenderer(textWriter);

                renderer.LinkRewriter = link =>
                {
                    string linkFile;
                    string anchor = null;

                    var filePath = (RelativePath)model.File;

                    if (PathUtility.IsRelativePath(link))
                    {
                        var index = link.IndexOf('#');
                        if (index == -1)
                        {
                            linkFile = link;
                        }
                        else if (index == 0)
                        {
                            return link;
                        }
                        else
                        {
                            linkFile = link.Remove(index);
                            anchor = link.Substring(index);
                        }
                        var path = filePath + (RelativePath)linkFile;
                        var file = (string)path.GetPathFromWorkingFolder();

                        linkToFiles.Add(HttpUtility.UrlDecode(file));
                        model.LinkToFiles = linkToFiles.ToImmutableHashSet();
                        return file + anchor;
                    }

                    return link;
                };

                var doc = Markdown.Parse(parameters.At<string>(0));

                renderer.Render(doc);

                writer.WriteSafeString(builder.ToString());
            });

            Handlebars.RegisterHelper("add", (writer, context, parameters) =>
            {
                writer.Write(parameters.At<int>(0) + parameters.At<int>(1));
            });

            Handlebars.Configuration.FormatterProviders.Add(new CustomDateTimeFormatter("yyyy-MM-dd"));

            return models;
        }

        public void Build(FileModel model, IHostService host)
        {
            if (model.Content is Dictionary<string, object> buildState)
            {
                var content = buildState["conceptual"] as Models.Roadmap;

                buildState["conceptual"] = Handlebars.Compile(Templates.GetHtmlTemplate())(new
                {
                    roadmap = content,
                    stylesheet = Templates.GetStylesheet(),
                    _model = model
                });
            }
        }

        public void Postbuild(ImmutableList<FileModel> models, IHostService host)
        {
        }

        public sealed class CustomDateTimeFormatter : IFormatter, IFormatterProvider
        {
            private readonly string _format;

            public CustomDateTimeFormatter(string format) => _format = format;

            public void Format<T>(T value, in EncodedTextWriter writer)
            {
                if (!(value is DateTime dateTime))
                    throw new ArgumentException("supposed to be DateTime");

                writer.Write($"{dateTime.ToString(_format)}");
            }

            public bool TryCreateFormatter(Type type, out IFormatter formatter)
            {
                if (type != typeof(DateTime))
                {
                    formatter = null;
                    return false;
                }

                formatter = this;
                return true;
            }
        }
    }
}