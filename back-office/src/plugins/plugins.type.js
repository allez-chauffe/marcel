// @flow

export type PropBase = { name: string }

export type Prop =
  | ({ type: 'string', value: string } & PropBase)
  | ({ type: 'number', value: number } & PropBase)
  | ({ type: 'boolean', value: boolean } & PropBase)
  | ({ type: 'json', value: mixed } & PropBase)
export type Plugin = {
  name: string,
  props: Prop[],
  icon: string,
  elementName: string,
}
