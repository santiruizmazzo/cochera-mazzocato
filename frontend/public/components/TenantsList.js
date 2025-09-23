const template = document.createElement("template");
template.innerHTML = /*html*/ `
    <style>
      .tenants-list {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(15.625rem, 1fr));
        gap: 0.625rem;
      }
    </style>

    <div class='tenants-list'></div>
`;

export default class TenantsList extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
    this.tenants = [];
  }

  async connectedCallback() {
    if (this.tenants.length != 0) return;

    this.tenants = await this.fetchTenants();
    this.createList();
  }

  async fetchTenants() {
    return await fetch(import.meta.env.VITE_API_URL + "/api/tenants")
      .then((response) => response.json())
      .then((json) => {
        return json["data"];
      })
      .catch((error) => {
        console.error("Error fetching tenants:", error);
      });
  }

  createList() {
    const list = this.shadowRoot.querySelector(".tenants-list");

    this.tenants.forEach((tenant) => {
      const card = document.createElement("tenant-card");
      card.setAttribute("id", tenant.id);
      card.setAttribute("dni", tenant.dni);
      card.setAttribute("name", tenant.name);
      card.setAttribute("last-name", tenant.last_name);
      card.setAttribute("entry-month", tenant.entry_month);
      card.setAttribute("email", tenant.email);
      card.setAttribute("address", tenant.address);
      card.setAttribute("phone", tenant.phone);
      list.appendChild(card);
    });
  }
}

customElements.define("tenants-list", TenantsList);
