import { useState } from "react";
import style from "./managezoom.module.scss";
import { Button, Modal, ModalBody, ModalFooter } from "reactstrap";

const ManageZoom = () => {
  const [modalOpen, setModalOpen] = useState(false);

  return (
    <div className={style.zoom}>
      <h1>Manage link zoom</h1>
      <div className={style.inputbox}>
        <form>
          <div>
            <h3>Link ZOOM Lama</h3>
            <input
              className={style.readonly}
              type="text"
              placeholder="https://zoom.us/w/99582712162?tk=_ILvh4FKnvxojs9q0ShiqEJsUyaIf4eE7qlYPIpmBQI.DQMAAAAXL5eZYhZFTWZ1dVFqQ1NBMjB5THVjMjBHakh3AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA&pwd=SzRUOFNIVldlRkR6SlFpc004OUs1Zz09"
              readOnly
            />
          </div>
          <br />
          <div>
            <h3>Link ZOOM Baru</h3>
            <input type="text" />
          </div>
          <br />
          <br />
        </form>
      </div>
    </div>
  );
};

export default ManageZoom;
