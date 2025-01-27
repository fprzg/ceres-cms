/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "ui/**/*.{tmpl.html,js}"],
  theme: {
    container: {
      padding: {
        DEFALUT: '15px',
      },
    },
    screens: {
      sm: '640px',
      md: '768px',
      lg: '960px',
      xl: '1200px',
    },
    fontFamily: {
      primary: 'DM Serif Display',
      secondary: 'Jost',
    },
    extend: {
      colors: {
        primary: {
          DEFAULT: '#3056D3',
          hover: '#305FFF',
        },
        secondary: '#13C296',
        accent: {
          DEFAULT: '#FBBF24',
            secondary: '#FBBF2F',
            hover: '#FBBFFF',
        }
      }
    },
  },
  plugins: [],
}

