import { Controller, Get, HttpStatus, UseGuards } from '@nestjs/common';
import { ProductService } from './product.service';
import HandleErrorException from '../../utils/errorHandling.utils';
import { JwtAuthGuard } from '../../guards/Jwt/jwt-auth.guard';
import { ResponseFormat } from 'src/utils/response.utils';

@Controller('product')
export class ProductController {
  constructor(private readonly productService: ProductService) {}

  @Get()
  async getAllProduct() {
    try {
      const res = await this.productService.getProducts();
      return ResponseFormat(res, HttpStatus.OK, 'product data');
    } catch (error) {
      throw HandleErrorException(error);
    }
  }

  @UseGuards(JwtAuthGuard)
  @Get('suggestion')
  async getSuggestedProduct() {
    try {
      const res = await this.productService.getSuggestedProduct();
      return ResponseFormat(res, HttpStatus.OK, 'suggested product data');
    } catch (error) {
      throw HandleErrorException(error);
    }
  }
}
