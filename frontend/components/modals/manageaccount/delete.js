import { Button, Modal, ModalBody } from "reactstrap";
import style from "./delete.module.scss";

const ModalDelete = (props) => {
  return (
    <>
      <Modal isOpen={props.show} cancel={props.close} toggle={props.close}>
        <div className="modal-header" style={{ backgroundColor: "#7E8A97" }}>
          <h5
            className="modal-title"
            id="exampleModalLabel"
            style={{ color: "white" }}
          >
            Konfirmasi Perubahan
          </h5>
          <Button
            aria-label="Close"
            className=" close"
            type="button"
            onClick={props.close}
          >
            <span aria-hidden={true}>×</span>
          </Button>
        </div>
        <ModalBody>
          <div className={style.body}>
            Apakah kamu yakin ingin menghapus {props.data.username}?
          </div>
          <div className={style.tombol}>
            <Button className={style.setuju} type="button">
              YA
            </Button>
            <Button className={style.tidak} type="button" onClick={props.close}>
              Tidak
            </Button>
          </div>
        </ModalBody>
      </Modal>
    </>
  );
};

export default ModalDelete;