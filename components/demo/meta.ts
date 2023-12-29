import { IMeta } from '@byted-apaas/designer-react';

const meta: IMeta = {
  props: [
    {
      name: 'key1',
      type: 'String',
      defaultValue: '字符串',
    },
    {
      name: 'onClick',
      type: 'Event',
      title: '点击时'
    }
  ]
};

export default meta;