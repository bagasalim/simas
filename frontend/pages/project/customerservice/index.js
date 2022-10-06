import Head from "next/head";
import Image from "next/image";
import styles from "../../../styles/Home.module.css";
import logo from "../../../public/logo.png";
import Sidebar from "../../../components/sidebarcs/sidebarn";
import style from "./index.module.scss";
import HalamanUtama from "../../../components/halamanutamacs/halamanutama";
import ManageWa from "../../../components/managewa/managewa";

export default function home() {
  return (
    <div className={style.home}>
      <Sidebar />
      <div className={style.homeContainer}>
        <div className={style.content}>
          <ManageWa />
        </div>
      </div>
    </div>
  );
}
