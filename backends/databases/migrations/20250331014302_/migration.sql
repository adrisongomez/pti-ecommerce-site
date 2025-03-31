/*
  Warnings:

  - Made the column `user_id` on table `orders` required. This step will fail if there are existing NULL values in that column.
  - Made the column `user_email` on table `orders` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE "orders" DROP CONSTRAINT "orders_user_id_fkey";

-- AlterTable
ALTER TABLE "addresses" ADD COLUMN     "zip_code" TEXT,
ALTER COLUMN "address_line_2" DROP NOT NULL;

-- AlterTable
ALTER TABLE "orders" ALTER COLUMN "user_id" SET NOT NULL,
ALTER COLUMN "user_email" SET NOT NULL;

-- AddForeignKey
ALTER TABLE "orders" ADD CONSTRAINT "orders_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE;
