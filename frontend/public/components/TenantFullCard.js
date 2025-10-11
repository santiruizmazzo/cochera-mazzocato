const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    * {
      box-sizing: border-box;
    }
    
    :host {
      aspect-ratio: 16 / 9;
      min-width: 65%;
      display: grid;
      grid-template-columns: 1fr 2fr;
    }

    .left-side {
      background-color: var(--clr-main-darker);
      display: flex;
      flex-direction: column;
      padding: 1.5rem;
      gap: 1rem;
    }

    .photo-placeholder {
      flex-grow: 1;
      background-image: linear-gradient(45deg, var(--clr-main), var(--clr-main-darkest));
    }

    .dni {
      display: block;
      margin: 0;
      padding: 0;
      font-size: 2rem;
      text-align: center;
      line-height: 1;
      font-weight: 600;
    }
    
    .right-side {
      background-color: var(--clr-main-darker);
      padding: 1.5rem;
    }

    p, h3 {
      margin: 0;
    }

    h3 {
      margin-bottom: 1rem;
    }

    .unknown {
      display: inline-block;
      background-color: var(--clr-main-darkest);
      color: var(--clr-main);
      font-size: 0.9rem;
      padding: 0.2rem;
      font-weight: 500;
      line-height: 1;
    }
  </style>

  <div class="left-side">
    <div class="photo-placeholder"></div>
    <h2 class="dni"></h2>
  </div>
  <div class="right-side">
    <p>Apellido/s</p>
    <h3 class="last-name"></h3>
    <p>Nombre/s</p>
    <h3 class="name"></h3>
    <p>Mes de ingreso</p>
    <h3 class="entry-month"></h3>
    <p>Teléfono</p>
    <h3 class="phone"></h3>
    <p>Email</p>
    <h3 class="email"></h3>
    <p>Domicilio</p>
    <h3 class="address"></h3>
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
    const dniElement = this.shadowRoot.querySelector(".dni");
    dniElement.innerHTML = this.tenantData["dni"];

    const nameElement = this.shadowRoot.querySelector(".name");
    nameElement.innerHTML = this.tenantData["name"];

    const lastNameElement = this.shadowRoot.querySelector(".last-name");
    lastNameElement.innerHTML = this.tenantData["last_name"];

    const entryMonthElement = this.shadowRoot.querySelector(".entry-month");
    entryMonthElement.innerHTML = formatEntryMonth(
      this.tenantData["entry_month"],
    );

    const phoneElement = this.shadowRoot.querySelector(".phone");
    phoneElement.innerHTML = formatPhone(this.tenantData["phone"]);

    // if (this.tenantData["phone"]) {
    //   phoneElement.innerHTML = this.tenantData["phone"];
    // } else {
    //   const unknownPlaceholder = document.createElement("p");
    //   unknownPlaceholder.classList.add("unknown");
    //   unknownPlaceholder.innerHTML = "Desconocido";
    //   phoneElement.appendChild(unknownPlaceholder);
    // }

    const addressElement = this.shadowRoot.querySelector(".address");

    if (this.tenantData["address"]) {
      addressElement.innerHTML = this.tenantData["address"];
    } else {
      const unknownPlaceholder = document.createElement("p");
      unknownPlaceholder.classList.add("unknown");
      unknownPlaceholder.innerHTML = "Desconocido";
      addressElement.appendChild(unknownPlaceholder);
    }

    const emailElement = this.shadowRoot.querySelector(".email");

    if (this.tenantData["email"]) {
      emailElement.innerHTML = this.tenantData["email"];
    } else {
      const unknownPlaceholder = document.createElement("p");
      unknownPlaceholder.classList.add("unknown");
      unknownPlaceholder.innerHTML = "Desconocido";
      emailElement.appendChild(unknownPlaceholder);
    }
  }
}

function formatEntryMonth(rawEntryMonth) {
  const [month, year] = rawEntryMonth.split("-").map(Number);
  const date = new Date(year, month - 1);
  const monthsFromDate = new Date().getMonth() - date.getMonth();
  const yearsFromDate = new Date().getFullYear() - date.getFullYear();

  const baseDate = date.toLocaleDateString("es-ES", {
    month: "long",
    year: "numeric",
  });
  const capitalizedDate = baseDate.charAt(0).toUpperCase() + baseDate.slice(1);

  const lessThanAYearComment =
    monthsFromDate < 1 ? "este mes" : `hace ${monthsFromDate} meses`;
  const comment =
    yearsFromDate < 1 ? lessThanAYearComment : `hace ${yearsFromDate} años`;

  const commentedDate = `${capitalizedDate} (${comment})`;
  return commentedDate;
}

function formatPhone(rawPhone) {
  if (!rawPhone) {
    // const unknownPlaceholder = document.createElement("p");
    // unknownPlaceholder.classList.add("unknown");
    // unknownPlaceholder.innerHTML = "Desconocido";
    // return unknownPlaceholder;
    // phoneElement.appendChild(unknownPlaceholder);
    return "Desconocido";
  }

  let newString;

  switch (true) {
    case rawPhone.startsWith("+54"):
      newString = `🇦🇷 +54 ${rawPhone.slice(3)}`;
      break;
    case rawPhone.startsWith("+598"):
      newString = `🇺🇾 +598 ${rawPhone.slice(4)}`;
      break;
  }
  return newString;
}

customElements.define("tenant-full-card", TenantFullCard);
