import TenantMiniCard from "./TenantMiniCard.js";

const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    :host {
      flex-grow: 1;
    }

    div {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(13rem, 1fr));
      gap: 0.625rem;
    }
  </style>

  <div></div>
`;

export default class TenantsCollection extends HTMLElement {
  tenants = [];

  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  async handleEvent(event) {
    if (event.type === "tenants:update") {
      const container = this.shadowRoot.querySelector("div");
      const card = TenantMiniCard.fromJson(event.detail);
      container.appendChild(card);
    }
  }

  async connectedCallback() {
    if (this.tenants.length != 0) return;

    document.addEventListener("tenants:update", this);

    this.renderPlaceholders();
    await this.fetchTenants();
    this.render();
  }

  async fetchTenants() {
    const TENANTS_URL = import.meta.env.VITE_API_URL + "/api/tenants";

    await fetch(TENANTS_URL)
      .then((response) => response.json())
      .then((json) => {
        this.tenants = json["data"];
      })
      .catch((error) => {
        console.error("Error fetching tenants:", error);
      });
  }

  renderPlaceholders() {
    const container = this.shadowRoot.querySelector("div");
    const COLLECTION_LENGTH = 15;

    for (let i = 0; i < COLLECTION_LENGTH; i++) {
      const emptyCard = document.createElement("tenant-mini-card");
      container.appendChild(emptyCard);
    }
  }

  render() {
    const container = this.shadowRoot.querySelector("div");
    container.innerHTML = "";

    this.tenants.forEach((tenant) => {
      const card = TenantMiniCard.fromJson(tenant);
      container.appendChild(card);
    });
  }
}

customElements.define("tenants-collection", TenantsCollection);
