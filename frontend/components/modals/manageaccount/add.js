import { Modal, ModalBody } from "reactstrap";
import style from "./add.module.scss";
import { useState } from "react";

const ModalAdd = (props) => {
  const [passwordShown, setPasswordShown] = useState(false);
  const [passwordConfirmShown, setPasswordConfirmShown] = useState(false);

  const addUser = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      //MASUKKAN BODY
      // nama: formData.get("username"),
      // email: formData.get("nama"),
      // kategori: formData.get("role"),
      // keterangan: formData.get("email"),
      // keterangan: formData.get("katasandi"),
    };
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_URL}create-account`, {
        method: "POST",
        body: JSON.stringify(body),
      });
      if (res.status != 200) {
        throw "gagal membuat user CS"();
      }
    } catch (e) {
      if (typeof e === "string") {
        alert("Gagal membuat user CS, silahkan refresh ulang");
      }
      return false;
    }
  };

  return (
    <>
      <Modal isOpen={props.show} toggle={props.close}>
        <ModalBody>
          <div style={{ padding: "20px" }}>
            <h4
              style={{
                textAlign: "center",
                paddingTop: "10px",
                paddingBottom: "10px",
              }}
            >
              Form Tambah Data
            </h4>
            <br />
            <form>
              <div className="form-group" style={{ marginBottom: "20px" }}>
                <label for="exampleInputEmail1">Username</label>
                <input
                  type="text"
                  className="form-control"
                  name="username"
                  aria-describedby="emailHelp"
                  placeholder="Masukkan username"
                  style={{
                    boxShadow: `rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px`,
                  }}
                />
              </div>
              <div className="form-group" style={{ marginBottom: "20px" }}>
                <label>Nama</label>
                <input
                  type="text"
                  className="form-control"
                  name="nama"
                  placeholder="Masukkan nama"
                  style={{
                    boxShadow: `rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px`,
                  }}
                />
              </div>
              <div className="form-group" style={{ marginBottom: "20px" }}>
                <label>Posisi</label>
                <select name="role" class="form-control">
                  <option value="2">Customer Service</option>
                </select>
              </div>
              <div className="form-group" style={{ marginBottom: "20px" }}>
                <label>Email</label>
                <input
                  type="email"
                  className="form-control"
                  name="email"
                  placeholder="Masukkan email"
                  style={{
                    boxShadow: `rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px`,
                  }}
                />
              </div>
              <div className="form-group" style={{ marginBottom: "10px" }}>
                <label for="exampleInputPassword1">Kata Sandi</label>
                <input
                  type={passwordShown ? "text" : "password"}
                  className="form-control"
                  name="katasandi"
                  placeholder="Masukkan kata sandi"
                  style={{
                    boxShadow: `rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px`,
                  }}
                />
                <input type="checkbox" class="form-check-input" onClick={() => setPasswordShown(!passwordShown)} />
                <label class="form-check-label" for="exampleCheck1" style={{ fontSize: "14px", paddingLeft: "5px" }}>
                  Show Password
                </label>
              </div>
              <div className="form-group" style={{ marginBottom: "10px" }}>
                <label for="exampleInputPassword1">Konfirmasi Kata Sandi</label>
                <input
                  type={passwordConfirmShown ? "text" : "password"}
                  className="form-control"
                  name="konfirmasi"
                  placeholder="Masukkan kata sandi kembali"
                  style={{
                    boxShadow: `rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px`,
                  }}
                />
                <input type="checkbox" class="form-check-input" onClick={() => setPasswordConfirmShown(!passwordConfirmShown)} />
                <label class="form-check-label" for="exampleCheck1" style={{ fontSize: "14px", paddingLeft: "5px" }}>
                  Show Password
                </label>
              </div>
              <button type="submit" className={style.buttonHijau} style={{ marginTop: "20px" }} onSubmit={addUser}>
                Kirim
              </button>
            </form>
          </div>
        </ModalBody>
      </Modal>
    </>
  );
};

export default ModalAdd;
