import '../styles/globals.css';
import type { AppProps } from 'next/app';
import { Toaster } from 'react-hot-toast';
import '@lizunlong/livekit-components-styles';
import '@lizunlong/livekit-components-styles/prefabs';
import { ThemeProvider } from '../components/ThemeProvider';

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <ThemeProvider attribute="class" defaultTheme="dark" enableSystem>
        <Component {...pageProps} />
        <Toaster />
      </ThemeProvider>
    </>
  );
}

export default MyApp;
