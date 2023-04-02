import { TextField } from "@mui/material";
import styles from '@/styles/Invoice.module.css';
import QR from "./QR";

type Props = {
  invoice: string;
  invoiceText?: string;
}

export default function Invoice({ invoice, invoiceText }: Props) {
  return (
    <div className={styles.container}>
      <QR size={250} content={invoice}/>
      <TextField
        multiline
        value={invoiceText || invoice}
        className={styles.invoice}
        onClick={(val) => {
          (val.target as any).select();
        }}
      />
    </div>
  );
}
