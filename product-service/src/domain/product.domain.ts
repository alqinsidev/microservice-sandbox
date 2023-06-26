export interface Product {
  id: string;
  createdAt: string;
  price: string;
  department: string;
  product: string;
}

export interface ProductIdr extends Product {
  price_idr: string;
}

export interface SortedProduct {
  mostExpensive: ProductIdr[];
  mostCheap: ProductIdr[];
}

export interface CurrencyResponse {
  meta: {
    last_updated_at: string;
  };
  data: {
    [currencyCode: string]: Currency;
  };
}

interface Currency {
  code: string;
  value: number;
}
