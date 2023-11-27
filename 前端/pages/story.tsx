// pages/story.tsx

import React, { useState, useEffect } from "react";
import StoryCard from "../components/StoryCard";
import Link from "next/link";

interface story {
  StoryID: number;
  StoryDetail: string;
  CID: number;
  Time: string;
  PublicAddress: string;
}

const Story = () => {
  const [stories, setStories] = useState<story[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const pubAddress =
    "80f442c18a3ad4f9944b5274a90add19b35cccd1fdd47ba9b28e91f67a3b19214723f7c9813e71c2e18a18bfc28414ed078934ace6adecd733fbb45f8d322646";
  const handleSearch = async () => {
    try {
      const response = await fetch("http://81.71.5.116:9090/storiesitem", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
      const data = await response.json();
      if (data) {
        setIsLoading(false);
      }
      console.log(data.storieslist);
      setStories(data.storieslist);
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    handleSearch();
  }, [pubAddress]);
  return (
    <div
      className="hero min-h-screen"
      style={{
        backgroundImage:
          "url(https://cdn.midjourney.com/167de30a-30b8-482b-8ee3-b1c513ef4332/0_3.webp)",
      }}
    >
      <div className="hero-overlay bg-opacity-60"></div>
      <div className="hero-content text-center text-neutral-content">
        <div className="w-full p-5">
          <h1 className="mb-5 text-6xl font-bold font-sans text-black">
            Let's See Other's Feelings
          </h1>
          <div
            className="mx-auto rounded-lg border glass flex-col p-5 overflow-x-auto"
            style={{ maxWidth: "80%" }}
          >
            {isLoading ? (
              <div className="product-display w-full mt-40 p-7 mx-auto rounded-lg border glass flex items-center justify-center">
                <span className="loading loading-bars loading-lg"></span>
                <span className="loading loading-bars loading-lg"></span>
                <span className="loading loading-bars loading-lg"></span>
              </div>
            ) : (
              stories.map((s, index) => (
                <Link key={index} href={`/[productId]`} as={`/${s.CID}`}>
                  <div className="card bg-base-100 shadow-xl text-black m-5 glass">
                    <div className="card-body">
                      <div
                        className="tooltip tooltip-open tooltip-right"
                        data-tip={`商品id:` + s.CID}
                      >
                        <h2 className="card-title text-left">用户地址：</h2>
                        <p className="text-left break-all">{s.PublicAddress}</p>
                      </div>
                      <p className="text-3xl m-5 text-left break-all">
                        "{s.StoryDetail}"
                      </p>
                      <p className="text-left">商品发布时间：{s.Time}</p>
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

export default Story;
