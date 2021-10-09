using System;
using System.IO;
using System.Reflection;

namespace Roadmap.DotFX
{
    public static class Templates
    {
        public static string GetStylesheet() => GetResourceText("roadmap.css");

        public static string GetHtmlTemplate() => GetResourceText("roadmap.html.handlebars");

        private static string GetResourceText(string filename)
        {
            var assembly = Assembly.GetExecutingAssembly();
            var resourceName = $"{typeof(Templates).Namespace}.{filename}";

            Console.WriteLine(string.Join(",", assembly.GetManifestResourceNames()));
            Console.WriteLine(resourceName);

            using (var stream = assembly.GetManifestResourceStream(resourceName))
            using (var reader = new StreamReader(stream))
            {
                return reader.ReadToEnd();
            }
        }
    }
}