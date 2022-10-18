import style from "./halamanutama.module.scss";
import Image from "next/image";
import foto1 from "../../public/assets/info1.jpg";
import foto2 from "../../public/assets/info2.jpg";

import {useState, useEffect} from 'react'


const HalamanUtama = () => {
  const data_user  = localStorage.getItem('user')
  const newData = JSON.parse(data_user)

  const date = newData.lastlogin;
  const [data, setData] = useState('');
  const item = localStorage.getItem('location')
  const obj = JSON.parse(item);
  const [error, setError] = useState(null);
  
  useEffect(() => {
    if(localStorage.getItem('location') === null) {
      navigator.geolocation.getCurrentPosition(function(position) {
        getLocation(position.coords.latitude, position.coords.longitude)
      });
    }
  }, []);

  const getLocation = async (lat, lang) => {
    try {
      const newUrl = `https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=${lat}&longitude=${lang}&localityLanguage=en`
      const res = await fetch(newUrl);
      const data = await res.json();
      setData(data);
      localStorage.setItem("location", JSON.stringify(data))
      //console.log(data)
    }
    catch (error) {
      setError(error);
    }
  }

  return (
 
      <div className={style.utama}>
        <h1 className={style.title}>Selamat Datang {newData.name}</h1>
        <hr />
        {obj !== null ? 
          <div className={style.alert}>
          <p>Login At : {date}</p> 
          <p>Current Location : {obj.locality} {obj.city} {obj.principalSubdivision} {obj.countryName} </p> 
          </div> : ''
        }
        <br />
        <br />
        <div className={style.informasi}>
          <h3 style={{ fontSize: "24px", fontWeight: "450" }}>Informasi</h3>

          <Image src={foto1} width={700} height={325} alt="foto1" />

          <p style={{ fontSize: "20px", textAlign: "justify" }}>Waspada Penipuan, Begini Tips Transaksi Aman di ATM Bank Sinarmas</p>
          <br />
          <br />

          <Image src={foto2} width={700} height={325} alt="foto2" />
          <p style={{ fontSize: "20px", textAlign: "justify" }}>Amankan Kartu Kredit dengan Cara Freeze Lewat Aplikasi Simobi+</p>
        </div>
      </div>
    
  );
};

export default HalamanUtama;
