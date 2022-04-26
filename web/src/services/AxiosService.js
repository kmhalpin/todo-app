import axios from "axios";
import InvariantError from "../exceptions/InvariantError"
import NotFoundError from "../exceptions/NotFoundError"
import ForbiddenError from "../exceptions/ForbiddenError"
import UnauthorizedError from "../exceptions/UnauthorizedError"

const AxiosService = axios.create({
  baseURL: '/api',
});

export const resInterceptor = AxiosService.interceptors.response.use(
  res => res,
  err => {
    if (axios.isAxiosError(err)) {
      switch (err.response.status) {
        case InvariantError.code:
          throw new InvariantError(err.response.data.message,
            Object.entries(err.response.data.errors)
              .map(([name, error]) => ({ name, errors: [error] })),
          )
        case NotFoundError.code:
          throw new NotFoundError(err.response.data.message)
        case ForbiddenError.code:
          throw new ForbiddenError(err.response.data.message)
        case UnauthorizedError.code:
          throw new UnauthorizedError(err.response.data.message)
        default:
      }
    }
    throw err;
  }
);

export default AxiosService;
