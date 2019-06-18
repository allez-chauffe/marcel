import { range } from 'lodash'

export default [
  {
    name: `Container 1`,
    eltName: `container-${1}`,
    icon: 'picture_in_picture_alt',
    props: {
      plugins: {
        name: 'plugins',
        description: 'The list of plugins',
        type: 'pluginList',
        value: [],
      },
    },
  },
  {
    name: 'marcel-plugin-hello',
    eltName: 'marcel-plugin-hello',
    icon: 'picture_in_picture_alt',
    props: {
      title: {
        name: 'title',
        description: 'The title to be displayed',
        type: 'string',
        value: 'World',
      },
    },
  },
  {
    name: 'DevFest Lille',
    eltName: 'devfest',
    icon: 'picture_in_picture_alt',
    props: {
      prop1: {
        name: 'prop1',
        description: '',
        type: 'string',
        value: 'value 1',
      },
      prop2: {
        name: 'prop2',
        description: '',
        type: 'string',
        value: 'value2',
      },
    },
  },
  ...range(20).map(i => ({
    name: `Plugin ${i}`,
    eltName: `plugin-${i}`,
    icon: 'picture_in_picture_alt',
    props: {
      prop1: {
        name: 'prop1',
        description: 'some description',
        type: 'string',
        value: 'hello world !',
      },
      prop2: {
        name: 'prop2',
        description: 'some description',
        type: 'number',
        value: 42,
      },
      prop3: {
        name: 'prop3',
        description: 'some description',
        type: 'boolean',
        value: true,
      },
    },
  })),
]
