function togglePasswordVisibility(inputId) {
    const input = document.getElementById(inputId);
    const passwordToggle = input.nextElementSibling;
    if (input.type === 'password') {
      input.type = 'text';
      passwordToggle.classList.remove('password-hide');
    } else {
      input.type = 'password';
      passwordToggle.classList.add('password-hide');
    }
  }
