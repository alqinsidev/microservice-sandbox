"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.ProductController = void 0;
const common_1 = require("@nestjs/common");
const product_service_1 = require("./product.service");
const errorHandling_utils_1 = require("../../utils/errorHandling.utils");
const jwt_auth_guard_1 = require("../../guards/Jwt/jwt-auth.guard");
const response_utils_1 = require("../../utils/response.utils");
let ProductController = class ProductController {
    constructor(productService) {
        this.productService = productService;
    }
    async getAllProduct() {
        try {
            const res = await this.productService.getProducts();
            return (0, response_utils_1.ResponseFormat)(res, common_1.HttpStatus.OK, 'product data');
        }
        catch (error) {
            throw (0, errorHandling_utils_1.default)(error);
        }
    }
    async getSuggestedProduct() {
        try {
            const res = await this.productService.getSuggestedProduct();
            return (0, response_utils_1.ResponseFormat)(res, common_1.HttpStatus.OK, 'suggested product data');
        }
        catch (error) {
            throw (0, errorHandling_utils_1.default)(error);
        }
    }
};
__decorate([
    (0, common_1.Get)(),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", Promise)
], ProductController.prototype, "getAllProduct", null);
__decorate([
    (0, common_1.UseGuards)(jwt_auth_guard_1.JwtAuthGuard),
    (0, common_1.Get)('suggestion'),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", Promise)
], ProductController.prototype, "getSuggestedProduct", null);
ProductController = __decorate([
    (0, common_1.Controller)('product'),
    __metadata("design:paramtypes", [product_service_1.ProductService])
], ProductController);
exports.ProductController = ProductController;
//# sourceMappingURL=product.controller.js.map