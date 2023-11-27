import React from 'react';

interface ProductItemProps {
  Pic: string;
  CommodityName: string;
  ExpectPrice: number;
}

const ProductItem: React.FC<ProductItemProps> = ({ Pic, CommodityName, ExpectPrice }) => (
  <div className="product-item p-4 rounded-lg border hover:shadow-lg transition-all duration-300 glass">
    <img src={Pic} alt={CommodityName} className="w-full h-32 object-cover mb-2 rounded-md" />
    <div className="text-sm font-bold mb-1">{CommodityName}</div>
    <div className="text-xs text-gray-600">{ExpectPrice}元</div>
  </div>
);

export default ProductItem;
