const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    :host {
      height: 100%;
      min-width: 65%;
      display: flex;
    }

    :host > div {
      display: flex;
      flex-direction: column;
      flex-grow: 1;
      background-color: var(--clr-main-darker);
    }
    
    h2, p {
      margin: 0;
    }

    h2 {
      text-align: center;
      font-size: 3.5rem;
      padding: 1rem 0;
    }

    #info {
      flex-grow: 1;
      padding: 1rem;
    }

    #entry-month {
      text-align: center;
      background-color: var(--clr-main-darkest);
      padding: 0.4rem 0;
    }

    .empty {
      align-items: center;
      justify-content: center;
    }
  </style>

  <div class="empty">
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
    const container = this.shadowRoot.querySelector("div");
    container.innerHTML = "";
    container.className = "loaded";

    const fullNameElement = document.createElement("h2");
    fullNameElement.innerHTML = `${this.tenantData["name"]} ${this.tenantData["last_name"]}`;
    container.appendChild(fullNameElement);

    const infoContainer = document.createElement("div");
    infoContainer.id = "info";

    const dniElement = document.createElement("p");
    dniElement.innerHTML = `DNI: ${this.tenantData["dni"]}`;
    infoContainer.appendChild(dniElement);

    const addressElement = document.createElement("p");
    addressElement.innerHTML = `Domicilio: ${this.tenantData["address"] || "❓"}`;
    infoContainer.appendChild(addressElement);

    const phoneElement = document.createElement("p");
    phoneElement.innerHTML = `Teléfono: ${this.tenantData["phone"] || "❓"}`;
    infoContainer.appendChild(phoneElement);

    const emailElement = document.createElement("p");
    emailElement.innerHTML = `Email: ${this.tenantData["email"] || "❓"}`;
    infoContainer.appendChild(emailElement);

    container.appendChild(infoContainer);

    const entryMonthElement = document.createElement("p");
    entryMonthElement.id = "entry-month";
    entryMonthElement.innerHTML = `Ingresó en ${formatDateMMYYYY(this.tenantData["entry_month"])}`;
    container.appendChild(entryMonthElement);
  }
}

function formatDateMMYYYY(input) {
  const [month, year] = input.split("-").map(Number);
  const date = new Date(year, month - 1);
  return date.toLocaleDateString("es-ES", { month: "long", year: "numeric" });
}

customElements.define("tenant-full-card", TenantFullCard);
