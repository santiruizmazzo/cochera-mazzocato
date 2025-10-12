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
      grid-template-columns: 1.25fr 2fr;
      background-color: var(--clr-main-darker);
    }

    .left-side {
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
      line-height: 1;
      font-weight: 600;
    }
    
    .right-side {
      padding: 1.5rem;
      display: flex;
      flex-direction: column;
      gap: 1rem;
    }

    fieldset {
      margin: 0;
      padding: 0;
      border: none;
    }

    legend {
      padding: 0;
      margin: 0;
      font-size: 1rem;
      font-weight: 500;
      font-style: italic;
      color: var(--clr-main-darkest);
    }

    p {
      margin: 0;
      font-weight: 500;
      font-size: 1.2rem;
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
    <fieldset>
      <legend>DNI</legend>
      <h2 class="dni"></h2>
    </fieldset>
  </div>
  <div class="right-side">
    <fieldset id="last-name">
      <legend>Apellido/s</legend>
    </fieldset>
    <fieldset id="name">
      <legend>Nombre/s</legend>
    </fieldset>
    <fieldset id="entry-month">
      <legend>Mes de ingreso</legend>
    </fieldset>
    <fieldset id="phone">
      <legend>Teléfono</legend>
    </fieldset>
    <fieldset id="email">
      <legend>Email</legend>
    </fieldset>
    <fieldset id="address">
      <legend>Domicilio</legend>
    </fieldset>
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
    dniElement.innerHTML = formatDni(this.tenantData["dni"]);

    const lastNameFieldset = this.shadowRoot.querySelector("#last-name");
    const lastNameElement = document.createElement("p");
    lastNameElement.innerHTML = this.tenantData["last_name"];
    lastNameFieldset.appendChild(lastNameElement);

    const nameFieldset = this.shadowRoot.querySelector("#name");
    const nameElement = document.createElement("p");
    nameElement.innerHTML = this.tenantData["name"];
    nameFieldset.appendChild(nameElement);

    const entryMonthFieldset = this.shadowRoot.querySelector("#entry-month");
    const entryMonthElement = document.createElement("p");
    entryMonthElement.innerHTML = formatEntryMonth(
      this.tenantData["entry_month"],
    );
    entryMonthFieldset.appendChild(entryMonthElement);

    const phoneFieldset = this.shadowRoot.querySelector("#phone");
    const phoneElement = formatPhone(this.tenantData["phone"]);
    phoneFieldset.appendChild(phoneElement);

    const addressFieldset = this.shadowRoot.querySelector("#address");
    const addressElement = document.createElement("p");

    if (this.tenantData["address"]) {
      addressElement.innerHTML = this.tenantData["address"];
    } else {
      addressElement.innerHTML = "Desconocido ¿?";
      addressElement.classList.add("unknown");
    }
    addressFieldset.appendChild(addressElement);

    const emailFieldset = this.shadowRoot.querySelector("#email");
    const emailElement = document.createElement("p");

    if (this.tenantData["email"]) {
      emailElement.innerHTML = this.tenantData["email"];
    } else {
      emailElement.innerHTML = "Desconocido ¿?";
      emailElement.classList.add("unknown");
    }
    emailFieldset.appendChild(emailElement);
  }
}

function formatDni(rawDni) {
  return new Intl.NumberFormat("es-ES").format(rawDni);
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
    yearsFromDate < 1
      ? lessThanAYearComment
      : `hace ${yearsFromDate} año${yearsFromDate > 1 ? "s" : ""}`;

  const commentedDate = `${capitalizedDate} (${comment})`;
  return commentedDate;
}

function formatPhone(rawPhone) {
  const phoneElement = document.createElement("p");

  if (!rawPhone) {
    phoneElement.classList.add("unknown");
    phoneElement.innerHTML = "Desconocido ¿?";
    return phoneElement;
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

  phoneElement.innerHTML = newString;
  return phoneElement;
}

customElements.define("tenant-full-card", TenantFullCard);
