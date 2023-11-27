// ProductDetail.tsx

import React, { useState, useEffect } from "react";
import { useRouter } from "next/router";
import { Product } from "@/pages";

interface ProductDetailProps {
  productId: string;
}

export interface ProductInfo {
  image: string;
  name: string;
  description: string;
  story: string;
  publicKey: string;
  condition: string;
  price: number;
}

const ProductDetail: React.FC<ProductDetailProps> = ({ productId }) => {
  const router = useRouter();
  const [productInfo, setProductInfo] = useState<Product | null>(null);
  const [privateKey, setPrivateKey] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [isPageLoading, setPageIsLoading] = useState(true);

  useEffect(() => {
    // 根据商品ID向后端请求商品信息
    // 用于从后端获取商品信息
    const fetchProductInfo = async () => {
      try {
        const response = await fetch(`http://81.71.5.116:9090/goodsdetail?commodityid=${productId}`,{
          method:"GET",
          headers: {
            "Content-Type": "application/json",
          },
        });
        const data = await response.json();
        if(data){
          setPageIsLoading(false);
        }
        console.log(data)
        setProductInfo(data);
      } catch (error) {
        console.error("Error fetching product information:", error);
      }
    };

    fetchProductInfo();
  }, [productId]);

  const handleConfirm = async () => {
    setIsLoading(true);


    //检查productInfo是否可用
    if(productInfo){
       //创建一个formdata来上传数据
    const formdata = new FormData();

    //把接口入参添加进formdata里
    formdata.append("price",String(productInfo?.ExpectPrice));
    formdata.append("buyeraddress","80f442c18a3ad4f9944b5274a90add19b35cccd1fdd47ba9b28e91f67a3b19214723f7c9813e71c2e18a18bfc28414ed078934ace6adecd733fbb45f8d322646");
    formdata.append("selleraddress",String(productInfo?.PublicAddress));
    formdata.append("commodityid",String(productInfo?.CommodityID));

    // 在这里调用交易接口，使用 privatekey 参数
    // 交易接口返回 true 表示交易成功
    // console.log(formdata);
    // console.log(productInfo);
    //   // 遍历FormData对象的键值对并打印
    //   for (const pair of formdata.entries()) {
    //     console.log(pair[0] + ": " + pair[1]);
    //   }

    const transactionSuccess = await fetch("http://81.71.5.116:9090/confirmorder", {
      method: "POST",
      body: formdata,
    });
    const data = await transactionSuccess.json();
    console.log(data);
    // const transactionSuccess=true;

    if (data) {
      // 交易成功后执行一些操作，比如显示提示，跳转页面等
      alert("交易已完成");
      const transaction_modal=document.getElementById("transaction_model") as HTMLDialogElement | null;
      if(transaction_modal){
        transaction_modal.close();
      }
      //跳转到buypage
      router.push('/')
    } else {
      // 交易失败时处理
      alert("交易失败，请重试");
    }

    setIsLoading(false);
    }
   
  };
  if (!productInfo) {
    // 加载中的状态，可以显示 loading 动画或其他提示
    return <div>Loading...</div>;
  }

  return (
    <div className="p-10">
      {/* 聊天咨询的modal */}
      <dialog id="chat_model" className="modal modal-bottom sm:modal-middle">
        <div className="modal-box">
          <h3 className="font-bold text-lg">Hello!</h3>
          <p className="py-4">Let's chat with the owner!</p>
          <div className="modal-action">
            <form method="dialog">
              {/* if there is a button in form, it will close the modal */}
              <button className="btn">Close</button>
            </form>
          </div>
        </div>
      </dialog>
      {/* 交易的modal，这里交易就不单独做一个页面了，只是一个modal来确认是否交易，如果用户输入私钥确认交易，则调用交易结果，等待交易确认后，则出现一个toast提示交易已完成，并且页面跳转到市场发现首页 */}
      <dialog id="transaction_model" className="modal">
      <div className="modal-box">
        <form method="dialog">
          {/* if there is a button in form, it will close the modal */}
          <button
            className="btn btn-sm btn-circle btn-ghost absolute right-2 top-2"
            // onClick={onClose}
          >
            ✕
          </button>
        </form>
        <h1 className="text-4xl font-bold">交易确认</h1>
        <p className="py-6">
          您确认希望以{productInfo.ExpectPrice}元购买{productInfo.CommodityName}吗？
        </p>

        {/* 输入私钥 */}
        <input
          type="password"
          placeholder="请输入您的私钥"
          value={privateKey}
          onChange={(e) => setPrivateKey(e.target.value)}
          className="input input-bordered mb-4"
        />

        {/* 确认和取消按钮 */}
        <div className="flex justify-end">
          <button
            className={`btn btn-primary mr-2 ${
              isLoading ? "cursor-not-allowed" : ""
            }`}
            onClick={handleConfirm}
            disabled={isLoading}
          >
            {isLoading ? "交易中..." : "确认交易"}
          </button>
        </div>
      </div>
    </dialog>
    {isPageLoading ? (
        <div className="product-display w-full mt-40 p-7 mx-auto rounded-lg border glass flex items-center justify-center">
          <span className="loading loading-bars loading-lg"></span>
          <span className="loading loading-bars loading-lg"></span>
          <span className="loading loading-bars loading-lg"></span>
        </div>
      ) : (
        <div className="card lg:card-side bg-base-100 shadow-xl m-5 p-5">
        <figure>
          <img src={productInfo.Pic} alt={productInfo.CommodityName} />
        </figure>
        <div className="card-body">
          <h1 className="text-5xl font-bold">{productInfo.CommodityName}</h1>
          <p className="py-6 italic text-sm">{productInfo.CommodityDescription}</p>
          <p className="py-6 break-all">卖家地址：{productInfo.PublicAddress}</p>
          <p className="py-6">商品折损程度：{productInfo.Extent}</p>
          <p className="py-6 break-all">商品故事：{productInfo.StoryDetail}</p>
          <h2 className="card-title">预期价格：{productInfo.ExpectPrice}元</h2>
          <div className="card-actions justify-end">
            <button
              className="btn btn-primary"
              onClick={() => {
                const chatModel = document.getElementById("chat_model") as HTMLDialogElement | null;
                if (chatModel) {
                  chatModel.showModal();
                }
              }}
            >
              联系卖家
            </button>
            <button className="btn btn-primary" onClick={()=>
              {
                const transaction_modal=document.getElementById("transaction_model") as HTMLDialogElement | null;
                if(transaction_modal){
                  transaction_modal.showModal();
                }
              }
            }>I WANT IT!</button>
          </div>
        </div>
      </div>
      )}

    </div>
  );
};

export default ProductDetail;
