import ClientError from "./ClientError";

class InvariantError extends ClientError {
  constructor(message, errors = []) {
    super(message);
    this.name = 'InvariantError';
    this.errors = errors;
  }
}

export default InvariantError;