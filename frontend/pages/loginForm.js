import { useEffect } from "react";
import styles from "../styles/Login.module.css";

import Router from "next/router";
import Image from "next/image";
export default function LoginForm() {
  const current = new Date();
  const month = '0' + (current.getMonth() + 1);
  const day = '0' + current.getDate();
  const hours = '0' + current.getHours();
  const minutes = '0' + current.getMinutes();
  const second = '0' + current.getSeconds();
  const dateRequest = current.getFullYear() + '-'
              + month.substring(month.length-2,month.length) + '-'
              + day.substring(day.length-2,day.length) + 'T'
              + hours.substring(hours.length-2,hours.length) + ':'
              + minutes.substring(minutes.length-2,minutes.length) + ':'
              + second.substring(second.length-2,second.length) + '.000Z';

  useEffect(() => {
    let user = localStorage.getItem("user");
    let token = localStorage.getItem("token");
    if (user != null && token != null) {
      user = JSON.parse(user);
      if (user.role == 1) {
        Router.replace("/project/admin");
        return;
      } else if (user.role == 2) {
        Router.replace("/project/customerservice");
        return;
      }
    }
    localStorage.removeItem("user");
    localStorage.removeItem("token");
  }, []);
  async function doLogin(e) {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      username: formData.get("username"),
      password: formData.get("password"),
    };

    const res = await fetch(`${process.env.NEXT_PUBLIC_URL}login`, {
      method: "POST",
      header: {
        Authorization: "token",
      },
      body: JSON.stringify(body),
    });
    const data = await res.json();

    if (data.token) {
      localStorage.setItem("token", data.token);

      localStorage.setItem("user", JSON.stringify(data.data));
      if (data.data.role == 1) {
        postLastLogin();
        Router.replace("/project/admin");
      } else if (data.data.role == 2) {
        postLastLogin();
        Router.replace("/project/customerservice");
      } else {
        console.log("Tidak ada Role");
      }
    } else {
      alert("Username Password salah");
    }
  }

  async function postLastLogin() {
    const data_user  = localStorage.getItem('user');
    const newData = JSON.parse(data_user);
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}updatelastlogin`, {
        method: "POST",
        headers: {
          Authorization: localStorage.getItem("token"),
        },
        body: JSON.stringify({
          username: newData.username,
          lastlogin: dateRequest,
        }),
      });
      if(res.status != 200){
        throw "gagal mendapatkan response update last login"();
      }
    } catch (error) {
      alert("Update Last Login Gagal");
    }
  }

  return (
    <div>
      <div className={styles.background}>
        <div className={styles.shape} />
        <div className={styles.shape} />
      </div>
      <div className="text-center mt-4">
        <Image src={"/logo.png"} width="378" height="100" alt="Logo" />

        <h2>Simas Contact dan Info</h2>
      </div>
      <form onSubmit={doLogin} id="formid" className={styles.form}>
        <h3>Masuk</h3>

        <label htmlFor="username" className={styles.label}>
          Username
        </label>
        <input type="text" placeholder="Masukan Username" id="username" name="username" className={styles.input} />

        <label htmlFor="password">Password</label>
        <input type="password" placeholder="Password" id="password" name="password" className={styles.input} />

        <a href="#" style={{ marginLeft: "70%", color: "#4A8CFF" }}>
          Lupa Kata Sandi ?
        </a>
        <button className={styles.button}>Masuk</button>
      </form>
    </div>
    
  );
}
