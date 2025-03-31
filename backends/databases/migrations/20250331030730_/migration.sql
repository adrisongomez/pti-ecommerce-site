/*
  Warnings:

  - Changed the type of `totalPrice` on the `orders` table. No cast exists, the column would be dropped and recreated, which cannot be done if there is data, since the column is required.

*/
-- AlterTable
ALTER TABLE "orders" DROP COLUMN "totalPrice",
ADD COLUMN     "totalPrice" MONEY NOT NULL;
