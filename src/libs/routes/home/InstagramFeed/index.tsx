import { FEED_INSTAGRAM } from "@/assets/data";
import Image from "@/libs/globals/components/images/Image";
import { FC } from "react";

const InstagramFeed: FC = () => {
  return (
    <div className="flex w-fit flex-row flex-wrap gap-6">
      {FEED_INSTAGRAM.map((v, i) => (
        <Image
          key={`instagram-feed-${i}`}
          className="rounded"
          src={v.imageUrl}
          height={v.hieght}
          width={v.width}
        />
      ))}
    </div>
  );
};

export default InstagramFeed;
