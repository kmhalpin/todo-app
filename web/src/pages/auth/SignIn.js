import { Col, Row, Button, Form, Input } from "antd";
import { useState } from "react";
import InvariantError from "../../exceptions/InvariantError";
import NotFoundError from "../../exceptions/NotFoundError";
import UnauthorizedError from "../../exceptions/UnauthorizedError";
import useApp from "../../hooks/useApp";

const SignIn = () => {
  const { func } = useApp('Sign In');

  const [signIn, setSignIn] = useState({
    username: '',
    password: '',
  });

  const [form] = Form.useForm();

  const signInUser = (signIn) => func.signIn({
    username: signIn.username,
    password: signIn.password,
  },
    () => func.navigate('/'),
    (err) => {
      if (err instanceof InvariantError) {
        form.setFields(err.errors);
      } else if (err instanceof NotFoundError) {
        form.setFields([
          {
            name: 'username',
            errors: ['pengguna tidak ditemukan']
          }
        ]);
      } else if (err instanceof UnauthorizedError) {
        form.setFields([
          {
            name: 'password',
            errors: ['password salah']
          }
        ]);
      }
    }
  );

  return (
    <Row justify='center'>
      <Col>
        <Form form={form} name="signIn" layout='vertical'>
          <Form.Item
            label="Username"
            name="username"
            initialValue={signIn.username}
            rules={[
              { required: false },
            ]}
          >
            <Input
              onChange={(e) => setSignIn({ ...signIn, username: e.target.value })}
            />
          </Form.Item>
          <Form.Item
            label="Password"
            name="password"
            initialValue={signIn.password}
            rules={[
              { required: false },
            ]}
          >
            <Input.Password
              onChange={(e) => setSignIn({ ...signIn, password: e.target.value })}
            />
          </Form.Item>
          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              onClick={() => signInUser(signIn)}
            >
              Sign In
            </Button>
          </Form.Item>
        </Form>
      </Col>
    </Row>
  );
}

export default SignIn;