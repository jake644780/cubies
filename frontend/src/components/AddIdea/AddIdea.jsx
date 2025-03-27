import styles from "./style.module.scss";

export default function AddIdea(){
    return (
    <div className={styles.content}>
        <form className={styles.container}>
            <div className={styles.title}>submit a new idea</div>
            <label htmlFor="category">category</label>
            <input type="text" id="category" placeholder="microservices, ml, etc." />
            <label htmlFor="content">content</label>
            <textarea id="content"></textarea>
            <button type="submit">submit</button>
        </form>
    </div>
    )
}