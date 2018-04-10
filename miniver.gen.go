package fonts

// name: "Miniver"
// designer: "Open Window"
// license: "OFL"
// category: "DISPLAY"
// date_added: "2011-12-19"
// fonts {
//   name: "Miniver"
//   style: "normal"
//   weight: 400
//   filename: "Miniver-Regular.ttf"
//   post_script_name: "Miniver"
//   full_name: "Miniver"
//   copyright: "Copyright (c) 2011 by Open Window (dathanboardman@gmail.com) with Reseved Font Name \"Miniver\""
// }
// subsets: "menu"
// subsets: "latin"

const MiniverRegular = `@font-face { font-family: 'Miniver'; font-style: normal; font-weight: 400; src: local('Miniver'), url('data:font/woff2;base64,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') format('woff2'); }`;