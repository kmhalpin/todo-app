import { Skeleton } from "antd";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import NotFoundError from "../../exceptions/NotFoundError";
import UnauthorizedError from "../../exceptions/UnauthorizedError";
import useApp from "../../hooks/useApp";
import { usePromise } from "../../hooks/usePromise";
import { GetTodoById } from "../../services/TodoService";
import TextEditor from "./TextEditor";

const DetailTodo = ({ edit = false }) => {
  const { id } = useParams();
  const [state, setState] = useState({
    id,
    title: "",
    time: "",
    note: "",
    edit,
  });
  const { accessToken, func } = useApp("", state, setState);

  const { loading } = usePromise(
    () => GetTodoById(accessToken, id),
    {
      onSuccess(res) { setState({ ...state, ...res.data.data, edit }) },
      onFailed: (err) => {
        if (err instanceof NotFoundError) {
          func.navigate('/')
        }
        if (err instanceof UnauthorizedError) {
          func.signOut()
        }
      }
    },
    edit,
  );

  useEffect(() => func.setTodoLoading(loading), [loading, func]);

  useEffect(() => setState(prev => ({ ...prev, edit })), [edit]);

  return <div className="content">
    {!loading
      ? <div>
        <TextEditor readonly={!edit} value={state.note} onChange={(note) => { setState({ ...state, note }) }} />
      </div>
      : <Skeleton active />
    }
  </div>
}

export default DetailTodo;