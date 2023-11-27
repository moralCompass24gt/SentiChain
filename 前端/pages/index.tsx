import { useEffect, useState } from "react";
import ProductItem from "@/components/productItem";
import Link from "next/link";

//显示声明product数据类型
export interface Product {
  CommodityID: number;
  CommodityName: string;
  CommodityDescription: string;
  ExpectPrice: number;
  Extent: number;
  Pic: string;
  PublicAddress: string;
  StoryDetail: string;
  Time: string;
}

const Buypage = () => {
  const [searchTerm, setSearchTerm] = useState(""); //搜索栏的输入
  const [product, setProduct] = useState<Product[]>([]); //产品列表
  const [isLoading, setIsLoading] = useState(false);
  //把用户输入的需求发送给后端，并获取返回的商品列表
  const handleSearch = async () => {
    try {
      const response = await fetch("http://81.71.5.116:9090/goodsitem", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
      const data = await response.json();
      if(data){
        setIsLoading(true);
      }
      console.log(data.commoditieslist);
      setProduct(data.commoditieslist);
      // setIsLoading(false);
    } catch (error) {
      console.error(error);
      setIsLoading(false);
    }
  };

  useEffect(() => {
    handleSearch();
  }, [searchTerm]);
  return (
    <div
      className="hero min-h-screen"
      style={{
        backgroundImage:
          'url("https://cdn.midjourney.com/1c799e23-c213-4bdb-8464-c01e870c6c00/0_0.webp")',
      }}
    >
      {/* 搜索框 */}
      <div className="join absolute top-7">
        <input
          className="input input-bordered join-item w-full"
          placeholder="商品名称或关键词"
          value={searchTerm}
          onChange={(e) => {
            setSearchTerm(e.target.value);
          }}
        />
        <button className="btn join-item rounded-r-full" onClick={handleSearch}>
          搜索
        </button>
      </div>
      {/* 分割线 */}
      <div className="flex flex-col w-full absolute top-20">
        <div className="divider divider-neutral">
          <span className="font-bold text-3xl">商 品 列 表</span>
        </div>
      </div>
      {/* 商品展示栏 */}
      {!isLoading ? (
        <div className="product-display w-full mt-40 p-7 mx-auto rounded-lg border glass flex items-center justify-center">
          <span className="loading loading-bars loading-lg"></span>
          <span className="loading loading-bars loading-lg"></span>
          <span className="loading loading-bars loading-lg"></span>
        </div>
      ) : (
        <div className="product-display w-full mt-40 p-7 mx-auto rounded-lg border glass">
          <div className="grid grid-cols-4 gap-6">
            {product.map((p, index) => (
              <Link key={index} href={`/[productId]`} as={`/${p.CommodityID}`}>
                <ProductItem key={index} {...p} />
              </Link>
            ))}
          </div>
        </div>
      )}
    </div>
  );
};
export default Buypage;
