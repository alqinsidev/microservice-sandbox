"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ResponseFormat = void 0;
const common_1 = require("@nestjs/common");
const ResponseFormat = (data = null, status = common_1.HttpStatus.OK, message = undefined) => ({
    data,
    status,
    message,
});
exports.ResponseFormat = ResponseFormat;
//# sourceMappingURL=response.utils.js.map