import ClientError from "./ClientError";

class NotFoundError extends ClientError {
  constructor(message) {
    super(message, NotFoundError.code);
    this.name = 'NotFoundError';
  }

  static code = 404;
}

export default NotFoundError;