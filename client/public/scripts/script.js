// const Http = new XMLHttpRequest();
// const url='/api/';
// Http.open("GET", url);
// Http.send();

// Http.onreadystatechange = (e) => {
//   items = JSON.parse(Http.responseText)
//   console.log(Http.responseText);
//   console.log(items);
// }
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
            <p class="item__location"><label class="localtion__label">Местоположение: </label>
            <span class="localtion__data">${item.location}</span></p>
            <p class="item__found-time"><label class="found-time__label">Найдено в </label>
            <span class="found-time__data">${new Date(item.found_time).toLocaleString()}</span></p>
          </div>
        </div>
      `;
      document.querySelector('.items-row').innerHTML += itemHTML;
    });
  })
  .catch(error => console.error('Error:', error));