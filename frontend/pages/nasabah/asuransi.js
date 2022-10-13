import Header from "../../components/Header";
import UserFooter from "../../components/userfooter";
import style from "./asuransi.module.scss";
import Image from "next/future/image";
import jumbotron from "../../public/jumbotron.png";
const Asuransi = () => {
  return (
    <div>
      <Header />
      <div>
        <Image className="{style.jumbotron}" src={jumbotron} alt="jumbotron" />
      </div>
      <div className={style.buttonpa}>
        <div>
          <button className={style.buttonpromo}>Promo</button>
        </div>
        <div>
          <button className={style.buttonasuransi}>Asuransi</button>
        </div>
      </div>
      <UserFooter />
    </div>
  );
};

export default Asuransi;
