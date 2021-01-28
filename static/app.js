const Controller = {
  search: (ev) => {
    ev.preventDefault();
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    const response = fetch(`/search?q=${data.query}`).then((response) => {
      response.json().then((results) => {
        Controller.updateTable(results);
      });
    });
  },

  updateTable: (results) => {
    // console.log("-------");
    // console.log(results);
    const table = document.getElementById("table-body");
    const rows = [];
    for (let result of results) {
      rows.push(`<tr>
        chapter: ${result.chapter}; line-number: ${result.line_number}<br>
        ${result.preview}
      <tr/>`);
    }
    table.innerHTML = rows;
  },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);
