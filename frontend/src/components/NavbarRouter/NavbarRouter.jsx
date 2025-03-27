import {Home} from '../Home';
import {AddIdea} from '../AddIdea';
import {Ideas} from '../Ideas';




import {Routes, Route} from "react-router-dom";

export default function NavbarRouter(){
    return (
      <Routes>
        <Route path="/" element={<Home/>}/>
        <Route path="/AddIdea" element={<AddIdea/>} />
        <Route path="/Ideas" element={<Ideas/>} />
      </Routes>
    )
}