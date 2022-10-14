import { Button, Modal, ModalBody, ModalFooter } from "reactstrap";
import style from "./modalDetailCart.module.scss";
import { useState, useEffect } from "react";
import Image from "next/future/image";
import cslaki from "../../public/promotion.jpg";
const ConfirmationDetailCart = (props) => {
    
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);
    return (
      <>
      <div className={style.container}>
        <Modal className={style.container} isOpen={props.show} cancel={props.close}>
          <div className="modal-header" style={{ backgroundColor: "white" }}>
            <h5 className="modal-title" id="exampleModalLabel">
              Serunya Oktober<br></br>
              <h6>Deskripsi <br></br>
              <h6>Ayo Belanja sekarang dapatkan diskon 50%</h6>
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
  