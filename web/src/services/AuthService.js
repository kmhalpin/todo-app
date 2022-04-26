import AxiosService from "./AxiosService"

export const Login = (username, password) => {
  return AxiosService.post('/auth', {
    username, password
  })
}