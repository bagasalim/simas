import React,{ useState, useEffect } from "react";
import Header from "../../components/Header";
import UserFooter from "../../components/userfooter";
import style from "./asuransi.module.scss";
import Image from "next/future/image";
import jumbotron from "../../public/jumbotron.png";
import ConfirmationModal from "../../components/modals/modalDetailCart";
import covid from "../../public/covid.jpg";
import { useRouter } from "next/router";

const Asuransi = () => {
  const [data, setData] = useState(null);
  const [newLink, setNewLink] = useState("");
  const [modalOpen, setModalOpen] = useState(false);
  const [body, setBodyData] = useState("");
  const router = useRouter();

  const onSubmit = async (e) => {
    const dataform = {
      newlink: newLink,
    };
    setBodyData(dataform);
    setNewLink(newLink);
    setModalOpen(true);
  };

  const promo = () => {
    router.push("/nasabah/promo");
  };

  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    try {

      const data = [
        { judul: "Simas Covid", premi: "Rp. 50.000",uangpertanggungan: "Rp. 10.000.000" },
        { judul: "Simas Covid", premi: "Rp. 50.000",uangpertanggungan: "Rp. 10.000.000" },
        { judul: "Simas Covid", premi: "Rp. 50.000",uangpertanggungan: "Rp. 10.000.000" },
        { judul: "Simas Covid", premi: "Rp. 50.000",uangpertanggungan: "Rp. 10.000.000" },
        { judul: "Simas Covid", premi: "Rp. 50.000",uangpertanggungan: "Rp. 10.000.000" }
      ];
      setData(data);
    }
    catch (e) {
      if (typeof e === "string") {
        alert("Gagal load data");
      }
    }
  }

  return (
    <div>
      <Header />
      <div>
        <Image className={style.jumbotron} src={jumbotron} alt="jumbotron" hidden/>
      </div>
      <div className={style.buttonpa}>
        <div>
          <button className={style.buttonpromo} onClick={promo}>Promo</button>
        </div>
        <div>
          <button className={style.buttonasuransiActive}>Asuransi</button>
        </div>
      </div>
      
      {getData}
        < div className="row justify-content-start" style={{paddingLeft:80, paddingRight:80}}>
        {data?.map((item, index) => (
        <div key={index} className="col-4" style={{paddingLeft:50, paddingRight:50}}>
          <div className={style.detailContent}>
              <Image className={style.imageDummy} src={covid} alt="covid" />
              <h2 className={style.textContent}>{item.judul}</h2>
              <h5 className={style.textContent}>Premi: {item.premi}</h5>
              <h5 className={style.textContent}>Uang Pertanggungan: {item.uangpertanggungan}</h5>
              <button className={style.buttonDetail} onClick={onSubmit} data={body}>Lihat Detail</button>
          </div>
        </div>
        ))}
        </div>

      <ConfirmationModal show={modalOpen} close={() => setModalOpen(false)} />;
      <UserFooter />

    </div>
  );
};

export default Asuransi;
