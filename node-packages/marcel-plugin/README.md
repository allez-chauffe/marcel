# Plugins

A plugin is a widget that can be added to a marcel media.
Basically it is just a web site served by marcel and displayed in an iframe.

**:warning: Some parts of this README might be outdated**

# Create a plugin

## Setup

1. Create a directory for the plugin

```sh
mkdir marcel-plugin-<plugin name> && cd marcel-plugin-<plugin name>
```

2. Create a directory name `frontend`

```sh
mkdir frontend && cd frontend
```

3. Create a node module. Don't forget to give a name to the module, the default will be `frontend`.

```sh
# with npm
npm init
# with yarn
yarn init
```

4. Add `marcel-plugin` as a dependency (`npm i marcel-plugin` or `yarn add marcel-plugin`)

```sh
# with npm
npm install marcel-plugin
# with yarn
yarn add marcel-plugin
```

5. Create a file `index.html`

The `frontend` directory will be served by marcel as a standard web site.
Hence the html, js and other static file must be put in this directory.

The plugin will be displayed in an iframe with `index.html` as `src`.

A plugin example is available [here](./example/simple)
[marcel-plugin-text](https://github.com/EmrysMyrddin/marcel-plugin-text) is an example of a very simple plugin.

## Structure of the frontend

The following script must be added in `index.html` in order to the interact with marcel.

```html
<script src="./nodes_modules/marcel-plugin/dist/index.js"></script>
```

Then, in a another script, a class that inherits `Marcel.Plugin` must be created:

```js
class MyPlugin extends Marcel.Plugin {
    <...>
}
```

This class is the representaton of the plugin's runtime. Some methods might be implemented in order to declare what the plugin does.

### The `render` method

```js
class MyPlugin extends Marcel.Plugin {
  render() {
    const { props1, props2 } = this.props
  }
}
```

This function is called every time marcel renders the plugin.
As its name implies, this function is responsible for rendering the plugin.
The object `this.props` will contain every properties entered by the user in the backoffice.

### The `propsDidChange` method:

```js
class MyPlugin extends Marcel.Plugin {
  propsDidChange(prevProps) {
    const { prop1, prop2 } = this.props
    if (prop1 !== prevProps.prop1) {
      this.updateMyPluginData()
    }
  }
}
```

This function should be used for any sides effect dependending on the props.
For example, making API calls in order to fetch some data.

This function will be called after the `render()` function.

The current properties are accessible in `this.props` (like in the `render` method), and the previous properties are accessible in the `prevProps` argument.
`prevProps` might be used to avoid expensive API calls if the properties didn't change.

### Initialize the plugin

Finally, a plugin must be initialized:

```js
Marcel.init(MyPlugin)
```

## Testing the plugin

A plugin might be tested outside marcel by serving the `frontend` directory with any http web server:

```sh
npm install --global serve
serve -s frontend
```

Then new props can be simulated by calling:

```js
Marcel.changeProps({ props1: 'new value' })
```

_Tips: This can either be done in the devtools or by adding it at the end of the plugin's script._

# `marcel-plugin` API

To receive the props and communicate with the marcel media, the standard HTML message API must be uesed.
A basic API built over the message API is provided by `marcel-plugin`.

## `Marcel.Plugin`

The main feature of `marcel-plugin` is the `Marcel.Plugin` base class.

By extending this class, the plugin is notified of props updates and might implement a basic state management.

```js
class MyPlugin extends Marcel.Plugin {
  constructor() {
    // The super constructor must always be called
    super({
      // Optional default props and default state
      defaultProps: { prop1: 'default value' },
      defaultState: { state1: 'default state' },
    })

    // Some initialization can be performed here
    // For example, storing dom elements
    this.container = document.getElementById('container')
    this.p = document.query('#container p')
  }

  render = () => {
    // Any UI updates needed to keep the DOM synchronised with the props and state
    // It is called a first time with default props and state and then each time props or state change

    const { stylesvar, prop1 } = this.props
    const { state1 } = this.state

    this.container.style.color = stylesvar['primary-color']
    this.p.innerText = `${state1} - ${prop1}`
  }

  pluginInitialized = () => {
    // Every side effects that should only be done once
    // but can't be done in the constructor because it depends on props goes here
    // It is called just after the first render with loaded props
    // WARNING: Most of the time propsChanged should be used
  }

  propsChanged = async (prevProps, prevState) => {
    // Every side effects depending on props goes here
    // It is called after render on each props or state change
    // Actual changes can be detected by comparing this.props and this.state with prevProps and prevState

    if (prevProps.prop1 !== this.props.prop1) {
      this.setState({ state1: await fetchSomeData(prop1) })
    }
  }
}
```

The render method will be called to let the UI synchronize with updated props and state.

The state can be used to keep track of some values and have automatically rerender the plugin on changes.
To update the state, `this.setState(newState)` must be called.
If `newState` is an object, it will be merged into the current state. If it is a function, it will be called with the current state and must return the new state.

# Tips and tricks

Since a plugin will be contained in a iframe that will take only a part of the full webpage, it is highly recommended to remove the border, padding and, must importantly, the scrollbar.
This can be done with this tiny css snippet:

```css
body {
  margin: 0;
  padding: 0;
  overflow: hidden;
  width: 100vw;
  height: 100vh;
}
```

The space reserved to the plugin might change a lot, from a lot of space to a very tiny space, with ratio of 1:1, 4:1, 5:3...
For this reason, a plugin must be as responsive as possible.
