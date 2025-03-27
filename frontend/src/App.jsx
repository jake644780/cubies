import {Navbar} from './components/Navbar';
import {Footer} from './components/Footer';
import {NavbarRouter} from './components/NavbarRouter';

import { BrowserRouter as Router} from "react-router-dom";


function App() {
  return (
    <Router>
      <Navbar />
      <NavbarRouter />
      <Footer />
    </Router>
  );
}

export default App;