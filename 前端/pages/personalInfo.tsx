// pages/story.tsx

import React, { useState, useEffect } from "react";
import StoryCard from "../components/StoryCard";
import Link from "next/link";
import { useRouter } from "next/router";

const PersonalInfo = () => {
  const [point, setPoint] = useState(0);
  const [isLoading,setIsLoading]=useState(true);
  const router=useRouter();
  const pubAddress='80f442c18a3ad4f9944b5274a90add19b35cccd1fdd47ba9b28e91f67a3b19214723f7c9813e71c2e18a18bfc28414ed078934ace6adecd733fbb45f8d322646';
  const handleSearch = async () => {
    try {
      const response = await fetch(`http://81.71.5.116:9090/goodsitemsell?address=${pubAddress}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(error);
    }
  };

  //查询积分
  const indexPoint=async()=>{
    try {
        const response = await fetch(`http://81.71.5.116:9090/searchscore?address=${pubAddress}`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        });
        const data = await response.json();
        console.log(data);
        if(data){
          setIsLoading(false);
        }
        setPoint(data.balance);
      } catch (error) {
        console.error(error);
      }
  };

  //跳转购买商品列表页面
  const handleBuyPage=()=>{
    router.push({
        pathname:"/userBuyPage",
        query:{address:pubAddress},
    });
  };
  //跳转卖出商品列表页面
  const handleSellPage=()=>{
    router.push({
        pathname:"/userSellPage",
        query:{address:pubAddress},
    });
  };

  useEffect(() => {
    handleSearch();
    //查询积分
    indexPoint();
  },[pubAddress]);
  return (
    <div
      className="hero min-h-screen"
      style={{
        backgroundImage:
          "url(https://cdn.midjourney.com/167de30a-30b8-482b-8ee3-b1c513ef4332/0_1.webp)",
      }}
    >
      <div className="hero-overlay bg-opacity-60"></div>
      <div className="hero-content text-center text-neutral-content">
        <div className="w-full p-5">
          <h1 className="mb-5 text-6xl font-bold font-sans text-white">
            个人信息
          </h1>
          <div
            className="mx-auto rounded-lg border glass flex-col p-5 overflow-x-auto"
            style={{ maxWidth: "80%" }}
          >
            <div className="flex">
            <div className="avatar placeholder m-5">
              <div className="bg-neutral text-neutral-content rounded-full w-24">
                <span className="text-3xl">D</span>
              </div>
            </div>
            <div className="flex-col flex items-center m-5">
              <p className="text-left mb-1 text-2xl">用户地址：</p>
              <p className="text-sm break-all">80f442c18a3ad4f9944b5274a90add19b35cccd1fdd47ba9b28e91f67a3b19214723f7c9813e71c2e18a18bfc28414ed078934ace6adecd733fbb45f8d322646</p>
              </div>
            </div>
            
             {/* 积分 */}
            <div className="m-5 card border hover:shadow-lg transition-all duration-300 glass cursor-pointer flex items-center" onClick={handleSearch}>
            <p className="text-white text-xl font-bold">积分:{isLoading?(<span className="loading loading-dots loading-md"></span>):(point)}</p>
            </div>

            {/* 购买记录 */}
            <div className=" m-5 card border hover:shadow-lg transition-all duration-300 glass cursor-pointer" onClick={handleBuyPage}>
            <p className="text-white text-xl font-bold">购买记录</p>
            </div>

            {/* 卖出记录 */}
            <div className="m-5 card border hover:shadow-lg transition-all duration-300 glass cursor-pointer" onClick={handleSellPage}>
            <p className="text-white text-xl font-bold">发布/卖出记录</p>
            </div>

          </div>
        </div>
      </div>
    </div>
  );
};

export default PersonalInfo;
