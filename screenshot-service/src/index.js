const puppeteer = require('puppeteer')
const [, , frontendURL, mediaID] = process.argv

async function main() {
  const browser = await puppeteer.launch({ args: ['--no-sandbox'] })
  const page = await browser.newPage()

  return takeScreenshot(page, `http://${frontendURL}/?mediaID=${mediaID}&bot=true`)
    .catch(error => {
      process.exitCode = 1
      console.error(error)
    })
    .then(() => browser.close())
}

async function takeScreenshot(page, url) {
  await page.setViewport({ width: 1920, height: 1080 })

  console.error(`Loading media URL : ${url}`)
  await page.goto(url, { waitUntil: 'networkidle', networkIdleTimeout: 2000 })

  const time = new Date().getTime()
  const fileName = `${time}.png`
  const path = `screenshots/${fileName}`

  console.error(`Taking screenshot : ${path}`)
  await page.screenshot({ path })
  console.log(fileName)
}

async function wait(timeout) {
  return new Promise(resolve => setTimeout(resolve, timeout * 1000))
}

main()
