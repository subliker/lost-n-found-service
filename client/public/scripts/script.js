

fetch('/api/items') // Измените URL на адрес вашего REST API сервера
  .then(response => response.json())
  .then(data => {
    console.log(data)
    data.forEach(item => {
      console.log(item)
      const itemHTML = `
        <div class="item">
          <div class="item__info">
            <h2 class="item__title">${item.name}</h2>
            <p class="item__location"><label class="localtion__label">📍 Местоположение: </label>
            <span class="localtion__data">${item.location}</span></p>
            <p class="item__found-time"><label class="found-time__label">🕒 Найдено в </label>
            <span class="found-time__data">${new Date(item.found_time).toLocaleString()}</span></p>
          </div>`+(item.photo_file_name!=""?`
          <img class="item__photo" alt="item_photo" src="/public/${item.photo_file_name}">`:'')+`
        </div>
      `;
      document.querySelector('.items-row').innerHTML += itemHTML;
    });
  })
  .catch(error => console.error('Error:', error));


document.getElementById("form").addEventListener("submit", function(event) {
  event.preventDefault();
  
  var formData = new FormData(document.getElementById("form"));
  formData.set("found_time", new Date(formData.get("found_time").replace(" ", "T") + "Z").toISOString())

  fetch("/api/item", { 
    method: "POST", 
    body: formData
  })
  .then(response => response.text())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
});

document.querySelector("#found-btn").addEventListener("click", function(e) {
  document.querySelector(".found__pop-up").classList.remove("hidden")
  document.querySelector("body").classList.add("no-scroll")
})
document.querySelector("#form-exit").addEventListener("click", function(e) {
  document.querySelector(".found__pop-up").classList.add("hidden")
  document.querySelector("body").classList.remove("no-scroll")
})
document.querySelector("#form-submit").addEventListener("click", function(e) {
  document.querySelector(".found__pop-up").classList.add("hidden")
  document.querySelector("body").classList.remove("no-scroll")
})