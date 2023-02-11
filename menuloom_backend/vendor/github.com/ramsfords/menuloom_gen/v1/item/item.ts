/* eslint-disable */

export const protobufPackage = "item";

export enum servingTime {
  ALLDAY = 0,
  MORNING = 1,
  LUNCH = 2,
  DINNER = 3,
  BREAKFAST = 4,
  BRUNCH = 5,
  UNRECOGNIZED = -1,
}

export interface image {
  /** @gotags: dynamodbav:"assetId" */
  assetId: string;
  /** @gotags: dynamodbav:"publicId" */
  publicId: string;
  /** @gotags: dynamodbav:"url" */
  url: string;
  /** @gotags: dynamodbav:"secureUrl" */
  secureUrl: string;
}

export interface categories {
  /** @gotags: dynamodbav:"categories" */
  categories: category[];
}

export interface category {
  /** @gotags: dynamodbav:"id" */
  id: string;
  /** @gotags: dynamodbav:"name" */
  name: string;
  /** @gotags: dynamodbav:"localName" */
  localName: string;
  /** @gotags: dynamodbav:"description" */
  description: string;
  /** @gotags: dynamodbav:"images" */
  images: image[];
  /** @gotags: dynamodbav:"servingTime" */
  servingTime: servingTime;
  /** @gotags: dynamodbav:"restaurantId" */
  restaurantId: string;
  /** @gotags: dynamodbav:"type" */
  type: string;
  /** @gotags: dynamodbav:"pk" */
  pk: string;
  /** @gotags: dynamodbav:"sk" */
  sk: string;
  /** @gotags: dynamodbav:"rank" */
  rank: number;
  /** @gotags: dynamodbav:"items" */
  items: item[];
}

export interface items {
  /** @gotags: dynamodbav:"items" */
  items: item[];
}

export interface item {
  /** @gotags: dynamodbav:"id" */
  id: string;
  /** @gotags: dynamodbav:"name" */
  name: string;
  /** @gotags: dynamodbav:"localName" */
  localName: string;
  /** @gotags: dynamodbav:"description" */
  description: string;
  /** @gotags: dynamodbav:"price" */
  price: string;
  /** @gotags: dynamodbav:"images" */
  images: image[];
  /** @gotags: dynamodbav:"spiceLevel" */
  spiceLevel: string;
  /** @gotags: dynamodbav:"isAvailable" */
  isAvailable: boolean;
  /** @gotags: dynamodbav:"cookingTime" */
  cookingTime: string;
  /** @gotags: dynamodbav:"reviews" */
  reviews: string[];
  /** @gotags: dynamodbav:"restaurantId" */
  restaurantId: string;
  /** @gotags: dynamodbav:"type" */
  type: string;
  /** @gotags: dynamodbav:"pk" */
  pk: string;
  /** @gotags: dynamodbav:"sk" */
  sk: string;
  /**
   * @gotags: dynamodbav:"categories"
   * it holds names of categories it belongs to
   */
  categories: string[];
  /** @gotags: dynamodbav:"tags" */
  tags: string[];
  /** @gotags: dynamodbav:"isVeg" */
  isVeg: boolean;
  /** @gotags: dynamodbav:"isVegan" */
  isVegan: boolean;
  /** @gotags: dynamodbav:"isGlutenFree" */
  isGlutenFree: boolean;
  /** @gotags: dynamodbav:"isDairyFree" */
  isDairyFree: boolean;
  /** @gotags: dynamodbav:"isNutFree" */
  isNutFree: boolean;
  /** @gotags: dynamodbav:"isEggFree" */
  isEggFree: boolean;
  /** @gotags: dynamodbav:"isSoyFree" */
  isSoyFree: boolean;
}

export interface itemResponse {
  /** @gotags: dynamodbav:"success" */
  success: boolean;
  /** @gotags: dynamodbav:"message" */
  message: string;
}
