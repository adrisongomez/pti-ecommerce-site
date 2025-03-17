import Gallery from "@/libs/globals/components/galleries/Gallery";
import Hero from "@/libs/globals/components/heros/Hero";
import { FC, useState } from "react";

const MainHeroGallery: FC = () => {
  const [currentValue, setCurrentValue] = useState(0);
  return (
    <Gallery disabled value={currentValue} onChange={(v) => setCurrentValue(v)}>
      <Hero
        title="Lights up Your Home"
        actionText="Read more"
        captionText="They can easily highlight your dining room decor and area a great accessory."
        imageUrl="https://shop.getty.edu/cdn/shop/products/G019AH_1200x.jpg?v=1621057374"
        onActionClick={console.log}
      />
    </Gallery>
  );
};

export default MainHeroGallery;
