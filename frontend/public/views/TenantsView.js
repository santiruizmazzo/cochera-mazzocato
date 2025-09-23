import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Inquilinos");
  }

  async getHtml() {
    return /*html*/ `
      <button class="open-button">Nuevo inquilino</button>
      <dialog class="modal" closedby="any">
        <form class="tenant-creation-form">
          <h3>Crear nuevo inquilino</h3>
          
          <label for="name">Nombre/s</label>
          <input id="name" name="name" type="text" 
                 required pattern="^[a-zA-ZáéíóúÁÉÍÓÚñÑüÜ\s]{2,50}$"
                 title="Solo letras y espacios, entre 2 y 50 caracteres"
                 placeholder="Pablo">
          
          <label for="last-name">Apellido/s</label>
          <input id="last-name" name="lastName" type="text" 
                 required pattern="^[a-zA-ZáéíóúÁÉÍÓÚñÑüÜ\s]{2,50}$"
                 title="Solo letras y espacios, entre 2 y 50 caracteres"
                 placeholder="Lamponne">
          
          <label for="dni">DNI</label>
          <input id="dni" name="dni" type="text"
                 required pattern="^[1-9][0-9]{0,9}$"
                 title="Solo números entre 1 y 4294967295 (sin puntos ni espacios)"
                 placeholder="20665961">
          
          <label for="email">Email</label>
          <input id="email" name="email" type="email" 
                 maxlength="100"
                 pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
                 title="Email válido con máximo 100 caracteres"
                 placeholder="lamponne@simuladores.com">
          
          <label for="address">Domicilio</label>
          <input id="address" name="address" type="text"
                 minlength="5" maxlength="100"
                 title="Debe tener entre 5 y 100 caracteres"
                 placeholder="Estrada 820, Concepción del Uruguay">
          
          <label for="phone">Teléfono</label>
          <div class="phone-input">
            <select name="dialing-code">
              <option value="+54">🇦🇷 +54</option>
              <option value="+598">🇺🇾 +598</option>
            </select>
            <input id="phone" name="phone" type="tel" maxlength="15" placeholder="3442518388">
          </div>
          
          <label>Mes de ingreso</label>
          <div class="entry-month-input">
            <select name="month" class="month-selector">
              <option value="01">Enero</option>
              <option value="02">Febrero</option>
              <option value="03">Marzo</option>
              <option value="04">Abril</option>
              <option value="05">Mayo</option>
              <option value="06">Junio</option>
              <option value="07">Julio</option>
              <option value="08">Agosto</option>
              <option value="09">Septiembre</option>
              <option value="10">Octubre</option>
              <option value="11">Noviembre</option>
              <option value="12">Diciembre</option>
            </select>
            <select name="year" class="year-selector"></select>
          </div>
          
          <button type="submit">Crear</button>
        </form>
      </dialog>
      <h2 class="section-title">Inquilinos registrados</h2>
      <tenants-list></tenants-list>
    `;
  }

  setUpJavascript() {
    const modal = document.querySelector(".modal");
    const openModal = document.querySelector(".open-button");

    openModal.addEventListener("click", () => {
      modal.showModal();
    });

    // CURRENT MONTH OPTION SELECTED DYNAMICALLY
    const today = new Date();
    const currentMonth = today.getMonth() + 1;
    const monthPadded = String(currentMonth).padStart(2, "0");
    const optionToSelect = document.querySelector(
      `.month-selector option[value="${monthPadded}"]`,
    );
    optionToSelect.setAttribute("selected", true);

    // YEAR SELECTOR DYNAMIC CONSTRUCTION
    const yearSelectorElement = document.querySelector(".year-selector");
    const currentYear = today.getFullYear();
    const yearsToFirstYear = currentYear - 2021;
    const maxYearOptions = 10;
    const totalYearOptions =
      yearsToFirstYear > maxYearOptions ? maxYearOptions : yearsToFirstYear;

    for (let i = 0; i < totalYearOptions; i++) {
      const option = document.createElement("option");
      const stringYear = `${currentYear - i}`;
      option.innerHTML = stringYear;
      option.setAttribute("value", stringYear);
      yearSelectorElement.appendChild(option);
    }

    const form = document.querySelector(".tenant-creation-form");

    form.addEventListener("submit", async (event) => {
      event.preventDefault();

      const formData = new FormData(form);
      const newTenantData = {
        name: formData.get("name"),
        last_name: formData.get("lastName"),
        dni: parseInt(formData.get("dni")),
        email: formData.get("email"),
        address: formData.get("address"),
        phone: `${formData.get("dialing-code")}${formData.get("phone")}`,
        entry_month: `${formData.get("month")}-${formData.get("year")}`,
      };

      fetch(import.meta.env.VITE_API_URL + "/api/tenants", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newTenantData),
      })
        .then((response) => response.json())
        .then((result) => {
          console.log("Success:", result);
        })
        .catch((error) => {
          console.error("Error:", error);
        });
    });
  }
}
