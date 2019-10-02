module.exports = [
  {
    path: 'marcel.json',
    content: ({ name, eltName }) => JSON.stringify({
      name: name,
      description: "",
      eltName: eltName,
      icon: "picture_in_picture_alt",
      frontend: {
        props: [
          {
            name: "firstName",
            description: "",
            type: "string",
            value: "Marcel"
          }
        ]
      }
    }, null, 2)
  },
  {
    path: 'frontend/package.json',
    content: ({ eltName }) => JSON.stringify({
      name: eltName,
      version: "1.0.0",
      license: "MIT",
      scripts: {
        start: "serve -s"
      }
    }, null, 2)
  },
  {
    path: 'frontend/.gitignore',
    content: () => 'node_modules\n'
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

          this.root.innerText = \`Hello \${firstName} !\`

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
