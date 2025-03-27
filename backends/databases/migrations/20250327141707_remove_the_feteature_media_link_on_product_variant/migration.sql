/*
  Warnings:

  - You are about to drop the column `feature_media_id` on the `product_variants` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "product_variants" DROP CONSTRAINT "product_variants_feature_media_id_fkey";

-- DropIndex
DROP INDEX "product_variants_feature_media_id_key";

-- AlterTable
ALTER TABLE "product_variants" DROP COLUMN "feature_media_id",
ADD COLUMN     "featureMediaSort" INTEGER;
