import Head from 'next/head';
import { Inter } from 'next/font/google';
import { useEffect, useState } from 'react';
import { GetStaticPropsResult } from 'next';
import { purple } from '@mui/material/colors';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { Limits } from '@/utils/types';
import Invoice from '@/components/Invoice';
import styles from '@/styles/Home.module.css';
import InvoiceFetcher from '@/components/InvoiceFetcher';
import {
  lnurlPath,
  invoicePath,
  defaultDescription,
  defaultMaxSendable,
  defaultMinSendable,
} from '@/utils/consts';

const inter = Inter({ subsets: ['latin'] });

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: purple[800],
    },
  },
  typography: {
    button: {
      textTransform: 'none',
    },
  }
});

type Props = {
  limits: Limits;
  description: string;
  lnurlEndpoint: string;
  invoiceEndpoint: string;
}

export default function Home({
  lnurlEndpoint, invoiceEndpoint, description, limits
}: Props) {
  const [lnurl, setLnurl] = useState(null);
  
  useEffect(() => {
    fetch(lnurlEndpoint)
      .then((res) => res.json())
      .then((res) => setLnurl(res));
  }, [lnurlEndpoint]);

  return (
    <ThemeProvider theme={darkTheme}>
      <Head>
        <title>{description}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <main className={`${styles.main} ${inter.className}`}>
        <h1>{description}</h1>
        { lnurl !== null ?
          <>
            <Invoice invoice={lnurl}/>

            <div className={styles.invoiceFetcher}>
              <InvoiceFetcher
                limits={limits}
                invoiceEndpoint={invoiceEndpoint}
              />
            </div>
          </> : null} 
      </main>
    </ThemeProvider>
  );
}

export async function getStaticProps(): Promise<GetStaticPropsResult<Props>> {
  let endpoint = process.env.ENDPOINT!;
  if (endpoint.endsWith('/')) {
    endpoint = endpoint.slice(0, -1);
  }

  return {
    props: {
      lnurlEndpoint: `${endpoint}${lnurlPath}`,
      invoiceEndpoint: `${endpoint}${invoicePath}`,
      description: process.env.INVOICE_DESCRIPTION || defaultDescription,
      limits: {
        min: Number(process.env.MIN_SENDABLE || defaultMinSendable),
        max: Number(process.env.MAX_SENDABLE || defaultMaxSendable)
      }
    },
  };
}
