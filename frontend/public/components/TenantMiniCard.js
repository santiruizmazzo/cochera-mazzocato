const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    :host {
      display: flex;
      flex-direction: column;
    }

    a {
      color: inherit;
      text-decoration: none;
      display: flex;
      margin: 0;
      padding: 0;
      border: none;
      background: transparent;
      cursor: pointer;
      flex-grow: 1;
    }

    div {
      display: flex;
      flex-direction: column;
      background-color: var(--clr-main-darker);
      padding: 1rem;
      flex-grow: 1;
      overflow-wrap: anywhere;
    }
    
    #name {
      margin: 0;
      color: var(--clr-main-darkest);
    }

    #last-name {
      margin: 0;
      flex-grow: 1;
    }

    svg {
      --side-size: 1.75rem;
      
      visibility: hidden;
      width: var(--side-size);
      height: var(--side-size);

      fill: var(--clr-main-darker);
      align-self: end;
    }

    div:hover {
      background-color: var(--clr-main-darkest);

      #name {
        color: var(--clr-main-darker);
      }

      #last-name {
        color: var(--clr-main);
      }

      svg {
        visibility: visible;
      }
    }
  </style>

  <a>
    <div>
      <h3 id="name"></h3>
      <h3 id="last-name"></h3>
      <svg viewBox="0 -960 960 960">
        <path d="M200-200v-560 179-19zm80-240h221q2-22 10-42t20-38H280zm0 160h157q17-20 39-32.5t46-20.5q-4-6-7-13t-5-14H280zm0-320h400v-80H280zm-80 480q-33 0-56.5-23.5T120-200v-560q0-33 23.5-56.5T200-840h560q33 0 56.5 23.5T840-760v258q-14-26-34-46t-46-33v-179H200v560h202q-1 6-1.5 12t-.5 12v56zm480-200q-42 0-71-29t-29-71 29-71 71-29 71 29 29 71-29 71-71 29M480-120v-56q0-24 12.5-44.5T528-250q36-15 74.5-22.5T680-280t77.5 7.5T832-250q23 9 35.5 29.5T880-176v56z"/>
      </svg>
    </div>
  </a>
`;

export default class TenantMiniCard extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  static fromJson(jsonTenant) {
    const card = new TenantMiniCard();
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
  }
}

customElements.define("tenant-mini-card", TenantMiniCard);
