import ClientError from "./ClientError";

class UnauthorizedError extends ClientError {
  constructor(message) {
    super(message, UnauthorizedError.code);
    this.name = 'UnauthorizedError';
  }

  static code = 401;
}

export default UnauthorizedError;