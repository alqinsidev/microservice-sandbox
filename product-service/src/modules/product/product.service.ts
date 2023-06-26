import { Injectable } from '@nestjs/common';
import axios from 'axios';
import {
  CurrencyResponse,
  Product,
  ProductIdr,
  SortedProduct,
} from 'src/domain/product.domain';

@Injectable()
export class ProductService {
  private readonly productUrl: string;
  private readonly currencyUrl: string;
  constructor() {
    this.productUrl =
      'https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product';
    this.currencyUrl =
      'https://api.currencyapi.com/v3/latest?apikey=aoypan2ESZYfDUJa42KknRTRgM8TjkwJyS1EFqrE&currencies=IDR';
  }

  async getProducts() {
    try {
      const response = await axios.get(this.productUrl);
      const { data } = response;
      const conversionRate = await getConversionRate(this.currencyUrl);
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

const getConversionRate = async (url: string) => {
  try {
    const response = await axios.get(url);
    const { data }: { data: CurrencyResponse } = response;
    const conversionRate = Number(data.data.IDR.value);

    return conversionRate;
  } catch (error) {
    return 15000;
  }
};
const convertUsdToIdr = (usd: string, conversionRate) => {
  return (Number(usd) * conversionRate).toString();
};
