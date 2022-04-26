import { Skeleton, Card, Space } from "antd";
import { useState } from "react";
import { Link } from "react-router-dom";
import UnauthorizedError from "../../exceptions/UnauthorizedError";
import useApp from "../../hooks/useApp";
import { usePromise } from "../../hooks/usePromise";
import { GetTodo } from "../../services/TodoService";
import { getDateString } from "../../utils";
import CreateTodo from "../todo/CreateTodo";

const Home = () => {
  const { accessToken, func } = useApp('Home');

  const [todos, setTodos] = useState([]);

  const { loading } = usePromise(
    () => GetTodo(accessToken),
    {
      onSuccess: (res) => { setTodos(res.data.data) },
      onFailed: (err) => {
        if (err instanceof UnauthorizedError) {
          func.signOut()
        }
      }
    }
  );

  return <div className="content">
    <CreateTodo />
    <Space direction='vertical' size='large' style={{ display: 'flex' }}>
      {!loading
        ? todos.map((todo) => <Link key={todo.id} to={`/todo/${todo.id}`}>
          <Card>
            <Card.Meta
              title={todo.title}
              description={
                <p>{getDateString(new Date(todo.time))}</p>
              }
            />
          </Card>
        </Link>)
        : <Skeleton active />
      }
    </Space>
  </div>
}

export default Home;