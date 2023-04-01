import { useQRCode } from "next-qrcode";

type Props = {
  size?: number;
  content: string;
};

export default function QR({ content, size }: Props) {
  const { Canvas } = useQRCode();

  return (
    <Canvas
      text={content}
      options={{
        margin: 2,
        width: size || 200,
      }}
    />
  );
}
