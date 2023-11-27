// pages/api/products/[productId].ts

import { NextApiRequest, NextApiResponse } from 'next';

interface ProductInfo {
  id: string;
  image: string;
  name: string;
  description: string;
  story: string;
  publicKey: string;
  condition: string;
  price: number;
}

const products: Record<string, ProductInfo> = {
  '1': {
    id: '1',
    image: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSCEZWpmjwyJXHTLVYBkX7kwYHRXw1N5hYzJr4WxGacT9PmRIcfTo7H8DnDNdnbMAFsBWo&usqp=CAU',
    name: '商品1',
    description: '商品1的描述',
    story: '商品1的故事',
    publicKey: '用户1的公钥地址',
    condition: '良好',
    price: 19.99,
  },
  // 添加更多商品信息...
};

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<ProductInfo | { error: string }>
) {
  const { productId } = req.query;

  try {
    const productInfo = products[productId as string];

    if (productInfo) {
      res.status(200).json(productInfo);
    } else {
      res.status(404).json({ error: 'Product not found' });
    }
  } catch (error) {
    console.error('Error fetching product information:', error);
    res.status(500).json({ error: 'Internal Server Error' });
  }
}
