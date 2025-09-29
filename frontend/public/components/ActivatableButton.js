const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    * {
      box-sizing: border-box;
      font-family: var(--font-family);
      font-size: inherit;
    }

    button {
      --separation: 0.6rem;
      --border-size: 0.2rem;
      --font-weight: 600;
      --stroke-color: var(--clr-contrast);
      --fill-color: var(--clr-main);

      min-width: 100%;
      max-height: fit-content;
      padding: var(--separation);
      font-weight: var(--font-weight);

      display: flex;
      align-items: center;
      justify-content: center;

      border: var(--border-size) solid var(--stroke-color);
      background-color: var(--stroke-color);
      color: var(--fill-color);

      &:hover {
        cursor: pointer;
      }

      svg {
        fill: var(--fill-color);
        width: 1.5rem;
        height: 1.5rem;
      }
    }

    button[disabled] {
      background-color: var(--clr-main-darkest);
      color: var(--clr-contrast);
      border: none;
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
    const newEvent = new Event(this.type, {
      bubbles: true,
      composed: true,
      cancelable: true,
    });
    this.dispatchEvent(newEvent);
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
