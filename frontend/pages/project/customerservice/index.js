import Head from "next/head";
import Image from "next/image";
import styles from "../../../styles/Home.module.css";
import logo from "../../../public/logo.png";
import Sidebar from "./sidebar";
import Halamanutama from "./content/halamanutama";

export default function home() {
  return (
    <div style={{ height: "100%", width: "100%" }} className="container">
      <div class="row">
        <div class="col-sm-3" style={{ padding: 0 }}>
          <Sidebar />
        </div>
        <div class="col-sm-9">
          <Halamanutama />
        </div>
      </div>
    </div>
  );
}
