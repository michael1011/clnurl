import { TextField } from "@mui/material";
import styles from '@/styles/Invoice.module.css';
import { lightningPrefix } from "@/utils/consts";
import QR from "./QR";

type Props = {
  invoice: string;
}

export default function Invoice({ invoice }: Props) {
  return (
    <div className={styles.container}>
      <QR size={250} content={`${lightningPrefix}${invoice}`}/>
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
