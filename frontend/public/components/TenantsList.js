const template = document.createElement("template");
template.innerHTML = /*html*/ `
    <style>
      .tenants-list {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
        gap: 10px;
      }
    </style>

    <div class='tenants-list'></div>
`;

export class TenantsList extends HTMLElement {
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
    return await fetch(window.ENV.API_URL + "/api/tenants")
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
