// @flow
export type PropBase = { name: string, description: string }

export type Prop =
  | ({ type: 'string', value: string } & PropBase)
  | ({ type: 'number', value: number } & PropBase)
  | ({ type: 'boolean', value: boolean } & PropBase)
  | ({ type: 'json', value: mixed } & PropBase)
export type Plugin = {
  name: string,
  props: { [propName: string]: ?Prop },
  icon: string,
  elementName: string,
}

export type State = Plugin[]
