// @flow
export type PropBase = { name: string }

//TODO: Refine this..
export type PropTyped = { type: string, value: mixed }
// | { type: 'string', value: string }
// | { type: 'number', value: number }
// | { type: 'boolean', value: boolean }
// | { type: 'json', value: mixed }

export type Prop = PropBase & PropTyped
export type Plugin = {
  name: string,
  props: Prop[],
  icon: string,
  elementName: string,
}
