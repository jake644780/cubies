import { useState, useEffect } from "react";
import styles from "./style.module.scss";

export default function Hero(){
    let [time, setTime] = useState(new Date().toLocaleTimeString());
    useEffect(() => {
        const interval = setInterval(() => {
          setTime(new Date().toLocaleTimeString());
        }, 1000);
    
        return () => clearInterval(interval); // Cleanup interval on unmount
      }, []);
    return(
         <div className={styles.container}>
            <div className={styles.content}>
                <div className={styles.title}>welcome back, Jake</div>
                <div className={styles.time}>current time is {time}</div>
            </div>
         </div>
    )
};
