// GoodsRecord.tsx

import React, { useState, useEffect } from "react";
import Link from "next/link";
import { useRouter } from "next/router";


interface GoodsRecordItem {
  CommodityDescription: string;
  PublicAddress: string;
  StoryDetail: string;
  Time: string;
  CommodityID:number;
  CommodityName:string;
  ExpectPrice:number;
  Extent:number;
  IsSold:boolean;
  Pic:string;
}

const userBuyRecord= () => {
  const [records, setRecords] = useState<GoodsRecordItem[]>([]);
  const [isLoading,setIsLoading]=useState(true);
  const router=useRouter();
  const {address}=router.query;

  useEffect(() => {
    // 根据 recordType 调用不同的接口
    const fetchRecords = async () => {
      try {
        const response = await fetch(`http://81.71.5.116:9090/goodsitembuy?address=${address}`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        });
        const data = await response.json();
        if(data){
          setIsLoading(false);
        }
        setRecords(data.commoditieslist);
      } catch (error) {
        console.error(error);
      }
    }
    fetchRecords();
  },[address]);

  return (
    <div className="hero min-h-screen"
    style={{
        backgroundImage:
          "url(https://cdn.midjourney.com/167de30a-30b8-482b-8ee3-b1c513ef4332/0_1.webp)",
      }}>
      <div className="hero-overlay bg-opacity-60"></div>
      <div className="hero-content text-center text-neutral-content">
        <div className="w-full p-5">
          <h1 className="mb-5 text-6xl font-bold font-sans text-white">
          User's Purchased Product Records
          </h1>
          <div
            className="mx-auto rounded-lg border glass flex-col p-5 overflow-x-auto"
            style={{ maxWidth: "80%" }}
          >
            {isLoading?(
              <div className="product-display w-full mt-40 p-7 mx-auto rounded-lg border glass flex items-center justify-center">
              <span className="loading loading-bars loading-lg"></span>
              <span className="loading loading-bars loading-lg"></span>
              <span className="loading loading-bars loading-lg"></span>
            </div>
            ):(
              records.map((record, index) => (
                <Link key={index} href={`/[productId]`} as={`/${record.CommodityID}`}>
                  <div className="card bg-base-100 shadow-xl text-black m-5 glass">
                    <div className="card-body">
                      <div
                        className="tooltip tooltip-open tooltip-right"
                        data-tip={`商品id:` + record.CommodityID}
                      >
                        <h2 className="card-title text-left">{record.CommodityName}</h2>
                        <p className="text-left break-all">购买价格：{record.ExpectPrice}</p>
                      </div>
                      <p className="text-3xl m-5 text-left break-all">"{record.StoryDetail}"</p>
                      <p className="text-left">商品购买时间：{record.Time}</p>
                    </div>
                  </div>
                </Link>
              ))
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default userBuyRecord;
