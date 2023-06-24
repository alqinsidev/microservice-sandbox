import { HttpException, HttpStatus, Logger } from '@nestjs/common';

export class NotFoundError extends Error {
  constructor(message: string) {
    super(message);
    this.name = 'Not Found Error';
  }
}

export class BadRequestError extends Error {
  constructor(message: string) {
    super(message);
    this.name = 'Bad Request Error';
  }
}

export class ForbiddenError extends Error {
  constructor(message: string) {
    super(message);
    this.name = 'Forbidden Request Error';
  }
}

const HandleErrorException = (error: any) => {
  if (error instanceof NotFoundError) {
    throw new HttpException(error.message, HttpStatus.NOT_FOUND);
  } else if (error instanceof BadRequestError) {
    throw new HttpException(error.message, HttpStatus.BAD_REQUEST);
  } else if (error instanceof ForbiddenError) {
    throw new HttpException(error.message, HttpStatus.FORBIDDEN);
  } else {
    // ADD SENTRY HERE IF NECESSARY
    Logger.error(error);
    throw new HttpException(
      'internal server error',
      HttpStatus.INTERNAL_SERVER_ERROR,
    );
  }
};
export default HandleErrorException;