import {  Modal, ModalBody } from "reactstrap";
import style from "./modalDetailCart.module.scss";
import Image from "next/future/image";
import cslaki from "../../public/covid.jpg";
const ConfirmationDetailCart = (props) => {
    return (
      <>
      <div className={style.container}>
        <Modal className={style.container} isOpen={props.show} cancel={props.close}>
          <div className="modal-header" style={{ backgroundColor: "white" }}>
            <h5 className="modal-title" id="exampleModalLabel">
              Simas Covid<br></br>
              <h6>Deskripsi <br></br>
              <h6>Premi mulai dari Rp. 50.000</h6>
              </h6>
            </h5>
            <button aria-label="Close" className=" close" type="button" onClick={props.close}>
              <span aria-hidden={true}>×</span>
            </button>
          </div>
          <ModalBody>
            <div className={style.body}> 
            <Image className={style.cslaki} src={cslaki} alt="cslaki" />
            </div>
            <div>
              Syarat Dan Ketentuan
            </div>
            <div>
              1.Gatau ngisi apaan pusing mau nulis apaan otak ngeblank aja tiba-tiba
            </div>
            <div>
            2.Gatau ngisi apaan pusing mau nulis apaan otak ngeblank aja tiba-tiba
            </div>
            <div>
            3.Gatau ngisi apaan pusing mau nulis apaan otak ngeblank aja tiba-tiba
            </div>
          </ModalBody>
         </Modal>
         </div>
      </>
    );
  };
  
  export default ConfirmationDetailCart;
  