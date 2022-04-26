import AxiosService from "./AxiosService"

export const GetTodo = (token) => {
  return AxiosService.get('/todo', {
    headers: {
      'Authorization': `Bearer ${token}`,
    }
  })
}

export const GetTodoById = (token, id) => {
  return AxiosService.get(`/todo/${id}`, {
    headers: {
      'Authorization': `Bearer ${token}`,
    }
  })
}

export const PostTodo = (token, payload) => {
  return AxiosService.post('/todo', payload, {
    headers: {
      'Authorization': `Bearer ${token}`,
    }
  })
}

export const PutTodo = (token, id, payload) => {
  return AxiosService.put(`/todo/${id}`, payload, {
    headers: {
      'Authorization': `Bearer ${token}`,
    }
  })
}

export const DeleteTodoById = (token, id) => {
  return AxiosService.delete(`/todo/${id}`, {
    headers: {
      'Authorization': `Bearer ${token}`,
    }
  })
}