//@flow
import { range } from 'lodash'

export default [
  {
    name: `Container 1`,
    elementName: `container-${1}`,
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
  ...range(20).map(i => ({
    name: `Plugin ${i}`,
    elementName: `plugin-${i}`,
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
