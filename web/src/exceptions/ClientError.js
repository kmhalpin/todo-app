class ClientError extends Error {
  constructor(message, statusCode = ClientError.code) {
    super(message);

    if (this.constructor.name === 'ClientError') {
      throw new Error('cannot instantiate abstract class');
    }

    this.statusCode = statusCode;
    this.name = 'ClientError';
  }

  static code = 400;
}

export default ClientError;