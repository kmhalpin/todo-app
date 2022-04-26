import { useContext, useEffect } from "react"
import { useNavigate } from "react-router-dom";
import AppContext from "../providers/app/AppContext"
import { useMounted } from "./useMounted";

const useApp = (title, todo = null, setTodo = (todo) => null) => {
  const navigate = useNavigate();
  const { func, ...ctx } = useContext(AppContext);
  const isMounted = useMounted();
  useEffect(() => {
    title && func.setTitle(title)
  }, [func, title]);
  useEffect(() => {
    todo && func.setTodo(todo)
  }, [func, todo])
  useEffect(() => {
    setTodo && func.setSetTodo(() => setTodo)
  }, [func, setTodo])
  return {
    ...ctx, isMounted, func: {
      ...func,
      navigate,
      signOut: () => {
        func.signOut();
        navigate('/login');
      }
    }
  }
}

export default useApp;
