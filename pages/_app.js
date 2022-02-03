import Head from "next/head";
import "../styles/globals.css";

function MyApp({ Component, pageProps }) {
  return (
    <>
      <Head>
        <title>LeetCode Stats</title>
        <meta
          name="description"
          content="Dynamically generated LeetCode stats for your GitHub READMEs!"
        />
        <link
          rel="icon"
          type="image/png"
          href="/favicon-16x16.png"
          sizes="16x16"
        />
        <link
          rel="icon"
          type="image/png"
          href="/favicon-32x32.png"
          sizes="32x32"
        />
        <link
          rel="icon"
          type="image/png"
          href="/favicon-96x96.png"
          sizes="96x96"
        />
        <link
          rel="icon"
          type="image/png"
          href="/favicon-192x192.png"
          sizes="192x192"
        />
      </Head>
      <Component {...pageProps} />
    </>
  );
}

export default MyApp;
