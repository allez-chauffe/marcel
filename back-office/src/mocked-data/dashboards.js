//@flow

export default {
  selectedPlugin: null,
  selectedDashboard: null,
  dashboards: {
    'Dashboard 1': {
      name: 'Dashboard 1',
      description: 'Some description',
      rows: 20,
      cols: 20,
      ratio: 16 / 9,
      plugins: {
        'plugin-1#0': {
          name: `Plugin 1`,
          elementName: `plugin-1`,
          instanceId: 'plugin-1#0',
          icon: 'picture_in_picture_alt',
          x: 0,
          y: 0,
          columns: 2,
          rows: 3,
          props: {
            prop1: { name: 'prop1', type: 'string', value: 'hello world !' },
            prop2: { name: 'prop2', type: 'number', value: 42 },
            prop3: { name: 'prop3', type: 'boolean', value: true },
          },
        },
      },
    },
    'Dashboard 2': {
      name: 'Dashboard 2',
      description: `
        Some description dlkfsldk flsdflksjk fljsdklflksdjfl ksdjfl
        ksdjlkfsdlkfjklsdflsfsdf sdf sd f sdf sdf sdf sdfsdfdsf dsfdsf
        sfsd f sdfsdfdsfds sdf kjsdkfljsd lkfjlksdj fksdjf lks djfl 
        ksdjlkfj io jioh oih oihoij ijioi upoupo
      `,
      rows: 20,
      cols: 20,
      ratio: 16 / 9,
      plugins: {
        'plugin-1#0': {
          name: `Plugin 1`,
          elementName: `plugin-1`,
          instanceId: 'plugin-1#0',
          icon: 'picture_in_picture_alt',
          x: 0,
          y: 0,
          columns: 2,
          rows: 3,
          props: {
            prop1: { name: 'prop1', type: 'string', value: 'hello world !' },
            prop2: { name: 'prop2', type: 'number', value: 42 },
            prop3: { name: 'prop3', type: 'boolean', value: true },
          },
        },
      },
    },
    'Dashboard 3': {
      name: 'Dashboard 3',
      description: 'Some description',
      rows: 20,
      cols: 20,
      ratio: 16 / 9,
      plugins: {
        'plugin-1#0': {
          name: `Plugin 1`,
          elementName: `plugin-1`,
          instanceId: 'plugin-1#0',
          icon: 'picture_in_picture_alt',
          x: 0,
          y: 0,
          columns: 2,
          rows: 3,
          props: {
            prop1: { name: 'prop1', type: 'string', value: 'hello world !' },
            prop2: { name: 'prop2', type: 'number', value: 42 },
            prop3: { name: 'prop3', type: 'boolean', value: true },
          },
        },
      },
    },
    'Dashboard 4': {
      name: 'Dashboard 4',
      description: 'Some description',
      rows: 20,
      cols: 20,
      ratio: 16 / 9,
      plugins: {
        'plugin-1#0': {
          name: `Plugin 1`,
          elementName: `plugin-1`,
          instanceId: 'plugin-1#0',
          icon: 'picture_in_picture_alt',
          x: 0,
          y: 0,
          columns: 2,
          rows: 3,
          props: {
            prop1: { name: 'prop1', type: 'string', value: 'hello world !' },
            prop2: { name: 'prop2', type: 'number', value: 42 },
            prop3: { name: 'prop3', type: 'boolean', value: true },
          },
        },
      },
    },
    'Dashboard 5': {
      name: 'Dashboard 5',
      description: 'Some description',
      rows: 20,
      cols: 20,
      ratio: 16 / 9,
      plugins: {
        'plugin-1#0': {
          name: `Plugin 1`,
          elementName: `plugin-1`,
          instanceId: 'plugin-1#0',
          icon: 'picture_in_picture_alt',
          x: 0,
          y: 0,
          columns: 2,
          rows: 3,
          props: {
            prop1: { name: 'prop1', type: 'string', value: 'hello world !' },
            prop2: { name: 'prop2', type: 'number', value: 42 },
            prop3: { name: 'prop3', type: 'boolean', value: true },
          },
        },
      },
    },
    'Dashboard 6': {
      name: 'Dashboard 6',
      description: `
        Some description dlkfsldk flsdflksjk fljsdklflksdjfl ksdjfl
        ksdjlkfsdlkfjklsdflsfsdf sdf sd f sdf sdf sdf sdfsdfdsf dsfdsf
        sfsd f sdfsdfdsfds sdf kjsdkfljsd lkfjlksdj fksdjf lks djfl 
        ksdjlkfj io jioh oih oihoij ijioi upoupo
        sdf sdf sdf sdfsdfdsf dsfdsf
        sfsd f sdfsdfdsfds sdf kjsdkfljsd lkfjlksdj fksdjf lks djfl 
        ksdjlkfj io jioh oih oihoij ijioi upoupo
        sdf sdf sdf sdfsdfdsf dsdsf
        sfsd f sdfsdfdsfds sdf kjsdkfljsd lkfjlksdj fksdjf lks djfl 
        ksdjlkfj io jioh oih oihoij ijioi upoupo
      `,
      rows: 20,
      cols: 20,
      ratio: 16 / 9,
      plugins: {
        'plugin-1#0': {
          name: `Plugin 1`,
          elementName: `plugin-1`,
          instanceId: 'plugin-1#0',
          icon: 'picture_in_picture_alt',
          x: 0,
          y: 0,
          columns: 2,
          rows: 3,
          props: {
            prop1: { name: 'prop1', type: 'string', value: 'hello world !' },
            prop2: { name: 'prop2', type: 'number', value: 42 },
            prop3: { name: 'prop3', type: 'boolean', value: true },
          },
        },
      },
    },
    'Dashboard 7': {
      name: 'Dashboard 7',
      description: 'Some description',
      rows: 20,
      cols: 20,
      ratio: 16 / 9,
      plugins: {
        'plugin-1#0': {
          name: `Plugin 1`,
          elementName: `plugin-1`,
          instanceId: 'plugin-1#0',
          icon: 'picture_in_picture_alt',
          x: 0,
          y: 0,
          columns: 2,
          rows: 3,
          props: {
            prop1: { name: 'prop1', type: 'string', value: 'hello world !' },
            prop2: { name: 'prop2', type: 'number', value: 42 },
            prop3: { name: 'prop3', type: 'boolean', value: true },
          },
        },
      },
    },
    'Dashboard 8': {
      name: 'Dashboard 8',
      description: 'Some description',
      rows: 20,
      cols: 20,
      ratio: 16 / 9,
      plugins: {
        'plugin-1#0': {
          name: `Plugin 1`,
          elementName: `plugin-1`,
          instanceId: 'plugin-1#0',
          icon: 'picture_in_picture_alt',
          x: 0,
          y: 0,
          columns: 2,
          rows: 3,
          props: {
            prop1: { name: 'prop1', type: 'string', value: 'hello world !' },
            prop2: { name: 'prop2', type: 'number', value: 42 },
            prop3: { name: 'prop3', type: 'boolean', value: true },
          },
        },
      },
    },
  },
}
