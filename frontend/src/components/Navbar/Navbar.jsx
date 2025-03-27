import { useState } from "react";
import styles from "./style.module.scss";
import { Link } from "react-router-dom";

export default function Navbar(){
    let [activeTab, setActiveTab] = useState("/");

    return(
        <div className={styles.navbar}>
                <div className={`${styles.element} ${activeTab === "/" ? styles.active : ""}`} onClick={()=>setActiveTab("/")}>
                    <Link to="/">Home</Link>
                </div>
                <div className={`${styles.element} ${activeTab === "/AddIdea" ? styles.active : ""}`} onClick={()=>setActiveTab("/AddIdea")}>
                    <Link to="/AddIdea">addIdea</Link>
                </div>
                <div className={`${styles.element} ${activeTab === "/Ideas" ? styles.active : ""}`} onClick={()=>setActiveTab("/Ideas")}>
                    <Link to="/Ideas">Ideas</Link>
                </div>
        </div>
    );
}
