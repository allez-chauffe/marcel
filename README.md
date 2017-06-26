# M.A.R.C.E.L. Project

## Smart Mirror (Miroir A Reflet Connecté Et Ludique)

MARCEL is a connected mirror that allows you to add plugins to display data about anything. You can for example add a weather forecast or a twitter plugin and see what the weather will be and the last tweet on a subject.

The main objective of the project is to be plug and play so anybody can add it on his own mirror and develop their own plugins.

## Setup

To populate every plugins, you have to clone the repository with :

```shell
git clone --recursive
```

or after you have cloned the repository :

```shell
git submodule update
```

## Content

Each directory has a specific usage.

* back-office - Manage the plugins and their configuration
* backend - Store the plugins and offer an API to request them
* frontend - Web application that load the plugins
* plugins - Links to repositories of available plugins

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details