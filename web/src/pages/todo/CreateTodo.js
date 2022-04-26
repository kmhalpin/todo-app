import { DatePicker, Form, Input, Button } from "antd";
import { useState } from "react";
import InvariantError from "../../exceptions/InvariantError";
import UnauthorizedError from "../../exceptions/UnauthorizedError";
import useApp from "../../hooks/useApp";
import { usePromise } from "../../hooks/usePromise";
import { PostTodo } from "../../services/TodoService";

const CreateTodo = () => {
  const { accessToken, func } = useApp('Buat Todo');
  const [state, setState] = useState({
    title: "",
    time: "",
  });

  const [form] = Form.useForm();

  const { loading, fetch } = usePromise(
    () => PostTodo(accessToken, state),
    {
      manual: true,
      onSuccess: (res) => func.navigate(`/todo/${res.data.data.id}/edit`),
      onFailed: (err) => {
        if (err instanceof InvariantError) {
          if (err.message.startsWith("parsing time"))
            err.errors.push({
              name: 'time',
              errors: ['required'],
            })
          form.setFields(err.errors);
        }
        if (err instanceof UnauthorizedError) {
          func.signOut()
        }
      }
    }
  )

  return <Form form={form} layout='vertical'>
    <Input.Group compact>
      <Form.Item
        label='Judul'
        name='title'
        rules={[{ required: false }]}
        initialValue={state.title}
      >
        <Input
          onChange={(e) => setState({ ...state, title: e.target.value })}
        />
      </Form.Item>
      <Form.Item
        label='Waktu'
        name='time'
        rules={[{ required: false }]}
      >
        <DatePicker showTime
          onChange={(date) => {
            setState({ ...state, time: date.toISOString() })
          }}
        />
      </Form.Item>
      <Form.Item label=" ">
        <Button
          disabled={loading}
          type="primary"
          htmlType="submit"
          onClick={() => fetch()}
        >
          Tambah
        </Button>
      </Form.Item>
    </Input.Group>
  </Form>
}

export default CreateTodo;