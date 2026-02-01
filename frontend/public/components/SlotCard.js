const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    :host {
      width: 100%;
      display: flex;
      flex-direction: column;
    }

    * {
      margin: 0;
    }

    h2 {
      line-height: 1;
      font-size: 3.5rem;
      font-weight: 500;
    }

    p {
      text-align: end;
    }

    a {
      text-decoration: none;
      font-weight: bold;

      &:hover {
        text-decoration: underline dotted;
        cursor: pointer;
      }
    }

    div {
      padding: 1rem;
      display: flex;
      flex-direction: column;
      gap: 1.5rem;
      container: slot;
    }

    .taken-and-paid {
      background-color: rgb(53, 200, 45);

      * {
        color: rgb(17, 63, 14);
      }
    }

    .taken-and-not-paid {
      background-color: rgb(250, 250, 166);

      * {
        color: rgb(116, 116, 25);
      }
    }

    .free {
      background-color: var(--clr-main-darker);

      * {
        color: var(--clr-main-darkest);
      }
    }

    .invisible {
      display: none;
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

    .assign-tenant-btn {
      background: none;
      border: none;
      display: flex;
      flex-direction: row;
      align-items: center;
      justify-content: center;
      gap: 0.3rem;
      width: fit-content;
      align-self: end;

      &:hover {
        cursor: pointer;
        scale: 1.03;
      }

      svg {
        fill: var(--clr-main-darkest);
        width: 1.7rem;
        height: 1.7rem;
      }
    }
  </style>

  <div class="loading">
    <h2 id="number">1</h2>
    <p id="owner">De <a id="tenant-id" data-link>Lamponne</a></p>
    <button id="assign-tenant" class="invisible">
      Asignar
      <svg viewBox="0 -960 960 960">
        <path d="M360-80v-529q-91-24-145.5-100.5T160-880h80q0 83 53.5 141.5T430-680h100q30 0 56 11t47 32l181 181-56 56-158-158v478h-80v-240h-80v240zm120-640q-33 0-56.5-23.5T400-800t23.5-56.5T480-880t56.5 23.5T560-800t-23.5 56.5T480-720"/>
      </svg>
    </button>
  </div>
`;

export default class SlotCard extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  static fromJson(jsonSlot) {
    const card = new SlotCard();
    card.slotId = jsonSlot["id"];
    card.slotNumber = jsonSlot["number"];
    card.tenantId = jsonSlot["tenant_id"];
    return card;
  }

  handleEvent(event) {
    event.preventDefault();
    let customEvent;

    if (event.target.id === "assign-tenant") {
      customEvent = new CustomEvent("slot:selected", {
        bubbles: true,
        composed: true,
        detail: { slotId: this.slotId, slotNumber: this.slotNumber },
      });
    } else {
      customEvent = new CustomEvent("navigate", {
        bubbles: true,
        composed: true,
        detail: { href: `/inquilinos/${this.tenantId}` },
      });
    }

    this.dispatchEvent(customEvent);
  }

  connectedCallback() {
    this.render();
    this.shadowRoot.querySelector("#tenant-id").addEventListener("click", this);
    this.shadowRoot
      .querySelector("#assign-tenant")
      .addEventListener("click", this);
  }

  render() {
    if (!this.slotId) return;

    if (this.tenantId) {
      this.renderTakenState();
    } else {
      this.renderFreeState();
    }
    this.shadowRoot.querySelector("#number").innerHTML = this.slotNumber;
  }

  renderTakenState() {
    this.shadowRoot.querySelector("div").className = "taken-and-not-paid";
    this.shadowRoot.querySelector("#tenant-id").innerHTML = `#${this.tenantId}`;
  }

  renderFreeState() {
    this.shadowRoot.querySelector("#owner").className = "invisible";
    this.shadowRoot.querySelector("#assign-tenant").className =
      "assign-tenant-btn open-modal-btn";
    this.shadowRoot.querySelector("div").className = "free";
  }
}

customElements.define("slot-card", SlotCard);
