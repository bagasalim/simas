import Head from "next/head";
import Image from "next/image";
import styles from "../../../styles/Home.module.css";
import logo from "../../../public/logo.png";
import Sidebar from "./sidebar";

export default function home() {
  return <Sidebar />;
}
