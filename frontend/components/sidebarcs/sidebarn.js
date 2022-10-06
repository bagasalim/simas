import style from "./sidebarn.module.scss";
import Image from "next/image";
import Logo from "../../public/logo.png";
import HomeIcon from "@mui/icons-material/Home";
import ChatIcon from "@mui/icons-material/Chat";
import CallIcon from "@mui/icons-material/Call";
import VideocamIcon from "@mui/icons-material/Videocam";
import LogoutIcon from "@mui/icons-material/Logout";
import WhatsAppIcon from "@mui/icons-material/WhatsApp";
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
            <ChatIcon className={style.icon} />
            <span>Manage Chat</span>
          </li>
          <li>
            <a href="https://mangadex.org">
              <CallIcon className={style.icon} />
              <span>Manage Call</span>
            </a>
          </li>
          <li>
            <a href="https://mangadex.org">
              <VideocamIcon className={style.icon} />
              <span>Manage Zoom</span>
            </a>
          </li>
          <li>
            <a href="https://mangadex.org">
              <WhatsAppIcon className={style.icon} />
              <span>Manage WA</span>
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
