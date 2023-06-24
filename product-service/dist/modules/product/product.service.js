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
exports.ProductService = void 0;
const common_1 = require("@nestjs/common");
const axios_1 = require("axios");
let ProductService = class ProductService {
    constructor() {
        this.apiUrl =
            'https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product';
    }
    async getProducts() {
        try {
            const response = await axios_1.default.get(this.apiUrl);
            const { data } = response;
            const conversionRate = await getConversionRate();
            const formattedData = data.map((product) => (Object.assign(Object.assign({}, product), { price_idr: convertUsdToIdr(product.price, conversionRate) })));
            return formattedData;
        }
        catch (error) {
            throw error;
        }
    }
    async getSuggestedProduct() {
        try {
            const rawData = await this.getProducts();
            const { mostExpensive, mostCheap } = sortProductsByPrice(rawData);
            return [...mostExpensive, ...mostCheap];
        }
        catch (error) {
            throw error;
        }
    }
};
ProductService = __decorate([
    (0, common_1.Injectable)(),
    __metadata("design:paramtypes", [])
], ProductService);
exports.ProductService = ProductService;
const sortProductsByPrice = (products) => {
    products.sort((a, b) => parseFloat(a.price) - parseFloat(b.price));
    const mostExpensive = products.slice(-Math.min(5, products.length));
    const mostCheap = products.slice(0, Math.min(5, products.length));
    return { mostExpensive, mostCheap };
};
const getConversionRate = async () => {
    try {
        throw 'error';
    }
    catch (error) {
        return 15000;
    }
};
const convertUsdToIdr = (usd, conversionRate) => {
    return (Number(usd) * conversionRate).toString();
};
//# sourceMappingURL=product.service.js.map