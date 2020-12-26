import "./App.css";

import { React, Bootstrap } from "./deps.ts";
import Logo from "./Logo.tsx";

function App() {
    return (
	   <div className="App">
		  <Bootstrap.Button variant="primary">Button</Bootstrap.Button>
	   </div>
    );
}

export default App;
