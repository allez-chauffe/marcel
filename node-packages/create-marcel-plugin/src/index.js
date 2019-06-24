const fs = require('fs')
const path = require('path')
const execSync = require('child_process').execSync
const toCamelCase = require('to-camel-case')

const [, , eltName] = process.argv

const fatal = (...message) => {
  console.error(...message)
  process.exit(-1)
}

if (!eltName) fatal('A plugin name should be given : $ yarn create marcel-plugin my-plugin')

const [firstChar, ...endOfName] = toCamelCase(eltName)
const name = [firstChar.toUpperCase(), ...endOfName].join('')
const pluginPath = path.resolve(eltName)
const frontendPath = path.join(pluginPath, 'frontend')

if (fs.existsSync(pluginPath)) fatal(`Can't create plugin : ${pluginPath} already exists`)

fs.mkdirSync(pluginPath)
fs.mkdirSync(frontendPath)

const packageDescriptor = `{
  "name": "${eltName}",
  "version": "1.0.0",
  "main": "index.js",
  "license": "MIT",
  "dependencies": {
  }
}
`

const pluginDescriptor = `{
  "name": "${name}",
  "description": "",
  "eltName": "${eltName}",
  "icon": "picture_in_picture_alt",
  "frontend": {
    "props": [
      {
        "name": "firstName",
        "description": "",
        "type": "string",
        "value": "Marcel"
      }
    ]
  }
}
`

const pluginIndex = `<!DOCTYPE html>
<html lang="en">
  <head>
    <title>${eltName}</title>
    <meta charset="utf-8" />
    <script src="./node_modules/marcel-plugin/dist/index.js"></script>
    <style>
      body {
        margin: 0;
        padding: 0;
        overflow: hidden;
        width: 100vw;
        height: 100vh;
      }
    </style>
  </head>

  <body>
    <div id="root"></div>

    <script>
      class ${name} extends Marcel.Plugin {
        constructor() {
          super()
          this.root = document.getElementById('root')
        }

        render() {
          const { firstName, stylesvar = {} } = this.props

          this.root.innerText = "Hello \`\${firstName} !"

          // stylesvar is a special property containing the global media theme.
          // You should use it to have a consistent style accross all the media.
          if (stylesvar['primary-color']) this.root.style.color = stylesvar['primary-color']
          if (stylesvar['font-family']) this.root.style.fontFamily = stylesvar['font-family']
        }
      }

      Marcel.init(${name})

      // uncomment this line to try the plugin in a browser :
      // Marcel.changeProps({ firstName: 'Marcel' })
    </script>
  </body>
</html>
`

fs.writeFileSync(path.join(pluginPath, 'package.json'), packageDescriptor)
fs.writeFileSync(path.join(pluginPath, 'marcel.json'), pluginDescriptor)
fs.writeFileSync(path.join(frontendPath, 'index.html'), pluginIndex)

console.info('Installing dependencies...')
execSync(`cd ${pluginPath} && yarn && cd ..`)

console.info(`The plugin has successfully been generated. You can now go to the ${eltName} folder and begin to make awesome things !`)
