-- CreateEnum
CREATE TYPE "UserRole" AS ENUM ('CUSTOMER', 'ADMIN');

-- DropForeignKey
ALTER TABLE "users" DROP CONSTRAINT "users_billing_address_id_fkey";

-- DropForeignKey
ALTER TABLE "users" DROP CONSTRAINT "users_shipping_address_id_fkey";

-- AlterTable
ALTER TABLE "addresses" ADD COLUMN     "deletedAt" TIMESTAMP(3),
ADD COLUMN     "role" "UserRole" NOT NULL DEFAULT 'CUSTOMER';

-- AlterTable
ALTER TABLE "users" ALTER COLUMN "billing_address_id" DROP NOT NULL,
ALTER COLUMN "shipping_address_id" DROP NOT NULL;

-- AddForeignKey
ALTER TABLE "users" ADD CONSTRAINT "users_shipping_address_id_fkey" FOREIGN KEY ("shipping_address_id") REFERENCES "addresses"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "users" ADD CONSTRAINT "users_billing_address_id_fkey" FOREIGN KEY ("billing_address_id") REFERENCES "addresses"("id") ON DELETE SET NULL ON UPDATE CASCADE;
