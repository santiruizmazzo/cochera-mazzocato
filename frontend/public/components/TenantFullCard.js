const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    div {
      min-width: 25dvw;
      min-height: 7rem;
      background-color: var(--clr-main-darker);
      padding: 1rem;
    }
  </style>

  <div>
    <h2 id="full-name"></h2>
    <p id="dni"></p>
    <p id="address"></p>
    <p id="phone"></p>
    <p id="email"></p>
    <p id="entry-month"></p>
    <slot></slot>
  </div>
`;

export default class TenantFullCard extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  async connectedCallback() {
    await this.fetchTenant();
    this.render();
  }

  async fetchTenant() {
    const TENANT_URL =
      import.meta.env.VITE_API_URL + `/api/tenants/${this.getAttribute("id")}`;

    await fetch(TENANT_URL)
      .then((response) => response.json())
      .then((json) => {
        this.tenantData = json;
      })
      .catch((error) => {
        console.error("Error fetching tenants:", error);
      });
  }

  render() {
    this.shadowRoot.querySelector("slot").style.display = "none";
    this.shadowRoot.querySelector("#full-name").innerHTML =
      `${this.tenantData["name"]} ${this.tenantData["last_name"]}`;
    this.shadowRoot.querySelector("#dni").innerHTML =
      `DNI ${this.tenantData["dni"]}`;
    this.shadowRoot.querySelector("#address").innerHTML =
      `Domicilio: ${this.tenantData["address"]}`;
    this.shadowRoot.querySelector("#phone").innerHTML =
      `Teléfono: ${this.tenantData["phone"]}`;
    this.shadowRoot.querySelector("#email").innerHTML =
      `Email: ${this.tenantData["email"]}`;
    this.shadowRoot.querySelector("#entry-month").innerHTML =
      `Mes de ingreso: ${this.tenantData["entry_month"]}`;
  }
}

customElements.define("tenant-full-card", TenantFullCard);
