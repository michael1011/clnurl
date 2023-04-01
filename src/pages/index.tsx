import Head from 'next/head'
import { Inter } from 'next/font/google'
import Button from '@mui/material/Button'
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { GetServerSidePropsContext, GetServerSidePropsResult } from 'next'
import QR from '@/components/QR'
import { useState } from 'react';
import Invoice from '@/components/Invoice';
import styles from '@/styles/Home.module.css'
import { defaultDescription, invoicePath, lightningPrefix, lnurlPath } from '@/utils/consts'

const inter = Inter({ subsets: ['latin'] })

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

type Props = {
  lnurl: string;
  invoiceUrl: string;
  description: string;
}

export default function Home({ lnurl, invoiceUrl, description }: Props) {
  const [invoice, setInvoice] = useState<string | null>(null);

  // TODO: input for amount
  const [amount, setAmount] = useState<number>(10000);

  const fetchInvoice = async () => {
    // TODO: handle error
    const res = await fetch(`${invoiceUrl}?amount=${amount}`);
    const { pr } = await res.json();

    setInvoice(pr);
  };
  
  return (
    <ThemeProvider theme={darkTheme}>
      <Head>
        <title>{description}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <main className={`${styles.main} ${inter.className}`}>
        <h1>{description}</h1>
        <QR content={lnurl}/>

        <Button variant='contained' onClick={fetchInvoice}>Fetch Invoice</Button>
        { invoice ? <Invoice invoice={invoice}/> : <></>}
      </main>
    </ThemeProvider>
  )
}

export async function getServerSideProps(
  { res }: GetServerSidePropsContext,
): Promise<GetServerSidePropsResult<Props>> {
  res.setHeader(
    'Cache-Control',
    'public, s-maxage=3600, stale-while-revalidate=59'
  )
  
  let endpoint = process.env.ENDPOINT!;
  if (endpoint.endsWith('/')) {
    endpoint = endpoint.slice(0, -1);
  }

  const lnurlRes = await fetch(`${endpoint}${lnurlPath}`)

  return {
    props: {
      lnurl: `${lightningPrefix}${await lnurlRes.json()}`,
      invoiceUrl: `${endpoint}${invoicePath}`,
      description: process.env.INVOICE_DESCRIPTION || defaultDescription,
    },
  };
}
