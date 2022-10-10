import React from "react";
import style from "./ManageWa.module.scss";
import ConfirmationModal from "../modals/modalwadanzoom";
import { useState, useEffect } from "react";
import { useRouter } from "next/route";

const ManageWa = () => {
  const router = useRouter();
  const [data, setData] = useState(null);
  const [newLink, setNewLink] = useState("");
  const [error, setError] = useState(null);
  useEffect(() => {
    getWa();
  }, []);

  const [modalOpen, setModalOpen] = React.useState(false);
  const [body, setBodyData] = React.useState("");

  const onSubmit = async (e) => {
    const dataform = {
      newlink: newLink,
    };
    setBodyData(dataform);
    setModalOpen(true);
  };
  async function getWa(e) {
    const res = await fetch(`${process.env.NEXT_PUBLIC_URL}/getlink?linktype=Wa`, {
      method: "GET",
      headers: {
        Authorization: localStorage.getItem("token"),
      },
    });
    const data = await res.json();
    setData(data);
    console.log(data);
  }
  async function putWa() {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    setModalOpen(false);

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/updatelink?linktype=Zoom`, {
        method: "PUT",
        headers: {
          Authorization: localStorage.getItem("token"),
        },
        body: JSON.stringify({ linkvalue: body.newlink, UpdatedBy: "system" }),
      });
      const resData = await res.json();
      const d = { ...data };
      d.data.linkvalue = body.newlink;
      console.log(d);
      setData(d);
      setData(data);
      alert("Update Sukses");
    } catch (error) {
      setError(error);
    }
  }

  return (
    <div className={style.zoom}>
      <h1>Manage Link WhatsApp</h1>
      <div className={style.inputbox}>
        <form
          onSubmit={(e) => {
            e.preventDefault();
          }}
        >
          <div>
            <h3>Link WhatsApp Lama</h3>
            <input className={style.readonly} type="text" placeholder="jabj" readOnly disabled="true" />
          </div>
          <br />
          <div>
            <h3>Link WhatsApp Baru</h3>
            <input type="text" name="newlink" required value={newLink} onChange={(e) => setNewLink(e.target.value)} />
          </div>
          <br />
          <br />
          <button className={style.buttonHijau} onClick={onSubmit}>
            SIMPAN
          </button>
        </form>
      </div>
      <ConfirmationModal show={modalOpen} close={() => setModalOpen(false)} linktype={"Wa"} data={body} response={putWa} />;
    </div>
  );
};

export default ManageWa;
