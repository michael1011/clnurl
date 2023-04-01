import { useQRCode } from "next-qrcode";

type Props = {
  content: string;
};

export default function QR({ content }: Props) {
  const { Canvas } = useQRCode();

  return (
    <Canvas
      text={content}
      options={{
        margin: 2,
        width: 200,
        color: {
        },
      }}
    />
  );
}
