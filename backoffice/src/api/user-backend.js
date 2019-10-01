import store from '../store'
import fetcher from './fetcher'

const { get, post, put, del } = fetcher(() => store.getState().config.apiURI + 'auth/')

const userBackend = {
  getAllUsers: () => get('users/').then(res => res.json()),

  addUser: user => post('users/', user).then(res => res.json()),

  updateUser: user => put(`users/${user.id}`, user),

  deleteUser: id => del(`users/${id}`),
}

export default userBackend
