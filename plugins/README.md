# Plugins

A plugin is a widget that can be added to a marcel media. 
It can be composed of:
- A frontend, which will be an iframe added to the marcel media.
- (optional) A backend, if the plugin need to run it's own backend.

## :warning: Some parts of this README might be outdated

## Create your own plugin

Creating a marcel plugin is simple:
1. Create a new directory named "marcel-plugin-<your plugin name>"
2. Create a directory `frontend` with a file `index.html`.
3. (Optional) Create a directory `backend` with a docker image tarball.

The frontend directory will be served in HTTP by marcel and it will act as it's own root.
Due to this, you can put all the html, js and other static file that you need in this directory.
By default, the `index.html` is served.

The plugin will be an iframe that display the result of a GET request to `/` on this server.

For an example of a simple plugin, you can look at [marcel-plugin-text](https://github.com/EmrysMyrddin/marcel-plugin-text).

### Structure of the frontend

The first thing you will need is [marcel.js](https://github.com/EmrysMyrddin/marcel-plugin-text/raw/master/frontend/marcel.js), which is a small script that will help you handle the interaction with marcel.
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
Your class can implement two function: 

1. The render function:

```javascript
render() {
    const { props1, props2 } = this.props
}
```

This function is called every time marcel need to render your plugin. 
As the name of the function imply, you should do everything that involve rendering your page in this function.
The object `this.props` will contain every properties sent by marcel.

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
It will also receive the properties sent by marcel in the `this.props` object, but it also has the previous set of properties in the argument `prevProps`.
This can be useful to check if your properties have changed, and only trigger expensive API call if the properties did changed.

Finally, you should create an instance of your class:
```javascript
const instance = new MyPlugin()
```

Additionally, you can add the line:
```javascript
Marcel.Debug.changeProps({ props1: "new value" })
```

At the end of your file to force marcel to call `render()` and `propsDidChange(prevProps)` when you load the page which is really useful when you are still developing and debugging the plugin.

### Trying out your plugin

Since your plugin is just a normal website, you should be able to test it by just opening the index.html in your browser.
You can also use a development web server like `serve` to autoreload your page while you are working on it.

You can also put:
```javascript
Marcel.Debug.changeProps({ props1: "new value" })
```
In the Javascript console to simulate an update from marcel.

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
For this reason, you should try to make your plugin as responsive as possible.