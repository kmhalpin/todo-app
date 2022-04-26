import ClientError from "./ClientError";

class ForbiddenError extends ClientError {
  constructor(message) {
    super(message, ForbiddenError.code);
    this.name = 'ForbiddenError';
  }

  static code = 403;
}

export default ForbiddenError;