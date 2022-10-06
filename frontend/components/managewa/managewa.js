import React from "react";
import style from "./managewa.module.scss";
import BerhasilSubmit from "../modals/berhasilsubmit";
import ConfirmationModal from "../modals/modalwadanzoom";

const ManageWa = () => {
  const [modalOpen, setModalOpen] = React.useState(false);
  const onSubmit = async (e) => {
    e.preventDefault();
    modalOpen = setModalOpen(true);
  };
  const closed = async (e) => {
    e.preventDefault();
    modalOpen = setModalOpen(!modalOpen);
  };
  return (
    <div className={style.wa}>
      <h1>Manage Link WhatsApp</h1>
      <div className={style.inputbox} style={{ borderRadius: "16px" }}>
        <form onSubmit={onSubmit}>
          <div>
            <h3>Link WhatsApp Lama</h3>
            <input className={style.readonly} type="text" placeholder="https://api.whatsapp.com/send?phone=6288221500153" disabled="true" />
          </div>
          <br />
          <div>
            <h3>Link WhatsApp Baru</h3>
            <input type="text" />
          </div>
          <br />
          <br />
          <button className={style.buttonHijau}>SIMPAN</button>
        </form>
      </div>
      <>
        <ConfirmationModal show={modalOpen} close={() => setModalOpen(false)} onClick={closed} />
      </>
      <>
        <BerhasilSubmit show={modalOpen} close={() => setModalOpen(false)} onClick={closed} />
      </>
    </div>
  );
};

export default ManageWa;
