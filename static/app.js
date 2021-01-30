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
    console.log("-------");
    console.log(results);
    const table = document.getElementById("table-body");
    for (let result of results) {
      var row = table.insertRow(0);
      var cell1 = row.insertCell(0);

      // preview = result.preview.replace(/(?:\r\n|\r|\n)/g, "<br>");
      cell1.innerHTML = `<div style="border: thin solid gray; padding: 0.5em 0.8em;">
        Chapter ${result.chapter}; Line #${result.line_number}<br>
        <q>${result.preview}</q>
      </div>`;
    }
  },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);
