/*
  Warnings:

  - You are about to drop the `_ProductToTag` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `product_tags` table. If the table is not empty, all the data it contains will be lost.
  - You are about to drop the `tags` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `size` to the `medias` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "_ProductToTag" DROP CONSTRAINT "_ProductToTag_A_fkey";

-- DropForeignKey
ALTER TABLE "_ProductToTag" DROP CONSTRAINT "_ProductToTag_B_fkey";

-- DropForeignKey
ALTER TABLE "product_tags" DROP CONSTRAINT "product_tags_product_id_fkey";

-- DropForeignKey
ALTER TABLE "product_tags" DROP CONSTRAINT "product_tags_tag_id_fkey";

-- DropForeignKey
ALTER TABLE "product_variants" DROP CONSTRAINT "product_variants_feature_media_id_fkey";

-- AlterTable
ALTER TABLE "medias" ADD COLUMN     "size" INTEGER NOT NULL;

-- AlterTable
ALTER TABLE "product_medias" ADD COLUMN     "alt" TEXT;

-- AlterTable
ALTER TABLE "product_variants" ALTER COLUMN "feature_media_id" DROP NOT NULL;

-- AlterTable
ALTER TABLE "products" ADD COLUMN     "tags" TEXT[];

-- DropTable
DROP TABLE "_ProductToTag";

-- DropTable
DROP TABLE "product_tags";

-- DropTable
DROP TABLE "tags";

-- AddForeignKey
ALTER TABLE "product_variants" ADD CONSTRAINT "product_variants_feature_media_id_fkey" FOREIGN KEY ("feature_media_id") REFERENCES "product_medias"("id") ON DELETE SET NULL ON UPDATE CASCADE;
