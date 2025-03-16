import { PRODUCTS } from "@/assets/data";
import ProductCard from "@/libs/globals/components/cards/ProductCard";
import Hero from "@/libs/globals/components/heros/Hero";
import MainLayout from "@/libs/globals/components/layouts/MainLayout";
import SectionTitle from "@/libs/globals/components/sections/SectionTitle";
import CollectionsSection from "@/libs/routes/home/CollectionsSection";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  head: () => ({
    meta: [{ title: "Home Page" }],
  }),
  component: HomePage,
});

function HomePage() {
  return (
    <MainLayout>
      <div className="w-full p-6 lg:p-[auto]">
        <Hero
          title="Lights up Your Home"
          actionText="Read more"
          captionText="They can easily highlight your dining room decor and area a great accessory."
          imageUrl="https://shop.getty.edu/cdn/shop/products/G019AH_1200x.jpg?v=1621057374"
          onActionClick={console.log}
        />
        <section className="flex w-full flex-col items-start md:gap-3 xl:gap-6 2xl:flex-row">
          <SectionTitle className="mb-3 md:mr-[inherit] xl:mr-56">
            Products
          </SectionTitle>
          <div className="mb-6 flex w-full flex-1 flex-col flex-wrap items-stretch gap-6 sm:flex-row lg:gap-6">
            {PRODUCTS.slice(0, 6).map((p, i) => (
              <ProductCard
                key={`product-card-${i}`}
                title={p.name}
                variants={p.colorOptions.map((v) => ({
                  imageUrl: v.imageUrl,
                  colorSwatch: v.colorSwatch,
                  price: 100.0 * Math.random(),
                }))}
              />
            ))}
          </div>
        </section>
        <section>
          <SectionTitle className="mb-4">Categories</SectionTitle>
          <CollectionsSection />
        </section>
      </div>
    </MainLayout>
  );
}
