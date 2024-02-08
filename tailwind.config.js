/** @type {import("tailwindcss").Config} */
module.exports = {
    content: ["./templates/**/*.{templ,html}"],
    theme: {
        extend: {
            colors: {
                "disgust": {
                    '50': '#f0f9f4',
                    '100': '#dbf0e2',
                    '200': '#bae0ca',
                    '300': '#8cc9a9',
                    '400': '#5bac83',
                    '500': '#398f67',
                    '600': '#287352',
                    '700': '#205c43',
                    '800': '#1c4936',
                    '900': '#183c2e',
                    '950': '#0c221a',
                },
            },
        },
    },
    plugins: [],
}