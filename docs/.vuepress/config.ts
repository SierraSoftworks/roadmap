import { defineUserConfig, PageHeader } from 'vuepress-vite'
import defaultTheme from '@vuepress/theme-default'
import { getDirname, path } from '@vuepress/utils'

const __dirname = getDirname(import.meta.url)

import {googleAnalyticsPlugin} from "@vuepress/plugin-google-analytics"
import {registerComponentsPlugin} from "@vuepress/plugin-register-components"

function htmlDecode(input: string): string {
  return input.replace("&#39;", "'").replace("&amp;", "&").replace("&quot;", '"')
}

function fixPageHeader(header: PageHeader) {
  header.title = htmlDecode(header.title)
  header.children.forEach(child => fixPageHeader(child))
}

export default defineUserConfig({
  lang: 'en-GB',
  title: 'Road map',
  description: "Generate your team's road-maps from structured data.",

  head: [
    ['meta', { name: "description", content: "Road map is a tool which allows you to generate your team's road-maps from structured data." }],
    ['link', { rel: 'icon', href: '/favicon.ico' }]
  ],

  extendsPage(page, app) {
    const fixedHeaders = page.headers || []
    fixedHeaders.forEach(header => fixPageHeader(header))

    page.headers = fixedHeaders;
  },

  theme: defaultTheme({
    logo: 'https://cdn.sierrasoftworks.com/logos/icon.png',
    logoDark: 'https://cdn.sierrasoftworks.com/logos/icon_light.png',

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
  }),

  plugins: [
    googleAnalyticsPlugin({ id: "G-R57T3LCFD4" }),
    registerComponentsPlugin({
      componentsDir: path.resolve(__dirname, './components'),
    })
  ]
})