import {useState, useEffect} from 'react';
import Header from "~/components/Header";
import {Router} from 'next/router';
import Image from "next/future/image";
import desktopzoom from "../../public/desktopzoom.png";
import UserFooter from "../../components/userfooter";
import styles from "./Zoom.module.scss";

export default function Zoom(){
    const [link, setLink] = useState("")
    const [loading, setLoading] = useState(false)
    const openZoom =()=>{
        if(link == ""){
            alert("link zoom sedang bermasalah, silahkan merefresh ulang")
            return
        }
    }
    const getLinkZoom = async ()=>{
        
        try{
            const res = await fetch(`${process.env.NEXT_PUBLIC_URL}get-link/zoom`)
            if (res.status != 200){
                throw new "gagal mendapatkan pesan"
            }
            const data = await res.json()
            if(!data.data){
                throw new "gagal mendapatkan data"
            }
            setLink(data.data)
            setLoading(false)
            return true
        }catch(e){
            if(typeof e==="string"){
                alert("Link Zoom tidak ada, silahkan merefresh ulang")
            }
            return false
        }
        
    }
    useEffect(()=>{
        getLinkZoom()
    },[])
    return(
        <div>
            <Header/>
          <div className="col-6 border-end-0 border-3">
          </div>
            <div className="container-fluid mt-5  ">
                <h2 className="ms-3 fw-bold">Layanan CS - Video Zoom</h2>
                <div className="row mt-4">
                    <div className="col-6">
                        <Image src={desktopzoom} width={500} height={400} alt={"desktopzoom"} />
                    </div>
                    <div className="col-6">
                        <h3 className="fw-bold">Proses Video Zoom</h3>
                        <div style={{maxWidth:"80%"}} className="mt-4">
                            <ol>
                                <li>Silahkan join room zoom menggunakan nama dan email yang sesuai diisikan di data diri, apabila tidak sesuai maka tidak akan diproses </li>
                                <li>Tunggu terlebih dahulu di ruang utama hingga customer service mengundangan ke breakout room</li>
                                <li>Silahkan join breakout room </li>
                            </ol>
                        </div>
                       
                    </div>
                    
                </div>
            </div>
        <br />
        <hr />
        <br />
        <div class={styles.container}>
        <form id="formZoom">
            <h3 className={styles.titleForm}>Isi Data Diri</h3>

            <label htmlFor="namaZoom" className={styles.label}>Nama</label>
            <input type="text" placeholder="Masukan Nama" id="namaZoom" name="namaZoom" className={styles.input} />

            <label htmlFor="emailZoom" className={styles.label}>Email</label>
            <input type="text" placeholder="Masukkan Email" id="emailZoom" name="emailZoom" className={styles.input} />

            <label htmlFor="kategoriZoom" className={styles.label}>Kategori</label>
            <select id="kategoriZoom" name="kategoriZoom" className={styles.input}>
                <option value="perbankan">Perbankan</option>
                <option value="kartuKredit">Kartu Kredit</option>
                <option value="digitalLoan">Digital Loan</option>
                <option value="merchantQR">Merchant QRIS</option>
                <option value="pin">Mengganti PIN Channel</option>
                <option value="custcare">Berbicara dengan CustCare</option>
            </select>
            
            {
                            loading ?
                            ( <div>
                                Please Wait
                            </div> ):(
                                //  <button className="btn text-white" style={{background:"#2D8CFF"}} onClick={openZoom}>Melanjutkan ke Zoom</button>
                                    <button className={styles.button} onClick={openZoom}>Melanjutkan ke Zoom</button>
                            )
            }
        </form>
        </div>
        <UserFooter />
        </div>
    )
}