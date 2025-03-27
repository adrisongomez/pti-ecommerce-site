import { Box, BoxProps, Card, CircularProgress, useTheme } from "@mui/material";
import { CSSProperties, FC, useState } from "react";

type ImageProps = {
  width: CSSProperties["width"];
  height: CSSProperties["height"];
} & Omit<BoxProps<"img">, "onLoad" | "width" | "height" | "bgcolor">;

const Image: FC<ImageProps> = ({ width, height, ...props }) => {
  const [loading, setLoading] = useState(true);
  const theme = useTheme();
  return (
    <Card
      variant="elevation"
      onDragStart={(e) => {
        e.preventDefault();
      }}
      sx={{
        borderRadius: theme.shape.borderRadius,
        position: "relative",
        width: width,
        height: height,
      }}
    >
      <Box
        {...props}
        component="img"
        width="100%"
        height="100%"
        bgcolor={theme.palette.background.default}
        sx={{
          ...(props.sx ?? {}),
          objectFit: "contain",
        }}
        onLoad={() => {
          setLoading(false);
        }}
      />
      {loading && (
        <Box
          position="absolute"
          top={0}
          left={0}
          display="flex"
          alignItems="center"
          justifyContent="center"
          height="100%"
          width="100%"
        >
          <CircularProgress />
        </Box>
      )}
    </Card>
  );
};

export default Image;
