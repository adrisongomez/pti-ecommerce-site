import ProductCard from "@/libs/globals/components/cards/ProductCard";
import MainLayout from "@/libs/globals/components/layouts/MainLayout";
import { PLACEHOLDER_IMAGE } from "@/libs/globals/constants";
import SectionTitle from "@/libs/globals/components/sections/SectionTitle";
import { useSvcProductsServiceGetApiProducts } from "@/libs/globals/generated/queries";
import MainHeroGallery from "@/libs/routes/home/MainHeroGallery";
import { createFileRoute, Link } from "@tanstack/react-router";
import { useAppDispatch } from "@/libs/globals/hooks/redux";
import CartSlider from "@/libs/globals/redux/CartReducer";

export const Route = createFileRoute("/")({
  head: () => ({
    meta: [{ title: "Home Page" }],
  }),
  component: HomePage,
});

function HomePage() {
  const dispatch = useAppDispatch();
  const products = useSvcProductsServiceGetApiProducts({}, [
    "HomePage_Products_QUery",
  ]);
  return (
    <MainLayout>
      <div className="flex w-full flex-col gap-20 p-6 lg:p-[auto]">
        <section className="w-full">
          <MainHeroGallery />
        </section>
        <section className="flex w-full flex-col items-start md:gap-3">
          <SectionTitle className="mb-3 md:mr-[inherit] xl:mr-56">
            Products
          </SectionTitle>
          <div className="flex w-full flex-1 flex-col flex-wrap items-stretch gap-6 sm:flex-row lg:gap-6">
            {products.data?.data.map((p, i) => (
              <Link
                key={`product-card-${i}`}
                to="/products/$productId"
                params={{ productId: p.id.toString() }}
              >
                <ProductCard
                  title={p.title}
                  variants={p.variants.map((v, i) => ({
                    imageUrl: p.medias.at(i)?.url ?? PLACEHOLDER_IMAGE ?? "#",
                    colorSwatch: v.colorHex ?? "",
                    price: v.price,
                  }))}
                  onClick={(e) => {
                    e.preventDefault();
                  }}
                  onAddToCarcClick={() => {
                    console.log("here");
                    dispatch(CartSlider.actions.addCart(p));
                  }}
                />
              </Link>
            ))}
          </div>
        </section>
      </div>
    </MainLayout>
  );
}
