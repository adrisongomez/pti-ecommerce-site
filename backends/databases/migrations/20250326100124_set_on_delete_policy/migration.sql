-- DropForeignKey
ALTER TABLE "product_collections" DROP CONSTRAINT "product_collections_collection_id_fkey";

-- DropForeignKey
ALTER TABLE "product_collections" DROP CONSTRAINT "product_collections_product_id_fkey";

-- DropForeignKey
ALTER TABLE "product_medias" DROP CONSTRAINT "product_medias_media_id_fkey";

-- DropForeignKey
ALTER TABLE "product_medias" DROP CONSTRAINT "product_medias_product_id_fkey";

-- DropForeignKey
ALTER TABLE "product_variants" DROP CONSTRAINT "product_variants_product_id_fkey";

-- DropForeignKey
ALTER TABLE "products" DROP CONSTRAINT "products_vendor_id_fkey";

-- AlterTable
ALTER TABLE "products" ALTER COLUMN "vendor_id" DROP NOT NULL;

-- AddForeignKey
ALTER TABLE "products" ADD CONSTRAINT "products_vendor_id_fkey" FOREIGN KEY ("vendor_id") REFERENCES "vendors"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "product_collections" ADD CONSTRAINT "product_collections_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "product_collections" ADD CONSTRAINT "product_collections_collection_id_fkey" FOREIGN KEY ("collection_id") REFERENCES "collections"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "product_variants" ADD CONSTRAINT "product_variants_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "product_medias" ADD CONSTRAINT "product_medias_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "product_medias" ADD CONSTRAINT "product_medias_media_id_fkey" FOREIGN KEY ("media_id") REFERENCES "medias"("id") ON DELETE CASCADE ON UPDATE CASCADE;
