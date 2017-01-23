module.exports = {
    entry : "./index/entry.js",
    output : {
        path : __dirname,
        filename : "bundle.js"
    },
    module : {
        loaders : [
        { test : /\.css$/, loader : "style!css" }
        ]
    }
};
