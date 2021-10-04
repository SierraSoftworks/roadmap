import { defineClientAppEnhance } from '@vuepress/client'

const token = "b541c01a0ed44af898ecaa688c0b0ddd"

export default defineClientAppEnhance(() => {
    if (__VUEPRESS_DEV__ || __VUEPRESS_SSR__) return

    if (window.cfbeacon) {
        return
    }

    const cfScript = document.createElement("script")
    cfScript.src = "https://static.cloudflareinsights.com/beacon.min.js"
    cfScript.defer = true
    cfScript.setAttribute("data-cf-beacon", JSON.stringify({ token }))

    document.head.appendChild(cfScript)
    window.cfbeacon = true
})