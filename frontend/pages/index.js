import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";
import logo from "../public/logo.png";

export default function Home() {
  return (
    <div>
      Hello World
      <a href="project/customerservice" as="/cs">
        test
      </a>
    </div>
  );
}
