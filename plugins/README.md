# Plugins
A plugin is a widget that can be added to a MARCEL board. 
It can be composed of:
- A frontend, which will be an iframe added to the MARCEL board.
- (optionally) A backend, if the plugin need to run it's own backend.

## Create your own plugin

Creating a MARCEL plugin is simple:
1. Create a new directory named "marcel-plugin-<your plugin name>"
2. Create a directory `frontend` and, if needed, a directory `backend`
3. In the directory `frontend`, create a file `index.html`

Your frontend will be the `index.html` and will be the one displayed in MARCEL.
The javascript can be in the `index.html` or in separates files, but must stay in the `frontend` directory.

You can also have several HTML file if needed, but be aware that since your plugin will be contained in a `<iframe>`, your user won't have access to the navigation history and arrows.

For an example of a simple plugin, you can look at [marcel-plugin-text](https://github.com/EmrysMyrddin/marcel-plugin-text).

### Structure of the frontend / index.html

The first thing you will need is [marcel.js](https://github.com/EmrysMyrddin/marcel-plugin-text/raw/master/frontend/marcel.js), which is a small script that will help you handle the interaction with MARCEL.
Include it in your HTML:
```html
<script src="./marcel.js"></script>
```

Then, you should create a new class that inherit `Marcel.Plugin`:
```javascript
class MyPlugin extends Marcel.plugin {
    <...>
}
```
Your class should implement two function: 

1. The render function:

```javascript
render() {
    const { props1, props2 } = this.props
}
```

This function is called every time MARCEL need to render your plugin. 
As the name of the function imply, you should do everything that involve rendering your page in this function.
The object `this.props` will contain every properties sent by MARCEL.

2. The propsDidChange function:

```javascript
propsDidChange(prevProps) {
    const { props1, props2 } = this.props
    if (props1 !== prevProps.props1) {
        this.updateMyPluginData()
    }
}
```

This function will be called after the `render()` function. 
It will also receive the properties sent by MARCEL in the `this.props` object, but it also has the previous set of properties in the argument `prevProps`.
This can be useful to check if your properties have changed, and only trigger expensive API call if the properties did changed.

Finally, you should create an instance of your class:
```javascript
const instance = new MyPlugin()
```

Additionally, you can add the line:
```javascript
Marcel.Debug.changeProps({props1: "new value"})
```

At the end of your file to force MARCEL to call `render()` and `propsDidChange(prevProps)` when you load the page which is really useful when you are still developing and debugging the plugin.

### Trying out your plugin

Since your plugin is just a normal website, you should be able to test it by just opening the index.html in your browser.
You can also use a development web server like `serve` to autoreload your page while you are working on it.

### Advises and tips for your plugin
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
For this reason, you should try to make your plugin as responsive as you can.