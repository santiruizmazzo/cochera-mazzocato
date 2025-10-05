const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    a {
      color: inherit;
      text-decoration: none;
      display: block;
      margin: 0;
      padding: 0;
      border: none;
      background: transparent;
      cursor: pointer;
    }

    div {
      min-height: 7rem;
      background-color: var(--clr-main-darker);
      padding: 1rem;
    }

    h3, span {
      margin: 0;
    }

    #name {
      color: var(--clr-main-darkest);
    }
  </style>

  <a>
    <div>
      <h3 id="name"></h3>
      <h3 id="last-name"></h3>
      <span id="id"></span>
    </div>
  </a>
`;

export default class TenantCard extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  static fromJson(jsonTenant) {
    const card = new TenantCard();
    card.id = jsonTenant["id"];
    card.name = jsonTenant["name"];
    card.lastName = jsonTenant["last_name"];
    return card;
  }

  handleEvent(event) {
    event.preventDefault();
    const navigateEvent = new CustomEvent("navigate", {
      bubbles: true,
      composed: true,
      detail: { href: `/inquilinos/${this.id}` },
    });
    this.dispatchEvent(navigateEvent);
  }

  connectedCallback() {
    this.render();
    this.shadowRoot.querySelector("a").addEventListener("click", this);
  }

  render() {
    this.shadowRoot.querySelector("#name").innerHTML = this.name;
    this.shadowRoot.querySelector("#last-name").innerHTML = this.lastName;
    this.shadowRoot.querySelector("#id").innerHTML = `#${this.id}`;
  }
}

customElements.define("tenant-card", TenantCard);
