import { Box, BoxProps, Card, CircularProgress } from "@mui/material";
import { CSSProperties, FC, useState } from "react";

type ImageProps = {
  width: CSSProperties["width"];
  height: CSSProperties["height"];
} & Omit<BoxProps<"img">, "onLoad" | "width">;

const Image: FC<ImageProps> = ({ width, height, ...props }) => {
  const [loading, setLoading] = useState(true);
  return (
    <Card
      raised
      variant="outlined"
      sx={{
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
