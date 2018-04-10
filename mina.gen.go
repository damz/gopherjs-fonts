package fonts

// name: "Mina"
// designer: "Suman Bhandary, Natanael Gama, Mooniak"
// license: "OFL"
// category: "SANS_SERIF"
// date_added: "2017-02-20"
// fonts {
//   name: "Mina"
//   style: "normal"
//   weight: 400
//   filename: "Mina-Regular.ttf"
//   post_script_name: "Mina-Regular"
//   full_name: "Mina"
//   copyright: "Copyright 2015 The Mina Project Authors, with Reserved Font Name \'Exo\'"
// }
// fonts {
//   name: "Mina"
//   style: "normal"
//   weight: 700
//   filename: "Mina-Bold.ttf"
//   post_script_name: "Mina-Bold"
//   full_name: "Mina Bold"
//   copyright: "Copyright 2015 The Mina Project Authors, with Reserved Font Name \'Exo\'"
// }
// subsets: "bengali"
// subsets: "latin"
// subsets: "latin-ext"
// subsets: "menu"

const MinaRegular = `@font-face { font-family: 'Mina'; font-style: normal; font-weight: 400; src: local('Mina'), local('Mina-Regular'), url('data:font/woff2;base64,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') format('woff2'); }`;
const Mina700 = `@font-face { font-family: 'Mina'; font-style: normal; font-weight: 700; src: local('Mina Bold'), local('Mina-Bold'), url('data:font/woff2;base64,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') format('woff2'); }`;