/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.{templ,html}"],
    theme: {
        colors: {
            transparent: 'transparent',
            current: 'currentColor',
            'white': '#ffffff',
            'disgust': {
                darkest: '#0C3819',
                dark: "#30512b",
                default: "#398f67",
                light: "#76f7aa",
                bright: "#b6dfb0"
            }
        },
        extend: {},
    },
    plugins: [],
}