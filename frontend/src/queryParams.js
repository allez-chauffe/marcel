//Parse query parameters from URL and expose them in the location.queryParams variable
window.location.queryParams = {}
window.location.search
  .slice(1)
  .split('&')
  .forEach(function(pair) {
    if (pair === '') return
    const [key, value] = pair.split('=')
    window.location.queryParams[key] = value && decodeURIComponent(value.replace(/\+/g, ' '))
  })
