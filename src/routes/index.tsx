import { PRODUCTS } from "@/assets/data";
import ProductCard from "@/libs/globals/components/cards/ProductCard";
import Hero from "@/libs/globals/components/heros/Hero";
import MainLayout from "@/libs/globals/components/layouts/MainLayout";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/")({
  head: () => ({
    meta: [{ title: "Home Page" }],
  }),
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <MainLayout>
      <div className="w-full">
        <Hero
          title="Lights up Your Home"
          actionText="Read more"
          captionText="They can easily highlight your dining room decor and area a great accessory."
          imageUrl="https://shop.getty.edu/cdn/shop/products/G019AH_1200x.jpg?v=1621057374"
          onActionClick={console.log}
        />
        <section className="w-full md:flex-col md:gap-3 xl:flex-row xl:gap-6">
          <h3 className="text-2xl text-(--text-accent) uppercase">Products</h3>
          <div
            id="collection-section"
            className="mt-6 flex flex-1 flex-wrap gap-6"
          >
            {PRODUCTS.map((p) => (
              <ProductCard
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
      </div>
    </MainLayout>
  );
}
