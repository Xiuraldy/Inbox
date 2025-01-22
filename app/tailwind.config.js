module.exports = {
  content: ['./src/**/*.{html,vue,js,ts}', './public/index.html'],
  theme: {
    extend: {
      colors: {
        primary: '#01022e',
        secondary: '#4800FF',
        tertiary: '#12D57D',
        error: '#f91b24'
      },
      width: {
        90: '90%',
        '150px': '150px',
        '100px': '100px'
      },
      height: {
        90: '90%',
        80: '80%',
        70: '70%',
        10: '10%'
      }
    }
  },
  plugins: []
}
