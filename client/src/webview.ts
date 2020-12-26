import { WebView } from "https://deno.land/x/webview/mod.ts";

const jsBundle = await Deno.readTextFile("./dist/bundle.js");

const webview = new WebView({
    title: "Rhapso",
    url: `data:text/html,\
<!DOCTYPE HTML>
<html>
    <body>
        <div id="app">
        </div>
        <script type="text/javascript">${jsBundle}</script>
    </body>
</html>
`,
    width: 400,
    height: 200,
    resizable: true,
    debug: true,
    frameless: false,
});

await webview.run();
