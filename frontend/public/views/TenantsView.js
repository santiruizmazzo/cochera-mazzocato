import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Inquilinos");
  }

  async getHtml() {
    return /*html*/ `
      <button class="open-button">Nuevo inquilino</button>
      <dialog class="modal">
        <form class="tenant-creation-form">
          <h3>Crear nuevo inquilino</h3>
          <label for="name">Nombre/s</label>
          <input id="name" name="name" type="text">
          <label for="last-name">Apellido/s</label>
          <input id="last-name" name="lastName" type="text">
          <label for="dni">DNI</label>
          <input id="dni" name="dni" type="number">
          <label for="email">Email</label>
          <input id="email" name="email" type="email">
          <label for="address">Domicilio</label>
          <input id="address" name="address" type="text">
          <label for="phone">Teléfono</label>
          <div class="phone-input">
            <select name="dialing-code">
              <option value="+54">🇦🇷 +54</option>
              <option value="+598">🇺🇾 +598</option>
            </select>
            <input id="phone" name="phone" type="tel">
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
            <select name="year" class="year-selector">
              <option value="2025">2025</option>
              <option value="2024">2024</option>
              <option value="2023">2023</option>
              <option value="2022">2022</option>
              <option value="2021">2021</option>
            </select>
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
