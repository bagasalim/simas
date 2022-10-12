import React from "react";
import Image from "next/image";
import Logo from "~/public/logo.png";
export default function Header() {
  return (
    <div>
      <div style={{ paddingLeft: "20px", paddingTop: "10px" }}>
        <Image src={Logo} width="200px" height="50px" layout="fixed" />
      </div>
      <div style={{ background: "#CC100F", width: "100%", height: "50px" }}>
        <div className="d-flex flex-row-reverse me-5 align-items-center " style={{ height: "100%" }}>
          <a href="#layananinfo" className="pe-2 text-white px-5" style={{ fontWeight: "bold", fontSize: "20px" }}>
            Pusat Informasi
          </a>
          <a href="#layanancs" className="pe-2 text-white px-5" style={{ fontWeight: "bold", fontSize: "20px" }}>
            Layanan CS
          </a>
        </div>
      </div>
    </div>
  );
}
