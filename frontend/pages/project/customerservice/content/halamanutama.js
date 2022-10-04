import Head from "next/head";
import Image from "next/image";
import styles from "../../../../styles/Home.module.css";
import Sidebar from "../sidebar";
import info1 from "../assets/info1.jpg";
import info2 from "../assets/info2.jpg";

export default function home() {
  return (
    <div
      className="position-absolute top-50 start-50 translate-middle"
      style={{ paddingBot: "22px" }}
    >
      <h2>Selamat Datang, Customer Service</h2>
      <h3>Informasi</h3>
      <Image src={info1} width={1200} height={400} />
      <br />
      <h4>Waspada Penipuan, Begini Tips Transaksi Aman di ATM Bank Sinarmas</h4>
      <br />
      <Image src={info2} width={1200} height={400} />
      <br />
      <h4>Amankan Kartu Kredit dengan Cara Freeze Lewat Aplikasi Simobi+</h4>
    </div>
  );
}
