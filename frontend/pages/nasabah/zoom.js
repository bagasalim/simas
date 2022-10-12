import { useState, useEffect } from "react";
import Header from "~/components/Header";
import { Router, useRouter } from "next/router";
import { unstable_renderSubtreeIntoContainer } from "react-dom";

export default function Zoom() {
  const [link, setLink] = useState(false);
  const [loading, setLoading] = useState(false);
  const router = useRouter();

  const openZoom = () => {
    if (link == "") {
      alert("link zoom sedang bermasalah, silahkan merefresh ulang");
      return;
    }
    router.push(link);
  };

  const getLinkZoom = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}get-link/Zoom`);
      if (res.status != 200) {
        throw new "gagal mendapatkan pesan"();
      }
      const data = await res.json();
      if (!data.data) {
        throw new "gagal mendapatkan data"();
      }
      setLink(data.data.linkvalue);
      setLoading(false);
      return true;
    } catch (e) {
      if (typeof e === "string") {
        alert("Link Zoom tidak ada, silahkan merefresh ulang");
      }
      return false;
    }
  };

  const postDataZoom = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      nama: formData.get("namaZoom"),
      email: formData.get("emailZoom"),
      kategori: formData.get("kategoriZoom"),
      keterangan: formData.get("keluhanZoom"),
    };
    try {
      await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
        method: "POST",
        body: JSON.stringify(body),
      });
    } catch (e) {
      if (typeof e === "string") {
        alert("Gagal menginputkan form data diri, silahkan refresh ulang");
      }
      return false;
    }
  };

  useEffect(() => {
    getLinkZoom();
  }, []);

  return (
    <div>
      <Header />
      <div className="container-fluid mt-5  ">
        <h2 className="ms-3 fw-bold">Layanan CS - Video Zoom</h2>
        <div className="row mt-4">
          <div className="col-6"></div>
          <div className="col-6">
            <h3 className="fw-bold">Proses Video Zoom</h3>
            <div style={{ maxWidth: "80%" }} className="mt-4">
              <ol>
                <li>
                  Silahkan join room zoom menggunakan nama dan email yang sesuai
                  diisikan di data diri, apabila tidak sesuai maka tidak akan
                  diproses{" "}
                </li>
                <li>
                  Tunggu terlebih dahulu di ruang utama hingga customer service
                  mengundangan ke breakout room
                </li>
                <li>Silahkan join breakout room </li>
              </ol>
            </div>
            {loading ? (
              <div>Please Wait</div>
            ) : (
              <button
                className="btn text-white"
                style={{ background: "darkturquoise" }}
                onClick={openZoom}
              >
                Open Link Zoom
              </button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
