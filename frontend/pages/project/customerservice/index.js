import Head from "next/head";
import Image from "next/image";
import styles from "../../../styles/Home.module.css";
import logo from "../../../public/logo.png";
import Sidebar from "../../../components/sidebarcs/sidebarn";
import style from "./index.module.scss";
import HalamanUtama from "../../../components/halamanutamacs/halamanutama";
import ManageZoom from "../../../components/managezoom/managezoom";

export default function home() {
  return (
    <div className={style.home}>
      <Sidebar />
      <div className={style.homeContainer}>
        <div className={style.content}>
          <ManageZoom />
        </div>
      </div>
    </div>
  );
}
