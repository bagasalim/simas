import Head from "next/head";
import Image from "next/image";
import styles from "../../../../styles/Home.module.css";
import Sidebar from "../sidebar";
import info1 from "../assets/info1.jpg";
import info2 from "../assets/info2.jpg";

export default function home() {
  return (
    <div className={styles.contentcontainer} style={{ border: "solid" }}>
      <h1>Selamat Datang Customer Service</h1>
      <Image src={info1} width={1200} height={400} />
      <h4>Waspada Penipuan, Begini Tips Transaksi Aman di ATM Bank Sinarmas</h4>
      <Image src={info2} width={1200} height={400} />
    </div>
  );
}
