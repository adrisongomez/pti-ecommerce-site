/*
  Warnings:

  - You are about to drop the column `featureMediaSort` on the `product_variants` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "product_variants" DROP COLUMN "featureMediaSort",
ADD COLUMN     "feature_media_loc_no" INTEGER;
