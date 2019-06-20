# Plugins

A plugin is a widget that can be added to a marcel media.
A plugin is basicly a web site served by marcel and displayed in an iframe.

**:warning: Some parts of this README might be outdated**

# Create your own plugin

## Setup

1. Create a directory for your plugin

```bash
mkdir marcel-plugin-<your plugin name> && cd marcel-plugin-<your plugin name>
```

2. Create a directory name `frontend`

```bash
mkdir frontend && cd frontend
```

3. Create a node module. Don't forget to give a name to your module, the default will be `frontend`.

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
Due to this, you can put all the html, js and other static file that you need in this directory.

The plugin will be an iframe displaying `index.html`.

You can find a plugin example [here](./example/simple)
For an example of a simple plugin, you can look at [marcel-plugin-text](https://github.com/EmrysMyrddin/marcel-plugin-text).

## Structure of the frontend

The first thing you will need is `marcel-plugin` which is a small script that will help you handle the interaction with marcel.
Include it in your HTML:

```html
<script src="./nodes_modules/marcel-plugin/dist/index.js"></script>
```

Then, you should create a new class that inherit `Marcel.Plugin`:

```javascript
class MyPlugin extends Marcel.Plugin {
    <...>
}
```

This class is the representaton of your plugin. Some methods can be implemented to declare what your plugin do.

### The `render` method

```javascript
class MyPlugin extends Marcel.Plugin {
  render() {
    const { props1, props2 } = this.props
  }
}
```

This function is called every time marcel needs to render your plugin.
As the name of the function imply, you should do everything that involve rendering your plugin in this function.
The object `this.props` will contains every properties entered by the user in the backoffice.

### The `propsDidChange` method:

```javascript
class MyPlugin extends Marcel.Plugin {
  propsDidChange(prevProps) {
    const { props1, props2 } = this.props
    if (props1 !== prevProps.props1) {
      this.updateMyPluginData()
    }
  }
}
```

This function should be used to make any sides effect dependending on props.
For example, if you need to fetch some data, you should make you API calls here.

This function will be called after the `render()` function.

You can access the current properties in `this.props` like in the `render` method, but you have also access to previous properties in the `prevProps` argument.
This can be useful only trigger expensive API call if some properties did changed.

### Initialize the plugin

Finally, you should initialize your plugin:

```javascript
Marcel.init(MyPlugin)
```

## Trying out your plugin

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
