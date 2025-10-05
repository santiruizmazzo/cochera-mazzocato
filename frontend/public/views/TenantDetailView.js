import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  async getHtml() {
    const TENANT_URL =
      import.meta.env.VITE_API_URL + `/api/tenants/${this.params.id}`;
    let tenantData = {};

    await fetch(TENANT_URL)
      .then((response) => response.json())
      .then((json) => {
        tenantData = json;
      })
      .catch((error) => {
        console.error("Error fetching tenants:", error);
      });

    return /*html*/ `
      <section class="tenant-detail-view">
        <h2>${tenantData["name"]} ${tenantData["last_name"]}</h2>
        <p>DNI: ${tenantData["dni"]}</p>
        <p>Teléfono: ${tenantData["phone"]}</p>
      </section>
    `;
  }
}
