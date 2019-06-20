import getPluginInstancesRec from './getPluginInstances'

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

it('should return plugins as it is', () => {
  expect(getPluginInstancesRec([plugin1, plugin2])).toEqual([plugin1, plugin2])
})

it('should return plugins and nested plugins', () => {
  const plugins = [
    {
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
          value: [plugin1, plugin2],
        },
      },
    },
  ]

  expect(getPluginInstancesRec(plugins)).toEqual([
    {
      ...plugins[0],
      props: {
        plugins: {
          ...plugins[0].props.plugins,
          value: ['plugin-1#0', 'plugin-2#0'],
        },
      },
    },
    ...plugins[0].props.plugins.value,
  ])
})

it('should return plugins and nested plugins on multiple levels', () => {
  const plugins = [
    {
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
          value: [
            plugin1,
            {
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
                  value: [plugin2],
                },
              },
            },
          ],
        },
      },
    },
  ]

  expect(getPluginInstancesRec(plugins)).toEqual([
    {
      ...plugins[0],
      props: {
        plugins: {
          ...plugins[0].props.plugins,
          value: ['plugin-1#0', 'container-1#1'],
        },
      },
    },
    plugin1,
    {
      ...plugins[0].props.plugins.value[1],
      props: {
        plugins: {
          ...plugins[0].props.plugins.value[1].props.plugins,
          value: ['plugin-2#0'],
        },
      },
    },
    plugin2,
  ])
})
