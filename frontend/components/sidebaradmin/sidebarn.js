import style from "./sidebarn.module.scss";
import Image from "next/image";
import Logo from "../../public/logo.png";
import HomeIcon from "@mui/icons-material/Home";
import AccountBoxIcon from "@mui/icons-material/AccountBox";
import InfoIcon from "@mui/icons-material/Info";
import LogoutIcon from "@mui/icons-material/Logout";
import Router from 'next/router';
const Sidebar = () => {
  const logout = ()=>{
      localStorage.removeItem("token")
      localStorage.removeItem("user")
      console.log(localStorage.getItem('token'))
      Router.replace('/loginForm');
  }
  return (
    <div className={style.sidebar}>
      <div className={style.top}>
        <Image src={Logo} />
      </div>
      <hr />
      <div className={style.center}>
        <ul>
          <p className={style.title}>MAIN MENU</p>
          <li>
            <HomeIcon className={style.icon} />
            <span>Halaman Utama</span>
          </li>
          <li>
            <AccountBoxIcon className={style.icon} />
            <span>Akun CS</span>
          </li>
          <li>
            <a href="https://mangadex.org">
              <InfoIcon className={style.icon} />
              <span>Informasi</span>
            </a>
          </li>
          <li>
            <a href="#" onClick={logout}>
              <LogoutIcon className={style.icon} />
              <span>Keluar</span>
            </a>
          </li>
        </ul>
      </div>
    </div>
  );
};

export default Sidebar;
