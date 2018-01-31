//@flow
import store from '../store'
import { User } from '../user'
import fetcher from './fetcher'

const { get, post, put, del } = fetcher(() => store.getState().config.authURI)

const userBackend = {
  getAllUsers: () => get('users/').then(res => res.json()),

  addUser: (user: User) => post('users/', user).then(res => res.json()),
  
  updateUser: (user: User) => put(`users/${user.id}`, user),

  deleteUser: (id: string) => del(`users/${id}`),
}

export default userBackend