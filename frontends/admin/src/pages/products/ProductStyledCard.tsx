import { Card, styled } from "@mui/material";

const ProductStyledCard = styled(Card)(({ theme }) => ({
  padding: theme.spacing(3),
  display: "flex",
  flexDirection: "column",
  gap: theme.spacing(3),
  borderRadius: theme.shape.borderRadius * 3,
}));
export default ProductStyledCard;
