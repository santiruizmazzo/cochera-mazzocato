const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    * {
      box-sizing: border-box;
    }
    
    :host {
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
      background-color: var(--clr-main-darkest);
    }

    .dni {
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
      overflow-wrap: anywhere;
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

    .loading {
      --anmtn-clr-base: var(--clr-main-darkest);
      --anmtn-clr-1: rgb(from var(--anmtn-clr-base) r g b / calc(alpha - 0.75));
      --anmtn-clr-2: rgb(from var(--anmtn-clr-base) r g b / calc(alpha - 0.55));

      color: transparent;
      background: linear-gradient(
        90deg,
        var(--anmtn-clr-2) 0%,
        var(--anmtn-clr-1) 50%,
        var(--anmtn-clr-2) 100%
      );
      background-size: 200% 100%;
      animation: pulse 1s infinite linear;
    }

    @keyframes pulse {
      0% {
        background-position: 0% 0;
      }
      100% {
        background-position: -200% 0;
      }
    }
  </style>

  <div class="left-side">
    <div class="photo-placeholder"></div>
    <fieldset>
      <legend>DNI</legend>
      <p class="dni loading">99.999.999</p>
    </fieldset>
  </div>

  <div class="right-side">
    <fieldset id="last-name">
      <legend>Apellido/s</legend>
      <p class="loading">Sonnet</p>
    </fieldset>

    <fieldset id="name">
      <legend>Nombre/s</legend>
      <p class="loading">Claude</p>
    </fieldset>

    <fieldset id="entry-month">
      <legend>Mes de ingreso</legend>
      <p class="loading">Marzo de 2023</p>
    </fieldset>

    <fieldset id="phone">
      <legend>Teléfono</legend>
      <p class="loading">+546666888888</p>
    </fieldset>

    <fieldset id="email">
      <legend>Email</legend>
      <p class="loading">claude@anthropic.com</p>
    </fieldset>

    <fieldset id="address">
      <legend>Domicilio</legend>
      <p class="loading">74 Lincoln Ave., San Francisco</p>
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
    this.renderDni();
    this.renderLastName();
    this.renderName();
    this.renderEntryMonth();
    this.renderPhone();
    this.renderAddress();
    this.renderEmail();
  }

  renderDni() {
    const dni = this.shadowRoot.querySelector(".dni");
    dni.innerHTML = formatDni(this.tenantData["dni"]);
    dni.className = "dni";
  }

  renderLastName() {
    const lastName = this.shadowRoot.querySelector("#last-name p");
    lastName.innerHTML = this.tenantData["last_name"];
    lastName.className = "";
  }

  renderName() {
    const name = this.shadowRoot.querySelector("#name p");
    name.innerHTML = this.tenantData["name"];
    name.className = "";
  }

  renderEntryMonth() {
    const entryMonth = this.shadowRoot.querySelector("#entry-month p");
    entryMonth.innerHTML = formatEntryMonth(this.tenantData["entry_month"]);
    entryMonth.className = "";
  }

  renderPhone() {
    const phone = this.shadowRoot.querySelector("#phone p");
    const formattedPhone = formatPhone(this.tenantData["phone"]);

    if (!formattedPhone) {
      phone.innerHTML = "Desconocido";
      phone.className = "unknown";
      return;
    }

    phone.innerHTML = formattedPhone;
    phone.className = "";
  }

  renderAddress() {
    const address = this.shadowRoot.querySelector("#address p");

    if (!this.tenantData["address"]) {
      address.innerHTML = "Desconocido";
      address.className = "unknown";
      return;
    }

    address.innerHTML = this.tenantData["address"];
    address.className = "";
  }

  renderEmail() {
    const email = this.shadowRoot.querySelector("#email p");

    if (!this.tenantData["email"]) {
      email.innerHTML = "Desconocido";
      email.className = "unknown";
      return;
    }

    email.innerHTML = this.tenantData["email"];
    email.className = "";
  }
}

function formatDni(rawDni) {
  return new Intl.NumberFormat("es-ES").format(rawDni);
}

function formatEntryMonth(rawEntryMonth) {
  const [month, year] = rawEntryMonth.split("-").map(Number);
  const date = new Date(year, month - 1);
  const monthsFromEntry = new Date().getMonth() - date.getMonth();
  const yearsFromEntry = new Date().getFullYear() - date.getFullYear();

  const baseDate = date.toLocaleDateString("es-ES", {
    month: "long",
    year: "numeric",
  });
  const capitalizedDate = baseDate.charAt(0).toUpperCase() + baseDate.slice(1);

  const lessThanAYearComment =
    monthsFromEntry < 1 ? "este mes" : `hace ${monthsFromEntry} meses`;

  const moreThanAYearComment = `hace ${yearsFromEntry} año${yearsFromEntry > 1 ? "s" : ""}`;

  const comment =
    yearsFromEntry < 1 ? lessThanAYearComment : moreThanAYearComment;

  return `${capitalizedDate} (${comment})`;
}

function formatPhone(rawPhone) {
  if (!rawPhone) {
    return rawPhone;
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
