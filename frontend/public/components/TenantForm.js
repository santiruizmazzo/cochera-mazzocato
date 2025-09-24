const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    form {
      display: grid;
      grid-template-columns: 1fr 3fr;
      gap: 0.85rem;

      h3 {
        text-align: center;
        grid-column: 1 / 3;
        margin: 0;
      }

      button {
        grid-column: 1 / 3;
        font-weight: bold;
        font-size: 1rem;

        &:hover {
          cursor: pointer;
        }
      }

      * {
        font-family: var(--font-family);
      }
    }

    label {
      span {
        color: red;
      }
    }

    .phone-input {
      display: flex;
      gap: 0.5rem;

      input {
        flex-grow: 4;
      }
      
      select {
        text-align: center;
        flex-grow: 1;
      }
    }

    .entry-month-input {
      display: flex;
      gap: 0.5rem;
    }

    .month-selector {
      font-family: var(--font-family);
      flex-grow: 1;
    }

    .year-selector {
      flex-grow: 1;
    }
  </style>

  <form>
    <h3>Registrar nuevo inquilino</h3>

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
            pattern="^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"
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
              minlength="10"
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

    <button type="submit">Confirmar</button>
  </form>
`;

export default class TenantForm extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
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
    monthOption.setAttribute("selected", true);
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
    const TENANTS_URL = import.meta.env.VITE_API_URL + "/api/tenants";

    form.addEventListener("submit", async (event) => {
      event.preventDefault();

      this.shadowRoot.querySelector("button").innerHTML = "Cargando...";

      const jsonTenant = this.createJsonTenant(form);

      fetch(TENANTS_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: jsonTenant,
      })
        .then((response) => response.json())
        .then((result) => {
          console.log("Success:", result);
          this.parentElement.close();
        })
        .catch((error) => {
          console.error("Error:", error);
        });
    });
  }

  createJsonTenant(form) {
    const tenantForm = new FormData(form);
    const tenant = {
      name: tenantForm.get("name"),
      last_name: tenantForm.get("lastName"),
      dni: parseInt(tenantForm.get("dni")),
      email: tenantForm.get("email"),
      address: tenantForm.get("address"),
      phone: `${tenantForm.get("dialing-code")}${tenantForm.get("phone")}`,
      entry_month: `${tenantForm.get("month")}-${tenantForm.get("year")}`,
    };

    return JSON.stringify(tenant);
  }
}

customElements.define("tenant-form", TenantForm);
