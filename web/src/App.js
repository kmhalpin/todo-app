import { withAppContext } from "./providers/app/AppContext";
import Layout from "./components/layouts/Layout";
import TodoLayout from "./components/layouts/TodoLayout";

import { Route, Routes, BrowserRouter } from "react-router-dom";
import { CreateTodo, DetailTodo, Home, SignIn } from "./pages";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route path="" element={<Home />} />
          <Route path="create" element={<CreateTodo />} />
          <Route path="login" element={<SignIn />} />
        </Route>
        <Route path="/" element={<TodoLayout />}>
          <Route path="todo/:id" element={<DetailTodo />} />
          <Route path="todo/:id/edit" element={<DetailTodo edit />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default withAppContext(App);
