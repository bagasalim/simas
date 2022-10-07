import { Button, Modal, ModalBody, ModalFooter } from "reactstrap";
import style from "./modalwadanzoom.module.scss";
import Success from "../modalsuccess/berhasilsubmit";
import React from "react";

const ConfirmationModal = (props) => {
  const [modalOpen, setModalOpen] = React.useState(false);

  const onClick = () => {
    modalOpen = setModalOpen(true);
    props.close;
  };

  return (
    <>
      <Modal modalClassName="modal" toggle={props.close} isOpen={props.show}>
        <div className="modal-header" style={{ backgroundColor: "#36506A" }}>
          <h5
            className="modal-title"
            id="exampleModalLabel"
            style={{ color: "white" }}
          >
            Konfirmasi Perubahan
          </h5>
          <button
            aria-label="Close"
            className=" close"
            type="button"
            onClick={props.close}
          >
            <span aria-hidden={true}>×</span>
          </button>
        </div>
        <ModalBody>
          <div className={style.body}>
            Apakah kamu yakin ingin mengubah link?
          </div>
        </ModalBody>
        <ModalFooter>
          <Button className={style.setuju} type="button" onClick={onClick}>
            YA
          </Button>
          <Button className={style.tidak} type="button" onClick={props.close}>
            TIDAK
          </Button>
        </ModalFooter>
      </Modal>
      <Success
        show={modalOpen}
        close={() => setModalOpen(false)}
        note={"Diubah"}
      />
    </>
  );
};

export default ConfirmationModal;
