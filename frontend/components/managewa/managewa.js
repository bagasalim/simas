import { useState } from "react";
import style from "./managewa.module.scss";
import { Button, Modal, ModalBody, ModalFooter } from "reactstrap";

const ManageWa = () => {
  const [modalOpen, setModalOpen] = useState(false);

  return (
    <div className={style.wa}>
      <h1>Manage Link WhatsApp</h1>
      <div className={style.inputbox} style={{ borderRadius: "16px" }}>
        <form>
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
          <button className={style.buttonHijau} onClick={() => setModalOpen(true)}>
            SIMPAN
          </button>
          {modalOpen ? <div>test</div> : null}
        </form>
      </div>
    </div>
  );
};

export default ManageWa;
