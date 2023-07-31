import { useQRCode } from "next-qrcode";

type Props = {
  size?: number;
  content: string;
};

export default function QR({ content, size }: Props) {
  const { SVG } = useQRCode();

  return (
    <a href={content}>
      <SVG
        text={content}
        options={{
          margin: 2,
          width: size || 200,
        }}
      />
    </a>
  );
}
