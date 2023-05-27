/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{vue,js}'],
  theme: {
    extend: {
      colors: {
        primary: '#6062EF',
        'primary-light': '#8687F3',
        grey: '#7A8094',
        'grey-light': '#C1C3CF',
      },
    },
  },
  plugins: [],
};
