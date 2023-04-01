import QR from "./QR";
import styles from '@/styles/Invoice.module.css'

type Props = {
  invoice: string;
}

export default function Invoice({ invoice }: Props) {
  return (
    <div className={styles.container}>
      <QR content={invoice}/>
      <p className={styles.invoice}>{invoice}</p>
    </div>
  );
}
