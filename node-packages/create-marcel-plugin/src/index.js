#!/usr/bin/env node

const fs = require('fs')
const path = require('path')
const { execSync } = require('child_process')
const toCamelCase = require('to-camel-case')

const fatal = (...message) => {
  console.error(...message)
  process.exit(-1)
}

const getPluginInfo = eltName => {
  const [firstChar, ...endOfName] = toCamelCase(eltName)
  const name = [firstChar.toUpperCase(), ...endOfName].join('')

  return {
    eltName,
    name,
    path: path.resolve(eltName),
  }
}

const files = [
  {
    path: 'marcel.json',
    content: ({ name, eltName }) => `{
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
  },
  {
    path: 'frontend/package.json',
    content: ({ eltName }) => `{
  "name": "${eltName}",
  "version": "1.0.0",
  "main": "index.js",
  "license": "MIT",
  "dependencies": {
    "marcel-plugin": "^1.0.0"
  }
}
`
  },
  {
    path: 'frontend/index.html',
    content: ({ name, eltName }) => `<!DOCTYPE html>
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
  },
]

const commands = [
  { command: 'yarn', cwd: 'frontend' },
]

if (process.argv.length < 3) fatal('A plugin name should be given : $ yarn create marcel-plugin my-plugin')

const plugin = getPluginInfo(process.argv[2])

// Initialize directory
if (fs.existsSync(plugin.path)) fatal(`Can't create plugin : directory ${plugin.path} already exists`)
fs.mkdirSync(plugin.path)

// Write files
files.forEach((file => {
  const filePath = path.resolve(plugin.path, file.path)
  fs.mkdirSync(path.dirname(filePath), { recursive: true })
  fs.writeFileSync(filePath, file.content(plugin))
}))

// Run commands
commands.forEach(({ command, cwd })=> {
  execSync(command, {
    cwd: cwd ? path.resolve(plugin.path, cwd) : plugin.path,
  })
})

console.info(`The plugin has successfully been generated. You can now go to the ${plugin.eltName} folder and start making awesome things !`)
