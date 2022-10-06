import React, {useState, useEffect} from 'react'

//import nookies from 'nookies';
import Router from 'next/router';


export default function loginForm() {
  const [loading, setLoading ]= useState(false)
  useEffect(()=>{
    const token = localStorage.getItem("token")
    let user = localStorage.getItem("user")
    if(token != null && user != null ){
      user = JSON.parse(user)
      if(user.role == 2){
        console.log("redirect")
        Router.replace('/project/customerservice');
        return
      }else if(user.role == 1){
        Router.replace('/project/admin');
        return
      }
    }
    localStorage.removeItem("token")
    localStorage.removeItem("user")
   
  },[])
    async function doLogin(e){
        if(loading){ return }
        setLoading(true)
        e.preventDefault(); 
        const formData = new FormData(e.currentTarget);
        const body = {
            username: formData.get("username"),
            password: formData.get("password")
         }
        if(body.username == "" || body.password == ""){
            alert("Masukkan email dan password")
            return 
        }
        try{
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/login`, {
                method: 'POST',
                
                body: JSON.stringify(body)
            });
            const data = await res.json();
            
            
            if(data.token){
                localStorage.setItem("token", data.token);
                localStorage.setItem("user", JSON.stringify(data.data))
                if(data.data.role == 1){
                    
                    Router.replace('/project/admin');
                    return
                }else if(data.data.role == 2){
                    Router.replace('/project/customerservice');
                    return
                }else{
                    console.log("Tidak ada Role");
                }
            }else{
                alert("Username Password salah");
            }
        }catch(e){
            alert("server sedang bermasalah")
        }
        setLoading(false)
        
           
        
          
    }

    
    return (
        <>
        <link rel="preconnect" href="https://fonts.gstatic.com" />
        <link
          rel="stylesheet"
          href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.4/css/all.min.css"
        />
        <link
          href="https://fonts.googleapis.com/css2?family=Poppins:wght@300;500;600&display=swap"
          rel="stylesheet"
        />
        {/*Stylesheet*/}
        <style
          media="screen"
          dangerouslySetInnerHTML={{
            __html:
              "\n      *,\n*:before,\n*:after{\n    padding: 0;\n    margin: 0;\n    box-sizing: border-box;\n}\nbody{\n    background-color: #ffffff;\n}\n\n.logo-login{\n    position: absolute;\n    left: 44%;\n    top: 5%;   \n}\n\n.text-login{\n    position: absolute; \n    left: 38%;\n    top: 18%;\n    font-family: 'Montserrat';\n    font-style: normal;\n    font-weight: 600;\n    font-size: 35px;\n    line-height: 49px;  \n    color: #000000;\n    backdrop-filter: blur(10px);\n}\n\nform{\n    width: 587px;\n    height: 500px;\n    background-color: rgba(255,255,255,0.13);\n    position: absolute;\n    transform: translate(-50%,-50%);\n    top: 50%;\n    left: 50%;\n    border-radius: 10px;\n    backdrop-filter: blur(10px);\n    border: 2px solid rgba(255,255,255,0.1);\n    box-shadow: 0 0 40px rgba(8,7,16,0.6);\n    padding: 50px 35px;\n    font-family: 'Montserrat';\n    font-style: normal;\n}\nform *{\n    font-family: 'Poppins',sans-serif;\n    color: #000000;\n    letter-spacing: 0.5px;\n    outline: none;\n    border: none;\n}\nform h3{\n    font-size: 32px;\n    font-weight: 500;\n    line-height: 42px;\n    text-align: center;\n}\n\nlabel{\n    display: block;\n    margin-top: 30px;\n    font-size: 16px;\n    font-weight: 500;\n}\ninput{\n    display: block;\n    height: 50px;\n    width: 100%;\n    background-color: rgba(1, 1, 1, 0.07);\n    border-radius: 3px;\n    padding: 0 10px;\n    margin-top: 8px;\n    font-size: 14px;\n    font-weight: 300;\n}\n::placeholder{\n    color: #020202;\n}\nbutton{\n    margin-top: 50px;\n    width: 100%;\n    background-color: #36506A;\n    color: #ffffff;\n    padding: 15px 0;\n    font-size: 18px;\n    font-weight: 600;\n    border-radius: 16px;\n    cursor: pointer;\n}\n\n.background{\n    width: 430px;\n    height: 520px;\n    position: absolute;\n    transform: translate(-50%,-50%);\n    left: 43%;\n    top: 50%;\n}\n.background .shape{\n    height: 200px;\n    width: 200px;\n    position: absolute;\n    border-radius: 50%;\n}\n.shape:first-child{\n    background: linear-gradient(\n        #081A2C,\n        #1D3144,\n        #36506A\n    );\n    /* right: -320px;\n    bottom: -80px; */\n    /* left: -80px; */\n    right: -76%;\n    top: -80px;\n    \n}\n.shape:last-child{\n    background: linear-gradient(\n        #081A2C,\n        #1D3144,\n        #36506A\n    );\n    /* right: -900px;\n    top: -80px; */\n    right: 70%;\n    bottom: -80px;\n    \n}\n\n"
          }}
        />
      
        <div className="text-login">
        <div className="logo-login">
          {/* <Image src={logo} style={{width : "20px", position : "relative"}}/> */}
        </div>
            <h3>Simas Contact &amp; Info</h3>
      
        </div>
        <div className="background">
            <div className="shape" />
            <div className="shape" />
        </div>
        <form onSubmit={doLogin}>
            <h3>Masuk</h3>
            <label htmlFor="username">Username</label>
            <input type="text" placeholder="Masukan Username" id="username" name ="username"/>
            <hr />
            <label htmlFor="password">Password</label>
            <input type="password" placeholder="Password" id="password" name ="password" />
            <a href="#" style={{ marginLeft: "70%", color: "#4A8CFF" }}>
              Lupa Kata Sandi ?
            </a>
            <button >{loading?"Please wait":"Masuk"}</button>
        </form>
        </>
    );
}