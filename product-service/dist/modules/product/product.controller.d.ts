import { ProductService } from './product.service';
export declare class ProductController {
    private readonly productService;
    constructor(productService: ProductService);
    getAllProduct(): Promise<{
        data: any;
        status: number;
        message: string;
    }>;
    getSuggestedProduct(): Promise<{
        data: any;
        status: number;
        message: string;
    }>;
}
