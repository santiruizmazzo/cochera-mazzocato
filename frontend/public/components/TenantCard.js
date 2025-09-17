const template = document.createElement("template");
template.innerHTML = /*html*/ `
    <style>
      .tenant-card {
        background-color: var(--gray-color);
        padding: 20px;
      }

      p {
        margin: 0;
      }

      h3 {
        margin-top: 0;
      }
    </style>

    <div class='tenant-card'></div>
`;

export class TenantCard extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  static get observedAttributes() {
    return [
      "id",
      "dni",
      "name",
      "last-name",
      "email",
      "address",
      "phone",
      "entry-month",
    ];
  }

  get id() {
    return this.getAttribute("id");
  }

  set id(value) {
    this.setAttribute("id", value);
  }

  get dni() {
    return this.getAttribute("dni");
  }

  set dni(value) {
    this.setAttribute("dni", value);
  }

  get name() {
    return this.getAttribute("name");
  }

  set name(value) {
    this.setAttribute("name", value);
  }

  get lastName() {
    return this.getAttribute("last-name");
  }

  set lastName(value) {
    this.setAttribute("last-name", value);
  }

  get email() {
    return this.getAttribute("email");
  }

  set email(value) {
    this.setAttribute("email", value);
  }

  get address() {
    return this.getAttribute("address");
  }

  set address(value) {
    this.setAttribute("address", value);
  }

  get phone() {
    return this.getAttribute("phone");
  }

  set phone(value) {
    this.setAttribute("phone", value);
  }

  get entryMonth() {
    return this.getAttribute("entry-month");
  }

  set entryMonth(value) {
    this.setAttribute("entry-month", value);
  }

  attributeChangedCallback(attrName, oldVal, newVal) {
    const card = this.shadowRoot.querySelector(".tenant-card");

    switch (attrName) {
      case "id":
        break;
      case "dni":
        if (newVal == "null") return;
        const dniElement = document.createElement("p");
        dniElement.innerHTML = `DNI: ${this.dni}`;
        card.appendChild(dniElement);
        break;
      case "name":
      case "last-name":
        if (this.name && this.lastName) {
          const fullNameElement = document.createElement("h3");
          fullNameElement.innerHTML = `${this.name} ${this.lastName}`;
          card.prepend(fullNameElement);
        }
        break;
      case "entry-month":
        const entryMonthElement = document.createElement("p");
        entryMonthElement.innerHTML = `Mes de ingreso: ${this.entryMonth}`;
        card.appendChild(entryMonthElement);
        break;
      case "email":
        if (newVal == "null") return;
        const emailElement = document.createElement("p");
        emailElement.innerHTML = `Email: ${this.email}`;
        card.appendChild(emailElement);
        break;
      case "phone":
        if (newVal == "null") return;
        const phoneElement = document.createElement("p");
        phoneElement.innerHTML = `Teléfono: ${this.phone}`;
        card.appendChild(phoneElement);
        break;
      case "address":
        if (newVal == "null") return;
        const addressElement = document.createElement("p");
        addressElement.innerHTML = `Domicilio: ${this.address}`;
        card.appendChild(addressElement);
        break;
      default:
    }
  }
}

customElements.define("tenant-card", TenantCard);
