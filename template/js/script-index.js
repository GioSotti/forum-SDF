document.addEventListener('DOMContentLoaded', function() {
  let cards = document.querySelectorAll('.flip-card');

  cards.forEach(function(card) {
    card.addEventListener('click', function(event) {
      if (event.target.classList.contains('button-card') ||
          event.target.classList.contains('comentary') ||
          event.target.classList.contains('comentary-text')||
          event.target.classList.contains('comment-input')  ||
          event.target.classList.contains('form-control')  ||
          event.target.classList.contains('text-container') || 
          event.target.classList.contains('myTextarea') ||
          event.target.classList.contains('submit_commentary')  
          
          
        )
           {
        event.stopPropagation();
      } else {  
        let imageItem = card.querySelector('.image-item');
        if (imageItem && !imageItem.contains(event.target)) {
          if (card.classList.contains('active')) {
            card.classList.remove('active');
          } else {
            card.classList.add('active');
          }
        }
      }
    });
  });

});

/*dropdown*/
function toggleDropdown(event) {
  event.preventDefault();
  var dropdownMenu = document.getElementById("dropdownMenu");
  dropdownMenu.style.display = (dropdownMenu.style.display === "none") ? "block" : "none";
}
let docTitle = document.title;
window.addEventListener("blur", () => {
  document.title = "Reviens chacal !";
});

window.addEventListener("focus", () => {
  document.title = docTitle;
});
var container = document.querySelector('.comentary');

container.addEventListener('scroll', function() {
  var scrollTop = container.scrollTop;
});