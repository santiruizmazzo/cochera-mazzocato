const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    div {
      min-height: 8rem;
      background-color: var(--clr-bg-light);
      padding: 1.25rem;
    }

    h3, span {
      margin: 0;
    }

    #name {
      color: var(--clr-bg-dark);
    }
  </style>

  <div>
    <h3 id="name"></h3>
    <h3 id="last-name"></h3>
    <span></span>
  </div>
`;

export default class TenantCard extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  static fromJson(jsonTenant) {
    const card = new TenantCard();
    card.dni = jsonTenant["dni"];
    card.name = jsonTenant["name"];
    card.lastName = jsonTenant["last_name"];
    return card;
  }

  connectedCallback() {
    this.render();
  }

  render() {
    this.shadowRoot.querySelector("#name").innerHTML = this.name;
    this.shadowRoot.querySelector("#last-name").innerHTML = this.lastName;
    this.shadowRoot.querySelector("span").innerHTML = `DNI ${this.dni}`;
  }
}

customElements.define("tenant-card", TenantCard);
