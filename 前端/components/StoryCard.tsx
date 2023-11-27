// components/StoryCard.tsx

import React from "react";

interface StoryCardProps {
  id: string;
  productName: string;
  user: string;
  story: string;
  timestamp: string;
}

const StoryCard: React.FC<StoryCardProps> = ({
  id,
  productName,
  user,
  story,
  timestamp,
}) => {
  const handleClick = () => {
    // Handle the click event, e.g., navigate to the product details page
    // You can use Next.js router for navigation
  };

  return (
   <div></div>
  );
};

export default StoryCard;
