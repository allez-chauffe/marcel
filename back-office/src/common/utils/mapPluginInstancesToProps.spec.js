//@flow
import mapPluginInstancesToProps from './mapPluginInstancesToProps'

const plugin1 = {
  name: `Plugin 1`,
  eltName: `plugin-1`,
  instanceId: 'plugin-1#0',
  icon: 'picture_in_picture_alt',
  x: 0,
  y: 0,
  cols: 2,
  rows: 3,
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
}

const plugin2 = {
  name: `Plugin 2`,
  eltName: `plugin-2`,
  instanceId: 'plugin-2#0',
  icon: 'picture_in_picture_alt',
  x: 0,
  y: 0,
  cols: 2,
  rows: 3,
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
}

it('should return the plugin instance', () => {
  const pluginInstances = {
    'plugin-1#0': plugin1,
    'plugin-2#0': plugin2,
  }
  expect(mapPluginInstancesToProps(pluginInstances)('plugin-1#0')).toEqual(
    plugin1,
  )
})

it('should return the plugin instance with plugin instances mapped', () => {
  const container = {
    name: `Container 1`,
    eltName: `container-1`,
    instanceId: 'container-1#0',
    icon: 'picture_in_picture_alt',
    x: 0,
    y: 0,
    cols: 2,
    rows: 3,
    props: {
      plugins: {
        name: 'plugins',
        description: 'The list of plugins',
        type: 'pluginList',
        value: ['plugin-1#0', 'plugin-2#0'],
      },
    },
  }
  const pluginInstances = {
    'container-1#0': container,
    'plugin-1#0': plugin1,
    'plugin-2#0': plugin2,
  }

  expect(
    mapPluginInstancesToProps(pluginInstances)(container.instanceId),
  ).toEqual({
    ...container,
    props: {
      plugins: {
        ...container.props.plugins,
        value: [plugin1, plugin2],
      },
    },
  })
})

it('should return plugin instance and nested plugin instances on multiple levels', () => {
  const container1 = {
    name: `Container 1`,
    eltName: `container-1`,
    instanceId: 'container-1#0',
    icon: 'picture_in_picture_alt',
    x: 0,
    y: 0,
    cols: 2,
    rows: 3,
    props: {
      plugins: {
        name: 'plugins',
        description: 'The list of plugins',
        type: 'pluginList',
        value: ['plugin-1#0', 'container-1#1'],
      },
    },
  }

  const container2 = {
    name: `Container 1`,
    eltName: `container-1`,
    instanceId: 'container-1#1',
    icon: 'picture_in_picture_alt',
    x: 0,
    y: 0,
    cols: 2,
    rows: 3,
    props: {
      plugins: {
        name: 'plugins',
        description: 'The list of plugins',
        type: 'pluginList',
        value: ['plugin-2#0'],
      },
    },
  }
  const pluginInstances = {
    'container-1#0': container1,
    'container-1#1': container2,
    'plugin-1#0': plugin1,
    'plugin-2#0': plugin2,
  }

  expect(
    mapPluginInstancesToProps(pluginInstances)(container1.instanceId),
  ).toEqual({
    ...container1,
    props: {
      plugins: {
        ...container1.props.plugins,
        value: [
          plugin1,
          {
            ...container2,
            props: {
              plugins: {
                ...container2.props.plugins,
                value: [plugin2],
              },
            },
          },
        ],
      },
    },
  })
})
