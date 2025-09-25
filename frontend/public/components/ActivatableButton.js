const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    * {
      box-sizing: border-box;
    }

    button {
      min-width: 100%;
      font-weight: bold;
      font-size: 1rem;
      font-family: var(--font-family);

      &:hover {
        cursor: pointer;
      }
    }
  </style>

  <button>
    <slot></slot>
  </button>
`;

export default class ActivatableButton extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  static get observedAttributes() {
    return ["type"];
  }

  set type(value) {
    this.setAttribute("type", value);
  }

  get type() {
    return this.getAttribute("type");
  }

  attributeChangedCallback(name, old, now) {
    if (name === "type") {
      this.shadowRoot.querySelector("button").type = now;
    }
  }

  handleEvent(event) {
    if (event.type === "click" && this.type === "submit") {
      const newEvent = new Event("submit", {
        bubbles: true,
        composed: true,
        cancelable: true,
      });
      this.dispatchEvent(newEvent);
    }
  }

  connectedCallback() {
    this.shadowRoot.querySelector("button").addEventListener("click", this);
  }

  deactivate() {
    const button = this.shadowRoot.querySelector("button");
    button.disabled = true;
    button.style.cursor = "default";
    button.innerHTML = "Cargando...";
  }

  activate() {
    const button = this.shadowRoot.querySelector("button");
    button.disabled = false;
    button.style.cursor = "pointer";
    button.innerHTML = "Confirmar";
  }
}

customElements.define("activatable-button", ActivatableButton);
