import {  useEffect } from "react";
import styles from "../styles/Login.module.css";

import Router from "next/router";
import Image from 'next/image'
export default function LoginForm(){
  useEffect(()=>{
    let user=localStorage.getItem("user")
    let token=localStorage.getItem("token")
    if(user !=null && token!=null){
      user = JSON.parse(user)
      if(user.role == 1){
        Router.replace("/project/admin");
        return
      }else if(user.role ==  2){
        Router.replace("/project/customerservice");
        return
      }
      
    }
    localStorage.removeItem("user")
    localStorage.removeItem("token")

  },[])
  async function doLogin(e) {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      username: formData.get("username"),
      password: formData.get("password"),
    };
    const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/login`, {
      method: "POST",
      header: {
        Authorization: "token",
      },
      body: JSON.stringify(body),
    });
    const data = await res.json();

    if (data.token) {
      localStorage.setItem("token", data.token);
      localStorage.setItem("user",JSON.stringify(data.data))
      if (data.data.role == 1) {
        Router.replace("/project/admin");
      } else if (data.data.role == 2) {
        Router.replace("/project/customerservice");
      } else {
        console.log("Tidak ada Role");
      }
    } else {
      alert("Username Password salah");
    }
  }

  return (
    <div>
        <div className={styles.background}>
          <div className={styles.shape} />
          <div className={styles.shape}  />
        </div>
        <div className="text-center mt-5">
          <div>
          <Image src={"/logo.png"} width="20%" alt="Logo"/>
            {/* <img src="/logo.png" style={{width:"20%"}}/> */}
    
          </div>
          <h2>Simas Contact dan Info</h2>
        </div>
        <form onSubmit={doLogin} id="formid" className={styles.form}>
          <h3>Masuk</h3>
        
            <label htmlFor="username" className={styles.label}>Username</label>
            <input type="text" placeholder="Masukan Username" id="username" name ="username" className={styles.input} />
        
            <label htmlFor="password">Password</label>
            <input type="password" placeholder="Password" id="password" name ="password" className={styles.input}/>
        
          <a href="#" style={{ marginLeft: "70%", color: "#4A8CFF" }}>
            Lupa Kata Sandi ?
          </a>
          <button className={styles.button}>Masuk</button>
        </form>
    </div>
  );
}
