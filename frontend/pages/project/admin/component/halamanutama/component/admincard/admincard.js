import style from "./admincard.module.scss";
import Foto from "../assets/fotocs.jpg";
import Image from "next/image";

const AdminCard = () => {
  const names = [
    "Afif Fatur Rahman",
    "Andre Diwa",
    "Bagas Alim santoso",
    "Braike Rema Alfian",
    "Calvin Yonathan",
    "Cindy Surjawan",
  ];
  return (
    <div className={style.admincard}>
      {names.map((name) => (
        <div className={style.item}>
          <div className={style.foto}>
            <Image src={Foto} />
          </div>
          <p>{name}</p>
        </div>
      ))}
    </div>
  );
};

export default AdminCard;
