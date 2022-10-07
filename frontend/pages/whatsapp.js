import {useState, useEffect} from 'react'
import React from 'react';

export default function Whatsapp() {
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);
  useEffect(() => {
    doLogin()
  }, []);
  async function doLogin(e){
    const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/getlink?linktype=WA`, {
        method: 'GET',
        headers:{
            //Authorization : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFkbWluMiIsIk5hbWUiOiJhZG1pbiIsIlJvbGUiOjAsImV4cCI6MTY2NDk2MjQ1NH0.XViII96_Vvr6DGMSCapSVOxOq0qkkUZnCtMVEGaTiKk"
            Authorization : localStorage.getItem("token")
        },
      });
      const data = await res.json();
      setData(data);
      console.log(data);

      
}
return(
<div>
{/* {
!data ? <div>Loading...</div>
:(
data.data.linkvalue
)} */}
<a href={!data?"":data.data.linkvalue}><button>Continue to Whatsapp</button></a>
</div>
)

}