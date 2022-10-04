import Footer from "./footer";
import styles from "../styles/Home.module.css";

const Layout = ({ children }) => {
  return (
    <div className="content">
      {children}
      <div className={styles.footer}>
        <Footer />
      </div>
    </div>
  );
};

export default Layout;
