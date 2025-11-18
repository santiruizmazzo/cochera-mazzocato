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

    div:not(.loading):hover {
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

    .loading {
      --animation-time: 1.2s;
      --anmtn-clr-base: var(--clr-main-darkest);
      --anmtn-clr-1: rgb(from var(--anmtn-clr-base) r g b / calc(alpha - 0.75));
      --anmtn-clr-2: rgb(from var(--anmtn-clr-base) r g b / calc(alpha - 0.55));
      
      background: linear-gradient(
        90deg,
        var(--anmtn-clr-2) 0%,
        var(--anmtn-clr-1) 50%,
        var(--anmtn-clr-2) 100%
      );
      background-size: 200% 100%;
      animation: pulse var(--animation-time) infinite linear;

      * {
        visibility: hidden;
      }
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

  <a>
    <div class="loading">
      <h3 id="name">Saul</h3>
      <h3 id="last-name">Goodman</h3>
      <svg viewBox="0 -960 960 960">
          <path d="M560-440h200v-80H560zm0-120h200v-80H560zM200-320h320v-22q0-45-44-71.5T360-440t-116 26.5-44 71.5zm160-160q33 0 56.5-23.5T440-560t-23.5-56.5T360-640t-56.5 23.5T280-560t23.5 56.5T360-480M160-160q-33 0-56.5-23.5T80-240v-480q0-33 23.5-56.5T160-800h640q33 0 56.5 23.5T880-720v480q0 33-23.5 56.5T800-160z"/>
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
    if (!this.name || !this.lastName) {
      return;
    }
    this.shadowRoot.querySelector("#name").innerHTML = this.name;
    this.shadowRoot.querySelector("#last-name").innerHTML = this.lastName;
    this.shadowRoot.querySelector("div").className = "";
  }
}

customElements.define("tenant-mini-card", TenantMiniCard);
