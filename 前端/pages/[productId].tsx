// pages/[productId].tsx

import { useRouter } from "next/router";
import ProductDetail from "../components/ProductDetail";

const ProductPage = () => {
  const router = useRouter();
  const { productId } = router.query;

  if (!productId) {
    // 可以显示加载中的状态或其他错误提示
    return (
      <div className="product-display w-full mt-40 p-7 mx-auto rounded-lg border glass flex items-center justify-center">
        <span className="loading loading-bars loading-lg"></span>
        <span className="loading loading-bars loading-lg"></span>
        <span className="loading loading-bars loading-lg"></span>
      </div>
    );
  }

  return <ProductDetail productId={productId as string} />;
};

export default ProductPage;
