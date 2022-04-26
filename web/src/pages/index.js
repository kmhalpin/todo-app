import { withAuthenticated } from "../providers/app/AppContext";
import SignIn from "./auth/SignIn";
import home from "./home/Home";
import createTodo from "./todo/CreateTodo";
import detailTodo from "./todo/DetailTodo";

export {
  SignIn,
};
export const Home = withAuthenticated(home);
export const CreateTodo = withAuthenticated(createTodo);
export const DetailTodo = withAuthenticated(detailTodo);