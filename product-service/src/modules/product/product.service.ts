import { Injectable } from '@nestjs/common';
import axios from 'axios';
import { Product, ProductIdr, SortedProduct } from 'src/domain/product.domain';

@Injectable()
export class ProductService {
  private readonly apiUrl: string;
  constructor() {
    this.apiUrl =
      'https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product';
  }

  async getProducts() {
    try {
      const response = await axios.get(this.apiUrl);
      const { data } = response;
      const conversionRate = await getConversionRate();
      const formattedData: ProductIdr[] = data.map((product: Product) => ({
        ...product,
        price_idr: convertUsdToIdr(product.price, conversionRate),
      }));
      return formattedData;
    } catch (error) {
      throw error;
    }
  }

  async getSuggestedProduct() {
    try {
      const rawData = await this.getProducts();
      const { mostExpensive, mostCheap } = sortProductsByPrice(rawData);
      return [...mostExpensive, ...mostCheap];
    } catch (error) {
      throw error;
    }
  }
}

const sortProductsByPrice = (products: ProductIdr[]): SortedProduct => {
  products.sort((a, b) => parseFloat(a.price) - parseFloat(b.price));

  const mostExpensive = products.slice(-Math.min(5, products.length));
  const mostCheap = products.slice(0, Math.min(5, products.length));

  return { mostExpensive, mostCheap };
};

const getConversionRate = async () => {
  try {
    //fetch api
    //return usd_idr value
    throw 'error';
  } catch (error) {
    return 15000;
  }
};
const convertUsdToIdr = (usd: string, conversionRate) => {
  return (Number(usd) * conversionRate).toString();
};
