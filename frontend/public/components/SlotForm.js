const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    * {
      font-family: var(--font-family);
    }

    form {
      display: grid;
      grid-template-columns: 2fr 1fr 2fr;
      gap: 0.85rem;
    }

    select {
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

    .right-arrow {
      justify-self: center;
      width: 2.4rem;
      height: 2.4rem;
    }

    .slot-icon {
      width: 2.5rem;
      height: 2.5rem;
    }

    .slot-number-container {
      justify-self: center;
      display: flex;
      align-items: center;
      justify-items: center;
      gap: 0.5rem;
      font-size: 1.3rem;
      font-weight: bold;
    }

    div:last-child {
      grid-column: 1 / -1;
    }
  </style>

  <form>
    <select name="tenant-id" class="tenant-selector"></select>
    <svg class="right-arrow" viewBox="0 -960 960 960">
      <path d="m600-200-57-56 184-184H80v-80h647L544-704l56-56 280 280z"/>
    </svg>
    <div class="slot-number-container">
      <svg class="slot-icon" viewBox="0 -960 960 960">
        <path d="M160-80q-33 0-56.5-23.5T80-160v-640q0-33 23.5-56.5T160-880h640q33 0 56.5 23.5T880-800v640q0 33-23.5 56.5T800-80zm200-320q-17 0-28.5-11.5T320-440t11.5-28.5T360-480t28.5 11.5T400-440t-11.5 28.5T360-400m240 0q-17 0-28.5-11.5T560-440t11.5-28.5T600-480t28.5 11.5T640-440t-11.5 28.5T600-400M200-516v264q0 14 9 23t23 9h16q14 0 23-9t9-23v-48h400v48q0 14 9 23t23 9h16q14 0 23-9t9-23v-264l-66-192q-5-14-16.5-23t-25.5-9H308q-14 0-25.5 9T266-708zm106-64 28-80h292l28 80z"/>
      </svg>
      <span></span>
    </div>

    <div>
      <activatable-button type="submit">Confirmar</activatable-button>
      <error-box></error-box>
    </div>
  </form>
`;

export default class SlotForm extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
    this.tenants = [];
  }

  static get observedAttributes() {
    return ["slot-id", "slot-number"];
  }

  set slotId(value) {
    this.setAttribute("slot-id", value);
  }

  get slotId() {
    return parseInt(this.getAttribute("slot-id"));
  }

  set slotNumber(value) {
    this.setAttribute("slot-number", value);
    this.renderSlotNumber();
  }

  get slotNumber() {
    return parseInt(this.getAttribute("slot-number"));
  }

  loadData(data) {
    this.slotId = data["slotId"];
    this.slotNumber = data["slotNumber"];
  }

  connectedCallback() {
    this.render();
  }

  render() {
    if (this.tenants.length > 0) return;
    this.generateTenantSelector();
    this.setupFormSubmissionBehavior();
  }

  async generateTenantSelector() {
    const TENANTS_URL = import.meta.env.VITE_API_URL + "/api/tenants";

    await fetch(TENANTS_URL)
      .then((response) => response.json())
      .then((json) => {
        this.tenants = json["data"];
      })
      .catch((error) => {
        console.error("Error fetching tenants:", error);
      });

    const tenantSelector = this.shadowRoot.querySelector(".tenant-selector");

    this.tenants.forEach((tenantJson) => {
      const tenantOption = document.createElement("option");
      tenantOption.innerHTML = `${tenantJson["name"]} ${tenantJson["last_name"]}`;
      tenantOption.setAttribute("value", tenantJson["id"]);
      tenantSelector.appendChild(tenantOption);
    });
  }

  setupFormSubmissionBehavior() {
    const button = this.shadowRoot.querySelector("activatable-button");
    const form = this.shadowRoot.querySelector("form");
    const errorBox = this.shadowRoot.querySelector("error-box");

    button.addEventListener("submit", async (event) => {
      event.preventDefault();

      if (!form.reportValidity()) {
        return;
      }

      button.deactivate();
      const url = `${import.meta.env.VITE_API_URL}/api/slots/${this.slotId}`;
      const slotData = this.createJsonSlot(form);

      fetch(url, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: slotData,
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
          const slotEvent = new CustomEvent("slot:assigned", {
            detail: responseBody,
            bubbles: true,
            composed: true,
          });
          this.dispatchEvent(slotEvent);

          button.activate();
          errorBox.hide();
        })
        .catch((error) => {
          button.activate();
          errorBox.show(error.message);
        });
    });
  }

  createJsonSlot(form) {
    const slotForm = new FormData(form);

    const slot = {
      tenant_id: parseInt(slotForm.get("tenant-id")),
    };

    return JSON.stringify(slot);
  }

  renderSlotNumber() {
    const slotNumberSpan = this.shadowRoot.querySelector(
      ".slot-number-container span",
    );
    slotNumberSpan.innerHTML = `N° ${this.slotNumber}`;
  }

  clear() {
    this.shadowRoot.querySelector("form").reset();
    this.shadowRoot.querySelector("activatable-button").activate();
    this.shadowRoot.querySelector("error-box").hide();
  }
}

customElements.define("slot-form", SlotForm);
