# Plugins

A plugin is a widget that can be added to a marcel media.
Basically it is just a web site served by marcel and displayed in an iframe.

**:warning: Some parts of this README might be outdated**

# Create a plugin

## Setup

1. Create a directory for the plugin

```bash
mkdir marcel-plugin-<plugin name> && cd marcel-plugin-<plugin name>
```

2. Create a directory name `frontend`

```bash
mkdir frontend && cd frontend
```

3. Create a node module. Don't forget to give a name to the module, the default will be `frontend`.

```bash
# with npm
npm init
# with yarn
yarn init
```

4. Add `marcel-plugin` as a dependency (`npm i marcel-plugin` or `yarn add marcel-plugin`)

```bash
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

```javascript
class MyPlugin extends Marcel.Plugin {
    <...>
}
```

This class is the representaton of the plugin's runtime. Some methods might be implemented in order to declare what the plugin does.

### The `render` method

```javascript
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

```javascript
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
For example, if you need to fetch some data, you should make you API calls here.

This function will be called after the `render()` function.

The current properties are accessible in `this.props` (like in the `render` method), and the previous properties are accessible in the `prevProps` argument.
`prevProps` might be used to avoid expensive API calls if the properties didn't change.

### Initialize the plugin

Finally, you should initialize your plugin:

```javascript
Marcel.init(MyPlugin)
```

## Testing the plugin

To test your plugin outside marcel, you can juste serve the `frontend` directory with any http web server.

```bash
npm install --global serve
serve -s frontend
```

You can then simulate the receipt of props by adding this line at then end of your code

```javascript
Marcel.changeProps({ props1: 'new value' })
```

You can call this function any times you want to try out your plugin.

_Tips: You can use it in the devtools_

# `marcel-plugin` API

To reiceive your props and comunicate with the marcel media, you need to use standard HTML message API. To avoid the pain of implementing marcel's comunication protocol in each plugin, the `marcel-plugin` gives a basic API.

## `Marcel.Plugin`

The main feature of `marcel-plugin` is the `Marcel.Plugin` base class.

By extending this class, you get notified of props updates and you can use a basic state management.

```javascript
class MyPlugin extends Marcel.Plugin {
  constructor() {
    // Don't forget to call the super constructor
    super({
      // You can pass an option object to give the state and props default values
      defaultProps: { prop1: 'default value' },
      defaultState: { state1: 'default state' },
    })

    // You can initialise your plugin here
    // for example, you can store the dom elements depending of props and state value
    this.container = document.getElementById('container')
    this.p = document.query('#container p')
  }

  render = () => {
    // You should do here every UI update needed to keep the props and state synchronised with the DOM
    // It is called a first time with default props and then each time props or state changes

    const { stylesvar, prop1 } = this.props
    const { state1 } = this.state

    this.container.style.color = stylesvar['primary-color']
    this.p.innerText = `${state1} - ${prop1}`
  }

  pluginInitialized = () => {
    // You should do here every side effects that should only be done once
    // but can't be done in constructor because it depends on props.
    // It is called just after the first render with loaded props
    // WARNING: You probably want to use propsChanged in most cases since a side
    //          effect depending on props should probably be rerun on props change
  }

  propsChanged = async (prevProps, prevState) => {
    // You should do here every side effects that depends on props
    // It is called after render on each props or state change
    // You can check if the props or state you depend on actually changed
    // by comparing this.props and this.state with prevProps and prevState

    if (prevProps.prop1 !== this.props.prop1) {
      this.setState({ state1: await fetchSomeData(prop1) })
    }
  }
}
```

The render method will be called to let you synchronise your UI with updated props and state.

You can use the state to keep track of some value and have your plugin automatically rerender on change. To update the state, simply call `this.setState(newState)`. If `newState` is an object, it will be merged into the current state. If it is a function, it will be called with the current state and merge the returned object into the current state.

# Advises and tips for your plugin

Since your plugin will be contained in a iframe that will take only a part of the full webpage, it is highly recommended to remove the border, padding and, very important, the scrollbar.
You can do this with this tiny css snippet:

```css
body {
  margin: 0;
  padding: 0;
  overflow: hidden;
  width: 100vw;
  height: 100vh;
}
```

The user of your project can also give your plugin whatever space they want.
Due to this, you may have a lot of space, or a very tiny space, have a ratio of 1:1, 4:1, 5:3, ...
For this reason, you should try to make your plugin as responsive as possible.
