import { useState } from "react";
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import { Limits } from "@/utils/types";
import Invoice from "@/components/Invoice";
import { msatFactor } from "@/utils/consts";
import styles from '@/styles/InvoiceFetcher.module.css';

type Props = {
  limits: Limits;
  invoiceEndpoint: string;
};

export default function InvoiceFetcher({ invoiceEndpoint, limits }: Props) {
  const [invoice, setInvoice] = useState<string | null>(null);
  const [amount, setAmount] = useState<number | null>(null);
  const [inputError, setInputError] = useState<string | null>(null);

  const validateInput = (newAmount: number | null) => {
    if (newAmount === null || Number.isNaN(newAmount)) {
      setInputError('Amount has to be set');
      return;
    }

    if (newAmount > limits.max) {
      setInputError(`Amounts has to be ${limits.max / msatFactor} or less`);
      return;
    }

    if (newAmount < limits.min) {
      setInputError(`Amounts has to be ${limits.min / msatFactor} or greater`);
      return;
    }

    setInputError(null);
    setAmount(newAmount);
  };

  const fetchInvoice = async () => {
    if (amount === null) {
      validateInput(null);
      return;
    }
    
    const res = await fetch(
      `${invoiceEndpoint}?amount=${encodeURIComponent(amount)}`,
    );
    const { pr } = await res.json();
    
    setInvoice(pr);
  };
  
  return (
    <div className={styles.container}>
      <div className={styles.inputContainer}>
        <TextField
          label="Invoice Satoshis"
          type="number"
          variant="outlined"
          error={inputError !== null}
          helperText={inputError}
          onChange={(val) => {
            validateInput((val.target as any).valueAsNumber * msatFactor);
          }}
        />
        <div className={styles.buttonContainer}>
          <Button
            variant='contained'
            onClick={fetchInvoice}
            disabled={inputError !== null}
          >
            Fetch Invoice
          </Button>
        </div>

      </div>
      {invoice ? <Invoice
        invoice={invoice}
      /> : null}
    </div>
  );
}
  