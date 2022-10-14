import style from "./halamanutama.module.scss";
import AdminCard from "../admincard/admincard";
import Image from "next/image";
import foto1 from "../../public/assets/info1.jpg";
import foto2 from "../../public/assets/info2.jpg";

const HalamanUtama = () => {
  const data  = localStorage.getItem('user')
  const newData = JSON.parse(data)

  return (
    <div className={style.utama}>
      <h1 className={style.title}>Selamat Datang {newData.name}</h1>
      <hr />
      <br />
      <div className={style.admincards}>
        <h3 style={{ fontSize: "24px", fontWeight: "450" }}>List Customer Service</h3>
        <AdminCard />
      </div>
      <br />
      <br />
      <div className={style.informasi}>
        <h3 style={{ fontSize: "24px", fontWeight: "450" }}>Informasi</h3>
        <Image src={foto1} width={700} height={325} alt="foto" />
        <p style={{ fontSize: "20px", textAlign: "justify" }}>Waspada Penipuan, Begini Tips Transaksi Aman di ATM Bank Sinarmas</p>
        <br />
        <br />
        <Image src={foto2} width={700} height={325} alt="foto" />
        <p style={{ fontSize: "20px", textAlign: "justify" }}>Amankan Kartu Kredit dengan Cara Freeze Lewat Aplikasi Simobi+</p>
      </div>
    </div>
  );
};

export default HalamanUtama;
