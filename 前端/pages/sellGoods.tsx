// pages/sellGoods.tsx

import { useState } from "react";
import { useRouter } from "next/router";

const SellGoods = () => {
  const router = useRouter();
  const [privateKey, setPrivateKey] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [formData, setFormData] = useState({
    pic: null as File | null,
    extent: "",
    price: "",
    story: "",
    description: "",
    name: "",
    address:"80f442c18a3ad4f9944b5274a90add19b35cccd1fdd47ba9b28e91f67a3b19214723f7c9813e71c2e18a18bfc28414ed078934ace6adecd733fbb45f8d322646",
  });

  const handleChange = (name: string, value: string) => {
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  //   特殊处理图片上传
  const handleImgChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files=e.target.files;
    if (files&&files.length>0){
      const file = files[0];
      setFormData((prevData) => ({
        ...prevData,
        pic: file,
      }));
    }
    
  };

  //向后端提交商品信息
  const handleSubmit = async () => {
    // 只要点击提交，就把按钮设为不可用
    setIsLoading(true);
    // 创建一个 FormData 对象，用于上传文件
    const formdata = new FormData();

    // 将文件添加到 FormData
    if (formData.pic) {
      formdata.append("pic", formData.pic);
    }

    // 将其他表单字段添加到 FormData
    formdata.append("extent", formData.extent);
    formdata.append("price", formData.price);
    formdata.append("story", formData.story);
    formdata.append("description", formData.description);
    formdata.append("name", formData.name);
    formdata.append("address", formData.address);
    console.log(formData);
    console.log(formdata);
    // 向后端提交 FormData
    try {
      const response = await fetch("http://81.71.5.116:9090/releasegoods", {
        method: "POST",
        body: formdata,
      });

      // 处理后端返回的响应
      const data = await response.json();
      console.log(data)
      if (data) {
        alert("商品已发布");
        router.push("/");
      } else {
        alert("商品发布失败，请重试！");
      }
      // 处理其他逻辑...
      setIsLoading(false)
    } catch (error) {
      console.error("Error submitting form:", error);
      setIsLoading(false);
    }
  };

  return (
    <div
      className="hero min-h-screen bg-base-200"
      style={{
        backgroundImage:
          'url("https://cdn.midjourney.com/167de30a-30b8-482b-8ee3-b1c513ef4332/0_0.webp")',
        backgroundSize: "cover",
        backgroundPosition: "center",
      }}
    >
      {/* 这是确认发布商品的modal */}
      <dialog id="publishGoods_model" className="modal">
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
          <h1 className="text-4xl font-bold">出售确认</h1>
          <p className="py-6">
            您确认希望以{formData.price}元出售{formData.name}吗？
          </p>

          {/* 输入私钥 */}
          <input
            type="password"
            placeholder="请输入您的私钥"
            value={privateKey}
            onChange={(e) => setPrivateKey(e.target.value)}
            className="input input-bordered mb-4"
          />

          {/* 确认发布按钮 */}
          <div className="flex justify-end">
            <button
              className={`btn btn-primary mr-2 ${
                isLoading ? "cursor-not-allowed" : ""
              }`}
              onClick={() => handleSubmit()}
              disabled={isLoading}
            >
              {isLoading ? "发布中..." : "确认发布"}
            </button>
          </div>
        </div>
      </dialog>
      <div className="hero-overlay bg-opacity-40"></div>
      <div className="hero-content text-center flex-col lg:flex-row-reverse">
        <div className="w-full">
          <div className="text-center lg:text-left">
            <h1 className="text-6xl font-bold text-white font-sans">
              Share Your Goods & Story with us.
            </h1>
            {/* <p className="py-6">Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.</p> */}
          </div>
          <div className="card shrink-0 shadow-2xl bg-base-100 glass mt-10">
            <form className="card-body">
              <div className="form-control">
                <label className="label">
                  <span className="label-text">商品名称</span>
                </label>
                <input
                  type="text"
                  placeholder="name"
                  className="input input-bordered"
                  value={formData.name}
                  onChange={(e) => handleChange("name", e.target.value)}
                  required
                />
              </div>
              <div className="form-control">
                <label className="label">
                  <span className="label-text">商品简介</span>
                </label>
                <input
                  type="text"
                  placeholder="description"
                  className="input input-bordered"
                  value={formData.description}
                  onChange={(e) => handleChange("description", e.target.value)}
                  required
                />
              </div>
              <div className="form-control">
                <label className="label">
                  <span className="label-text">商品图片</span>
                </label>
                <input
                  type="file"
                  className="file-input file-input-bordered file-input-info w-full max-w-xs"
                  onChange={(e) => handleImgChange(e)}
                />
              </div>
              <div className="form-control">
                <label className="label">
                  <span className="label-text">商品折损程度</span>
                </label>
                <input
                  type="text"
                  placeholder="extent"
                  className="input input-bordered"
                  value={formData.extent}
                  onChange={(e) => handleChange("extent", e.target.value)}
                  required
                />
              </div>
              <div className="form-control">
                <label className="label">
                  <span className="label-text">商品期望价格</span>
                </label>
                <input
                  type="text"
                  placeholder="price"
                  className="input input-bordered"
                  value={formData.price}
                  onChange={(e) => handleChange("price", e.target.value)}
                  required
                />
              </div>
              <div className="form-control">
                <label className="label">
                  <span className="label-text">商品与你之间的故事</span>
                  <span className="label-text-alt">就让这份记忆长存世间~</span>
                </label>
                <textarea
                  className="textarea textarea-bordered h-24"
                  placeholder="story"
                  value={formData.story}
                  onChange={(e) => handleChange("story", e.target.value)}
                  required
                ></textarea>
              </div>
              <div className="form-control mt-6">
                <button
                  className="btn btn-primary"
                  onClick={() => {
                    const publish_modal = document.getElementById(
                      "publishGoods_model"
                    ) as HTMLDialogElement | null;
                    if (publish_modal) {
                      publish_modal.showModal();
                    }
                  }}
                  type="button"
                >
                  发布
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SellGoods;
