-- CreateTable
CREATE TABLE "_MediaToProduct" (
    "A" INTEGER NOT NULL,
    "B" INTEGER NOT NULL,

    CONSTRAINT "_MediaToProduct_AB_pkey" PRIMARY KEY ("A","B")
);

-- CreateIndex
CREATE INDEX "_MediaToProduct_B_index" ON "_MediaToProduct"("B");

-- AddForeignKey
ALTER TABLE "_MediaToProduct" ADD CONSTRAINT "_MediaToProduct_A_fkey" FOREIGN KEY ("A") REFERENCES "medias"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_MediaToProduct" ADD CONSTRAINT "_MediaToProduct_B_fkey" FOREIGN KEY ("B") REFERENCES "products"("id") ON DELETE CASCADE ON UPDATE CASCADE;
