import { PageHeader, Statistic, Row, Skeleton, Typography, Button, DatePicker, Modal } from "antd";
import moment from "moment";
import { Link } from "react-router-dom";
import InvariantError from "../../exceptions/InvariantError";
import UnauthorizedError from "../../exceptions/UnauthorizedError";
import useApp from "../../hooks/useApp";
import { usePromise } from "../../hooks/usePromise";
import { DeleteTodoById, PutTodo } from "../../services/TodoService";
import { getDateString } from "../../utils";

const TodoHeader = () => {
  const { todo, accessToken, func } = useApp();

  const { fetch: fetchUpdate } = usePromise(
    () => PutTodo(accessToken, todo.id, todo),
    {
      manual: true,
      onFailed: (err) => {
        if (err instanceof InvariantError) {
          if (err.message.startsWith("parsing time"))
            err.errors.push({
              name: 'time',
              errors: ['required'],
            })
          Modal.error({
            content: err.errors.map((v) => <p>{v.name} : {v.errors[0]}</p>)
          })
        }
        if (err instanceof UnauthorizedError) {
          func.signOut()
        }
      }
    }
  );

  const { fetch: fetchDelete } = usePromise(
    () => DeleteTodoById(accessToken, todo.id),
    {
      manual: true,
      onSuccess: () => func.navigate('/'),
      onFailed: (err) => {
        if (err instanceof UnauthorizedError) {
          func.signOut()
        }
      }
    }
  );

  return (
    <PageHeader
      title={!todo.loading
        ? (!todo.edit
          ? getDateString(new Date(todo.time))
          : <DatePicker showTime value={moment(todo.time)}
            onChange={(date) => {
              todo.setState({ ...todo, time: date ? date.toISOString() : todo.time })
            }} />
        ) : <Skeleton.Input active />}
      onBack={() => window.history.back()}
      extra={[
        <Button key="1" type="primary" {...todo.edit && {
          onClick: fetchUpdate
        }}>{
            !todo.edit
              ? <Link to={`/todo/${todo.id}/edit`}>Ubah</Link>
              : 'Simpan'
          }</Button>,
        <Button onClick={fetchDelete} key="2" type="default">Hapus</Button>
      ]}
    >
      <Row>
        <Statistic title='Judul'
          valueRender={() => !todo.loading ? <Typography.Title
            editable={todo.edit && { onChange(title) { todo.setState({ ...todo, title }) } }}
            level={4}
          >{todo.title}</Typography.Title> : <Skeleton.Input active />}
        />
      </Row>
    </PageHeader>
  )
}

export default TodoHeader;