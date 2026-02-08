const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    * {
      font-family: var(--font-family);
    }

    form {
      display: grid;
      grid-template-columns: 1fr 3fr;
      gap: 0.85rem;
    }

    input, select {
      border: none;
      background-color: var(--clr-main-darker);
      color: var(--clr-contrast);
      padding: 0.5rem;
      border-radius: 0;
      font-size: 0.9rem;

      &:focus {
        outline: 0.15rem solid var(--clr-contrast);
      }
    }

    select, ::picker(select) {
      appearance: base-select;
    }

    ::picker(select) {
      border: 0.15rem solid var(--clr-contrast);
    }

    span {
      color: red;
      font-weight: bold;
    }

    .phone-input, .entry-month-input {
      display: flex;
      gap: 0.5rem;
    }

    .phone-input input, .month-selector, .year-selector {
      flex-grow: 1;
    }

    div:last-child {
      grid-column: 1 / -1;
    }
  </style>

  <form>
    <label for="name">Nombre/s <span>*</span></label>
    <input id="name" name="name" type="text" 
            required
            minlength="2" maxlength="50"
            pattern="^[A-ZÁÉÍÓÚÑÜ][a-záéíóúñü]*(\\s[A-ZÁÉÍÓÚÑÜ][a-záéíóúñü]*)*$"
            title="Cada nombre empieza con mayúsculas, entre 2 y 50 caracteres"
            placeholder="Pablo">

    <label for="last-name">Apellido/s <span>*</span></label>
    <input id="last-name" name="lastName" type="text"
            required
            minlength="2" maxlength="50"
            pattern="^[A-ZÁÉÍÓÚÑÜ][a-záéíóúñü]*(\\s[A-ZÁÉÍÓÚÑÜ][a-záéíóúñü]*)*$"
            title="Cada apellido empieza con mayúsculas, entre 2 y 50 caracteres"
            placeholder="Lamponne">

    <label for="dni">DNI <span>*</span></label>
    <input id="dni" name="dni" type="text"
            required
            minlength="1" maxlength="10"
            pattern="^[1-9][0-9]{0,9}$"
            title="Solo números entre 1 y 4294967295 (sin puntos ni espacios)"
            placeholder="20665961">

    <label for="email">Email</label>
    <input id="email" name="email" type="email" 
            maxlength="100"
            title="Email válido con máximo 100 caracteres"
            placeholder="lamponne@simuladores.com">

    <label for="address">Domicilio</label>
    <input id="address" name="address" type="text"
            minlength="5" maxlength="100"
            title="Debe tener entre 5 y 100 caracteres"
            placeholder="Estrada 820, Concepción del Uruguay">

    <label for="phone">Teléfono</label>
    <div class="phone-input">
      <select name="dialing-code">
        <option value="+54">🇦🇷 +54</option>
        <option value="+598">🇺🇾 +598</option>
      </select>

      <input id="phone" name="phone" type="tel"
              minlength="8"
              maxlength="15"
              placeholder="3442518388">
    </div>

    <label>Mes de ingreso <span>*</span></label>
    <div class="entry-month-input">
      <select name="month" class="month-selector">
        <option value="01">Enero</option>
        <option value="02">Febrero</option>
        <option value="03">Marzo</option>
        <option value="04">Abril</option>
        <option value="05">Mayo</option>
        <option value="06">Junio</option>
        <option value="07">Julio</option>
        <option value="08">Agosto</option>
        <option value="09">Septiembre</option>
        <option value="10">Octubre</option>
        <option value="11">Noviembre</option>
        <option value="12">Diciembre</option>
      </select>
      
      <select name="year" class="year-selector"></select>
    </div>

    <div>
      <activatable-button type="submit">Confirmar</activatable-button>
      <error-box></error-box>
    </div>
  </form>
`;

export default class TenantForm extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
    this.tenant = null;
  }

  static get observedAttributes() {
    return ["mode"];
  }

  get mode() {
    return this.getAttribute("mode") || "create";
  }

  set mode(value) {
    this.setAttribute("mode", value);
  }

  load(tenant) {
    this.tenant = tenant;
    this.populateForm();
  }

  populateForm() {
    if (!this.tenant) return;

    const form = this.shadowRoot.querySelector("form");
    form.querySelector("#name").value = this.tenant.name;
    form.querySelector("#last-name").value = this.tenant.last_name;
    form.querySelector("#dni").value = this.tenant.dni;
    form.querySelector("#email").value = this.tenant.email || "";
    form.querySelector("#address").value = this.tenant.address || "";

    if (this.tenant.phone) {
      let dialingCode, localPhone;

      if (this.tenant.phone.startsWith("+54")) {
        dialingCode = this.tenant.phone.substring(0, 3);
        localPhone = this.tenant.phone.substring(3);
      }

      if (this.tenant.phone.startsWith("+598")) {
        dialingCode = this.tenant.phone.substring(0, 4);
        localPhone = this.tenant.phone.substring(4);
      }

      form.querySelector("[name='dialing-code']").value = dialingCode;
      form.querySelector("#phone").value = localPhone;
    }

    const [month, year] = this.tenant.entry_month.split("-");

    const monthOption = form.querySelector(
      `[name='month'] option[value="${month}"]`,
    );
    monthOption.selected = true;

    form.querySelector("[name='year']").value = year;
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.selectCurrentMonthOption();
    this.generateYearSelector();
    this.setupFormSubmissionBehavior();
  }

  selectCurrentMonthOption() {
    const currentMonth = new Date().getMonth() + 1;
    const paddedMonth = String(currentMonth).padStart(2, "0");
    const monthOption = this.shadowRoot.querySelector(
      `.month-selector option[value="${paddedMonth}"]`,
    );
    monthOption.selected = true;
  }

  generateYearSelector() {
    const currentYear = new Date().getFullYear();
    const FIRST_YEAR_OF_OPERATION = 2022;
    const MAX_OPTIONS = 10;

    const operativeYears = currentYear - FIRST_YEAR_OF_OPERATION + 1;
    const totalOptions = Math.min(operativeYears, MAX_OPTIONS);

    const yearSelector = this.shadowRoot.querySelector(".year-selector");

    for (let i = 0; i < totalOptions; i++) {
      const stringYear = `${currentYear - i}`;

      const yearOption = document.createElement("option");
      yearOption.innerHTML = stringYear;
      yearOption.setAttribute("value", stringYear);
      yearSelector.appendChild(yearOption);
    }
  }

  setupFormSubmissionBehavior() {
    const form = this.shadowRoot.querySelector("form");
    const button = this.shadowRoot.querySelector("activatable-button");
    const errorBox = this.shadowRoot.querySelector("error-box");
    const TENANTS_URL = import.meta.env.VITE_API_URL + "/api/tenants";

    button.addEventListener("submit", async (event) => {
      event.preventDefault();

      if (!form.reportValidity()) {
        return;
      }

      button.deactivate();

      const method = this.mode === "edit" ? "PATCH" : "POST";
      const url =
        this.mode === "edit" ? `${TENANTS_URL}/${this.tenant.id}` : TENANTS_URL;

      const tenant = this.createJsonTenant(form);

      fetch(url, {
        method: method,
        headers: { "Content-Type": "application/json" },
        body: tenant,
      })
        .then(async (response) => {
          const responseBody = await response.json();

          if (!response.ok) {
            const errorMessage =
              responseBody.detail ||
              `${response.status} (${response.statusText})`;
            throw new Error(errorMessage);
          }

          return responseBody;
        })
        .then((responseBody) => {
          const eventName =
            this.mode === "edit" ? "tenants:updated" : "tenants:created";

          const tenantEvent = new CustomEvent(eventName, {
            detail: responseBody,
            bubbles: true,
            composed: true,
          });
          this.dispatchEvent(tenantEvent);

          if (this.mode === "create") {
            form.reset();
          }

          button.activate();
          errorBox.hide();
        })
        .catch((error) => {
          button.activate();
          errorBox.show(error.message);
        });
    });
  }

  createJsonTenant(form) {
    const tenantForm = new FormData(form);

    const localPhone = tenantForm.get("phone");
    const fullPhone = localPhone
      ? `${tenantForm.get("dialing-code")}${localPhone}`
      : null;

    const tenant = {
      name: tenantForm.get("name"),
      last_name: tenantForm.get("lastName"),
      dni: parseInt(tenantForm.get("dni")),
      email: tenantForm.get("email") ? tenantForm.get("email") : null,
      address: tenantForm.get("address") ? tenantForm.get("address") : null,
      phone: fullPhone,
      entry_month: `${tenantForm.get("month")}-${tenantForm.get("year")}`,
    };

    return JSON.stringify(tenant);
  }

  clear() {
    this.shadowRoot.querySelector("form").reset();
    this.shadowRoot.querySelector("activatable-button").activate();
    this.shadowRoot.querySelector("error-box").hide();
  }
}

customElements.define("tenant-form", TenantForm);
