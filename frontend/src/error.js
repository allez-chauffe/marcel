function showError(error) {
  document.querySelector('#error').innerText = error
  document.querySelector('#error_container').style.visibility = 'visible'
}

function hideError() {
  document.querySelector('#error_container').style.visibility = 'hidden'
}
