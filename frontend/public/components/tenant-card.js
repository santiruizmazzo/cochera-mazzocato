const template = document.createElement("template");
template.innerHTML = `
    <style>
      .tenant-card {
        background-color: rgba(184, 126, 126, 1);
        padding: 20px;
      }

      p {
        margin: 0;
      }
    </style>

    <div class='tenant-card'></div>
`;

export class TenantCard extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: "closed" });
    this.root.append(template.content.cloneNode(true));
  }

  static get observedAttributes() {
    return [
      "id",
      "dni",
      "name",
      "lastName",
      "email",
      "address",
      "phone",
      "entryMonth",
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
    return this.getAttribute("lastName");
  }

  set lastName(value) {
    this.setAttribute("lastName", value);
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
    return this.getAttribute("entryMonth");
  }

  set entryMonth(value) {
    this.setAttribute("entryMonth", value);
  }

  attributeChangedCallback(attrName, oldVal, newVal) {
    switch (attrName) {
      default:
        const card = this.root.querySelector(".tenant-card");
        const tenantAttribute = document.createElement("p");
        tenantAttribute.innerHTML = newVal;
        card.appendChild(tenantAttribute);
    }
  }
}

customElements.define("tenant-card", TenantCard);
