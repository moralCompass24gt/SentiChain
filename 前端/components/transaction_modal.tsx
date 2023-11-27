/* 交易的modal，这里交易就不单独做一个页面了，只是一个modal来确认是否交易，如果用户输入私钥确认交易，则调用交易结果，等待交易确认后，则出现一个toast提示交易已完成，并且页面跳转到市场发现首页 */
import { useState } from "react";
import { ProductInfo } from "./ProductDetail";

interface TransactionModalProps {
  productInfo: ProductInfo;
  onClose: () => void;
}

const TransactionModal: React.FC<TransactionModalProps> = ({
  productInfo,
  onClose,
}) => {
  const [privateKey, setPrivateKey] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleConfirm = async () => {
    setIsLoading(true);
    // 在这里调用交易接口，使用 privatekey 参数
    // 假设交易接口返回 true 表示交易成功
    const transactionSuccess = await fetch("/api/product", {
      method: "POST",
      body: JSON.stringify({ privateKey }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = await transactionSuccess.json();

    if (transactionSuccess) {
      // 交易成功后执行一些操作，比如显示提示，跳转页面等
      alert("交易已完成");
      onClose();
    } else {
      // 交易失败时处理
      alert("交易失败，请重试");
    }

    setIsLoading(false);
  };

  return (
    <dialog id="transaction_model" className="modal">
      <div className="modal-box">
        <form method="dialog">
          {/* if there is a button in form, it will close the modal */}
          <button
            className="btn btn-sm btn-circle btn-ghost absolute right-2 top-2"
            onClick={onClose}
          >
            ✕
          </button>
        </form>
        <h1 className="text-5xl font-bold">交易确认</h1>
        <p className="py-6">
          您确认希望以{productInfo.price}元购买{productInfo.name}吗？
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
          <button className="btn btn-secondary" onClick={onClose}>
            取消
          </button>
        </div>
      </div>
    </dialog>
  );
};

export default TransactionModal;
