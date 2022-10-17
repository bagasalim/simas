import React,{ useState, useEffect } from "react";
import Header from "../../components/Header";
import UserFooter from "../../components/userfooter";
import style from "./asuransi.module.scss";
import Image from "next/future/image";
import jumbotron from "../../public/jumbotron.png";
import ConfirmationModal from "../../components/modals/modalDetailPromo";
import promotionDummy from "../../public/promotion.jpg";
import { useRouter } from "next/router";

const Promo = () => {
  const [data, setData] = useState(null);
  const [newLink, setNewLink] = useState("");
  const [modalOpen, setModalOpen] = React.useState(false);
  const [body, setBodyData] = React.useState("");
  const router = useRouter();

  const onSubmit = async (e) => {
    const dataform = {
      newlink: newLink,
    };
    setBodyData(dataform);
    setNewLink(newLink);
    setModalOpen(true);
  };

  const asuransi = () => {
    router.push("/nasabah/asuransi");
  };

  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    try {
      const data = [
        { judul: "Serunya Oktober", startdate : "1 Oktober 2022",enddate: "31 Oktober 2022", kodepromo: "S1234" },
        { judul: "Serunya Oktober", startdate : "1 Oktober 2022",enddate: "31 Oktober 2022", kodepromo: "S1234" },
        { judul: "Serunya Oktober", startdate : "1 Oktober 2022",enddate: "31 Oktober 2022", kodepromo: "S1234" },
        { judul: "Serunya Oktober", startdate : "1 Oktober 2022",enddate: "31 Oktober 2022", kodepromo: "S1234" },
        { judul: "Serunya Oktober", startdate : "1 Oktober 2022",enddate: "31 Oktober 2022", kodepromo: "S1234" }
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
          <button className={style.buttonpromoActive}>Promo</button>
        </div>
        <div>
          <button className={style.buttonasuransi} onClick={asuransi}>Asuransi</button>
        </div>
      </div>
      
      {getData}
        < div className="row justify-content-start" style={{paddingLeft:80, paddingRight:80}}>
        {data?.map((item, index) => (
        <div key={index} className="col-4" style={{paddingLeft:50, paddingRight:50}}>
          <div className={style.detailContent}>
              <Image className={style.imageDummy} src={promotionDummy} alt="promotionDummy" />
              <h2 className={style.textContent}>{item.judul}</h2>
              <h5 className={style.textContent}>Periode: {item.startdate} - {item.enddate}</h5>
              <h5 className={style.textContent}>Kode Promo: {item.kodepromo}</h5>
              <button className={style.buttonDetail} onClick={onSubmit } data={body}>Lihat Detail</button>
          </div>
        </div>
        ))}
        </div>

      <ConfirmationModal show={modalOpen} close={() => setModalOpen(false)} />;
      <UserFooter />

    </div>
  );
};

export default Promo;
