import React from "react";
import style from "./managezoom.module.scss";
import ConfirmationModal from "../modals/modalwadanzoom/modalwadanzoom";

const ManageZoom = () => {
  const [modalOpen, setModalOpen] = React.useState(false);
  const [body, setBodyData] = React.useState("");

  const onSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const dataform = {
      newlink: formData.get("newlink"),
    };
    body = setBodyData(dataform);
    modalOpen = setModalOpen(true);
  };

  return (
    <div className={style.zoom}>
      <h1>Manage link zoom</h1>
      <div className={style.inputbox}>
        <form onSubmit={onSubmit}>
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
            <input type="text" name="newlink" required />
          </div>
          <br />
          <br />
          <button className={style.buttonHijau}>SIMPAN</button>
        </form>
      </div>
      <ConfirmationModal
        show={modalOpen}
        close={() => setModalOpen(false)}
        linktype={"Zoom"}
        data={body}
      />
      ;
    </div>
  );
};

export default ManageZoom;
