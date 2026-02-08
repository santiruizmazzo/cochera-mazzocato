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
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
    this.tenants = [];
  }

  async handleEvent(event) {
    if (event.type === "tenants:created") {
      const container = this.shadowRoot.querySelector("div");
      const card = TenantMiniCard.fromJson(event.detail);
      container.appendChild(card);
    }
  }

  connectedCallback() {
    if (this.tenants.length != 0) return;

    document.addEventListener("tenants:created", this);

    this.renderPlaceholders();
    // await this.fetchTenants();
    // this.render();
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

  load(tenants) {
    this.tenants = tenants;
    this.render();
  }
}

customElements.define("tenants-collection", TenantsCollection);
