/*
  Warnings:

  - You are about to drop the `credentials` table. If the table is not empty, all the data it contains will be lost.
  - Added the required column `password_has` to the `users` table without a default value. This is not possible if the table is not empty.

*/
-- DropForeignKey
ALTER TABLE "credentials" DROP CONSTRAINT "credentials_user_id_fkey";

-- AlterTable
ALTER TABLE "users" ADD COLUMN     "password_has" TEXT NOT NULL;

-- DropTable
DROP TABLE "credentials";
