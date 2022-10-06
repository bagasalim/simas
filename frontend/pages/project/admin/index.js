import Head from "next/head";
import Image from "next/image";
import styles from "../../../styles/Home.module.css";
import logo from "../../../public/logo.png";
import Sidebar from "../../../components/sidebaradmin/sidebarn";
import style from "./index.module.scss";
import HalamanUtama from "../../../components/halamanutamaadmin/halamanutama";
import { useEffect, useState } from "react";
import Router from 'next/router';

export default function Index(){
  const [loading, setLoading] = useState(true)
  useEffect(()=>{
    const token = localStorage.getItem("token")
    let user = localStorage.getItem("user")
    if(token == null || user == null ){
      console.log("logout")
      Router.replace('/loginForm');
      return
    }
    user = JSON.parse(user)
    if(user.role != 1 ){
      if(user.role == 2){
        console.log("redirect")
        Router.replace('/project/customerservice');
        return
      }
      console.log("load", user)
      Router.replace('/loginForm');
      return
    }
    setLoading(false)
    console.log(token, user)
  },[])
  
  return (
    <div className={style.home}>
      <Sidebar />
      <div className={style.homeContainer}>
        <div className={style.content}>
          {
            loading ?
              (<div>
                <h1>Please wait</h1>
              </div>)
              :<HalamanUtama />
          }
          
        </div>
      </div>
    </div>
  );
}


