import Head from 'next/head';
import { Inter } from 'next/font/google';
import { purple } from '@mui/material/colors';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import { GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import { Limits } from '@/utils/types';
import Invoice from '@/components/Invoice';
import styles from '@/styles/Home.module.css';
import InvoiceFetcher from '@/components/InvoiceFetcher';
import {
  lnurlPath,
  invoicePath,
  lightningPrefix,
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
  lnurl: string;
  limits: Limits;
  invoiceUrl: string;
  description: string;
}

export default function Home({
  lnurl, invoiceUrl, description, limits
}: Props) {
  return (
    <ThemeProvider theme={darkTheme}>
      <Head>
        <title>{description}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <main className={`${styles.main} ${inter.className}`}>
        <h1>{description}</h1>
        <Invoice invoice={`${lightningPrefix}${lnurl}`} invoiceText={lnurl}/>

        <div className={styles.invoiceFetcher}>
          <InvoiceFetcher invoiceUrl={invoiceUrl} limits={limits} />
        </div>
      </main>
    </ThemeProvider>
  );
}

export async function getServerSideProps(
  { res }: GetServerSidePropsContext,
): Promise<GetServerSidePropsResult<Props>> {
  res.setHeader(
    'Cache-Control',
    'public, s-maxage=3600, stale-while-revalidate=59'
  );
  
  let endpoint = process.env.ENDPOINT!;
  if (endpoint.endsWith('/')) {
    endpoint = endpoint.slice(0, -1);
  }

  const lnurlRes = await fetch(`${endpoint}${lnurlPath}`);

  return {
    props: {
      invoiceUrl: `${endpoint}${invoicePath}`,
      lnurl: await lnurlRes.json(),
      description: process.env.INVOICE_DESCRIPTION || defaultDescription,
      limits: {
        min: Number(process.env.MIN_SENDABLE || defaultMinSendable),
        max: Number(process.env.MAX_SENDABLE || defaultMaxSendable)
      }
    },
  };
}
