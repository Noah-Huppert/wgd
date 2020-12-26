import React from "https://dev.jspm.io/react@17.0.0";
import ReactDOM from "https://dev.jspm.io/react-dom@17.0.0";

import App from "./components/App.tsx";

(ReactDOM as any).hydrate(<App />, document.getElementById("app"));
