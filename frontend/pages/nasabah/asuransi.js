import Header from "../../components/Header";
import UserFooter from "../../components/userfooter";
import style from "./asuransi.module.scss";
import Image from "next/future/image";
import jumbotron from "../../public/jumbotron.png";
import cslaki from "../../public/cslaki.png";
import React from "react";
import ConfirmationModal from "../../components/modals/modalDetailCart";
import {useState, useEffect } from "react";



const Asuransi = () => {
  
  const [data, setData] = useState(null);
  const [newLink, setNewLink] = useState("");
  const [modalOpen, setModalOpen] = React.useState(false);
  const [body, setBodyData] = React.useState("");

const onSubmit = async (e) => {
  const dataform = {
    newlink: newLink,
  };
  setBodyData(dataform);
  setModalOpen(true);
};
  return (
    <div>
      <Header />
      <div>
        <Image className={style.jumbotron} src={jumbotron} alt="jumbotron" hidden/>
      </div>
      <div className={style.buttonpa}>
        <div>
          <button className={style.buttonpromo}>Promo</button>
        </div>
        <div>
          <button className={style.buttonasuransi}>Asuransi</button>
        </div>
      </div>

      <div className={style.container}>
      <form id="cartZoom"  onSubmit={(e) => {
            e.preventDefault();
          }}>
      <div>
        Judul
        <input type="text" name="newlink" required value="Judul" onChange={(e) => setNewLink(e.target.value)}/>
      </div>
      <div>Deskripsi</div>
      <div>
      <Image className={style.cslaki} src={cslaki} alt="cslaki" />
      </div>
      <div>
        SnK
      </div>
      <button onClick={onSubmit}>
        Lihat Detail
      </button>
      </form>
    
      </div>
      <ConfirmationModal show={modalOpen} close={() => setModalOpen(false)} />;


      <UserFooter />
    </div>
  );
};

export default Asuransi;
