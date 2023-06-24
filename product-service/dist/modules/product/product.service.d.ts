import { ProductIdr } from 'src/domain/product.domain';
export declare class ProductService {
    private readonly apiUrl;
    constructor();
    getProducts(): Promise<ProductIdr[]>;
    getSuggestedProduct(): Promise<ProductIdr[]>;
}
