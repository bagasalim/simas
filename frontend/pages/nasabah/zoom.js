import {useState, useEffect} from 'react'
import Header from "~/components/Header"
import {Router} from 'next/router'
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
            <div className="container-fluid mt-5  ">
                <h2 className="ms-3 fw-bold">Layanan CS - Video Zoom</h2>
                <div className="row mt-4">
                    <div className="col-6">

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
                        {
                            loading ?
                            ( <div>
                                Please Wait
                            </div> ):(
                                 <button className="btn text-white" style={{background:"darkturquoise"}} onClick={openZoom}>Open Link Zoom</button>
                            )
                        }
                       
                    </div>
                    
                </div>
                
            </div>
        </div>
    )
}