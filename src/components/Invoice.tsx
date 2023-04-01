import QR from "./QR";

type Props = {
  invoice: string;
}

export default function Invoice({ invoice }: Props) {
  return (
    <div>
      <p style={{ maxWidth: '100px'}}>{invoice}</p>
      <QR content={invoice}/>
    </div>
  );
}
