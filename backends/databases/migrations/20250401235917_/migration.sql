/*
  Warnings:

  - Added the required column `message` to the `ChatMessage` table without a default value. This is not possible if the table is not empty.
  - Added the required column `source` to the `ChatMessage` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "ChatMessage" ADD COLUMN     "message" TEXT NOT NULL,
ADD COLUMN     "source" "ChatSource" NOT NULL;
