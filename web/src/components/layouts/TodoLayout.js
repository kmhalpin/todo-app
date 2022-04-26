import TodoHeader from "../header/TodoHeader"
import { Outlet } from "react-router-dom"

const TodoLayout = () => {
  return <>
    <TodoHeader />
    <Outlet />
  </>
}

export default TodoLayout;
