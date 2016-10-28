# MARCEL



## First
```bash
npm install
```

## Configuration

- Rename ApiService file (remove exemple.)
- Replace api key and url
  - [Maps]: https://console.developers.google.com/apis
  - [Openweathermap]: http://openweathermap.org/


## Launch the project
```bash
ng serve
```

## Creating a build
```bash
ng build
```

## Custom / Create ...
```bash
ng generate component test
```
Read Angular CLI documentation :)
[Angular CLI](https://github.com/angular/angular-cli)

### Base tag handling in index.html

When building you can modify base tag (`<base href="/">`) in your index.html with `--base-href your-url` option.

```bash
# Sets base tag href to /myUrl/ in your index.html
ng build --base-href /myUrl/
ng build --bh /myUrl/
```

### Running unit tests

```bash
ng test
```

Tests will execute after a build is executed via [Karma](http://karma-runner.github.io/0.13/index.html), and it will automatically watch your files for changes. You can run tests a single time via `--watch=false`.

### Running end-to-end tests

```bash
ng e2e
```

Before running the tests make sure you are serving the app via `ng serve`.

End-to-end tests are run via [Protractor](https://angular.github.io/protractor/).

<!--### Electron app
```terminal
$ npm run electron
```
-->

### Widgets

- [x] Date time
- [x] Weather
- [x] Xee Car
- [x] Bike Vlille
- [x] Home
- [ ] Twitter
- [ ] ...

## Hardware

- [Raspberry Pi]: http://amzn.to/28Q1ztX
- [Motion Sensor]: http://amzn.to/28Q1zdA
- [Relay]: http://amzn.to/28SjqEU
- Monitor
- [Mirror]: http://amzn.to/28PN0bd
