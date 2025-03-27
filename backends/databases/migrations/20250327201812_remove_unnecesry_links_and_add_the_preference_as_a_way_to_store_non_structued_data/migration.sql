/*
  Warnings:

  - You are about to drop the column `billing_address_id` on the `users` table. All the data in the column will be lost.
  - You are about to drop the column `shipping_address_id` on the `users` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "users" DROP CONSTRAINT "users_billing_address_id_fkey";

-- DropForeignKey
ALTER TABLE "users" DROP CONSTRAINT "users_shipping_address_id_fkey";

-- AlterTable
ALTER TABLE "users" DROP COLUMN "billing_address_id",
DROP COLUMN "shipping_address_id",
ADD COLUMN     "preference" JSONB;
