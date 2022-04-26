import { Skeleton } from "antd";
import { createContext, useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { usePromise } from "../../hooks/usePromise";
import { Login } from "../../services/AuthService";

const AppContext = createContext({
  todo: {
    id: "",
    title: "",
    time: "",
    edit: false,
    setState: (todo) => null,
    loading: true,
  },
  title: "",
  accessToken: "",
  func: {
    setTitle: (title) => null,
    setTodo: (todo) => null,
    setTodoLoading: (loading) => null,
    setSetTodo: (func) => (todo) => null,
    signIn: ({ username, password }, onSuccess = (res) => { }, onFailed = (err) => { }) => { },
    signOut: () => { }
  },
});

export const withAppContext = Component => props => {
  const [title, setTitle] = useState("");
  const [todo, setTodo] = useState({
    id: "",
    title: "",
    time: "",
    edit: false,
  });
  const [todoLoading, setTodoLoading] = useState(true);
  const [setTodoFunc, setSetTodo] = useState(() => todo.setState);
  const [accessToken, setAccessToken] = useState(localStorage.getItem("accessToken"));

  const { fetch: fetchSignIn } = usePromise(
    ({ username, password, onFailed, onSuccess }) => Login(username, password)
      .then(res => ({ res, onSuccess }))
      .catch(onFailed),
    {
      manual: true,
      onSuccess: ({ res, onSuccess }) => {
        setAccessToken(res.data.data.accessToken);
        localStorage.setItem("accessToken", res.data.data.accessToken);
        onSuccess();
      },
    }
  );

  const signIn = ({ username, password }, onSuccess = (res) => { }, onFailed = (err) => { }) =>
    fetchSignIn({ username, password, onFailed, onSuccess });

  const signOut = () => { setAccessToken(""); localStorage.setItem("accessToken", null); }

  return (
    <AppContext.Provider value={{
      title,
      todo: {
        ...todo,
        setState: setTodoFunc,
        loading: todoLoading,
      },
      accessToken,
      func: {
        setTitle,
        setTodo,
        setTodoLoading,
        setSetTodo,
        signIn,
        signOut,
      },
    }}>
      <Component {...props} />
    </AppContext.Provider>
  )
}

export const withAuthenticated = Component => props => {
  const navigate = useNavigate();
  const { accessToken } = useContext(AppContext);
  useEffect(() => {
    !accessToken && navigate("/login")
  });
  return accessToken ? <Component {...props} /> : <Skeleton active />
}

export default AppContext;
