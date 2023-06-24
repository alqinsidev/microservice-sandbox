"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ForbiddenError = exports.BadRequestError = exports.NotFoundError = void 0;
const common_1 = require("@nestjs/common");
class NotFoundError extends Error {
    constructor(message) {
        super(message);
        this.name = 'Not Found Error';
    }
}
exports.NotFoundError = NotFoundError;
class BadRequestError extends Error {
    constructor(message) {
        super(message);
        this.name = 'Bad Request Error';
    }
}
exports.BadRequestError = BadRequestError;
class ForbiddenError extends Error {
    constructor(message) {
        super(message);
        this.name = 'Forbidden Request Error';
    }
}
exports.ForbiddenError = ForbiddenError;
const HandleErrorException = (error) => {
    if (error instanceof NotFoundError) {
        throw new common_1.HttpException(error.message, common_1.HttpStatus.NOT_FOUND);
    }
    else if (error instanceof BadRequestError) {
        throw new common_1.HttpException(error.message, common_1.HttpStatus.BAD_REQUEST);
    }
    else if (error instanceof ForbiddenError) {
        throw new common_1.HttpException(error.message, common_1.HttpStatus.FORBIDDEN);
    }
    else {
        common_1.Logger.error(error);
        throw new common_1.HttpException('internal server error', common_1.HttpStatus.INTERNAL_SERVER_ERROR);
    }
};
exports.default = HandleErrorException;
//# sourceMappingURL=errorHandling.utils.js.map