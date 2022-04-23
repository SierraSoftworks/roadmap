import { defineUserConfig, PageHeader, DefaultThemeOptions } from 'vuepress'
import { path } from '@vuepress/utils'

function htmlDecode(input: string): string {
  return input.replace("&#39;", "'").replace("&amp;", "&").replace("&quot;", '"')
}

function fixPageHeader(header: PageHeader) {
  header.title = htmlDecode(header.title)
  header.children.forEach(child => fixPageHeader(child))
}

export default defineUserConfig<DefaultThemeOptions>({
  lang: 'en-GB',
  title: 'Road map',
  description: "Generate your team's road-maps from structured data.",

  head: [
    ['meta', { name: "description", content: "Road map is a tool which allows you to generate your team's road-maps from structured data." }],
    ['link', { rel: 'icon', href: '/favicon.ico' }]
  ],

  //bundler: "@vuepress/bundler-vite",

  clientAppEnhanceFiles: [
    path.resolve(__dirname, "enhance", "cloudflare.analytics.js")
  ],

  extendsPageData(page, app) {
    const fixedHeaders = page.headers || []
    fixedHeaders.forEach(header => fixPageHeader(header))

    return {
      headers: fixedHeaders,
    }
  },

  themeConfig: {
    logo: 'https://cdn.sierrasoftworks.com/logos/icon.png',

    repo: "SierraSoftworks/roadmap",
    docsDir: 'docs',
    navbar: [
      {
        text: "Getting Started",
        link: "/guide/README.md",
      },
      {
        text: "Tools",
        link: "/tools/README.md",
      },
      {
        text: "Report an Issue",
        link: "https://github.com/SierraSoftworks/roadmap/issues/new",
        target: "_blank"
      }
    ],

    sidebar: {
      '/guide/': [
        {
          text: "Getting Started",
          children: [
            '/guide/README.md',
          ]
        },
        {
          text: "Advanced",
          children: [
            '/guide/advanced/schema.md',
          ]
        }
      ],
      '/tools/': [
        {
          text: 'Tools',
          children: [
            '/tools/README.md',
          ]
        },
        {
          text: 'Visualizations',
          children: [
            '/tools/visualizations/graphviz/README.md',
          ]
        },
        {
          text: 'Documentation',
          children: [
            '/tools/documentation/docfx/README.md',
            '/tools/documentation/html/README.md',
            '/tools/documentation/markdown/README.md',
            '/tools/documentation/web-viewer/README.md',
          ]
        }
      ]
    }
  },

  plugins: [
    ["@vuepress/plugin-google-analytics", { id: "G-R57T3LCFD4" }],
    ["@vuepress/plugin-register-components", {
      componentsDir: path.resolve(__dirname, "./components")
    }],
  ]
})