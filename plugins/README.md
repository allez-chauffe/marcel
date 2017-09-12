# Plugins

## Development

To develop a new plugin, you just have to create a new directory with the name of your plugin and an HTML file that contains a Polymer component. You can see an example in the marcel-plugin-hello plugin which is one of the most basic plugin possible.

If you want to use the global style for example for font-family, you have to import the shared-styles :

```html
<link rel="import" href="/plugins/shared-styles/shared-styles.html" />
```

And then between the two template tags :

```html
 <style include="shared-styles"></style>
```

As the styles, you can also import other plugins to use them. For example in the welcome plugin, if you want to use the datetime plugin, you have to import it:

```html
<link rel="import" href="/plugins/marcel-plugin-datetime" />
```

And use it as a custom element

```html
<datetime-item></datetime-item>
```

## Build

The best way to build your plugin is to use macel-builder.

```shell
npm i -g marcel-builder
cd plugin-front-folder
marcel-builder
```

Il will create an index.html in a new build directory which will contain everything your plugin must have to run.