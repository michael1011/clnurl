import { TextField } from "@mui/material";
import QR from "./QR";
import styles from '@/styles/Invoice.module.css'

type Props = {
  invoice: string;
}

export default function Invoice({ invoice }: Props) {
  return (
    <div className={styles.container}>
      <QR size={250} content={invoice}/>
      <TextField
        multiline
        value={invoice}
        className={styles.invoice}
        onClick={(val) => {
          (val.target as any).select();
        }}
      />
    </div>
  );
}
