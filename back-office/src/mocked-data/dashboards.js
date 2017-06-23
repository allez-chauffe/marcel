//@flow

export default {
  selectedPlugin: null,
  selectedDashboard: null,
  dashboards: {
    dashboard1: {
      id: 'dashboard1',
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
    dashboard2: {
      id: 'dashboard2',
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
    dashboard3: {
      id: 'dashboard3',
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
    dashboard4: {
      id: 'dashboard4',
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
    dashboard5: {
      id: 'dashboard5',
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
    dashboard6: {
      id: 'dashboard6',
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
    dashboard7: {
      id: 'dashboard7',
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
  },
}
