const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    * {
      --separation: 0.5rem;
      --border-size: 0.2rem;
      --font-weight: 600;
      --stroke-color: var(--clr-contrast);
      --fill-color: var(--clr-main);
      --icon-side-length: 1.5rem;
      
      font: inherit;
    }

    button {
      display: flex;
      align-items: center;
      justify-content: center;
      gap: var(--separation);

      background-color: var(--fill-color);
      border: var(--border-size) solid var(--stroke-color);
      color: var(--stroke-color);
      max-height: fit-content;
      padding: var(--separation);
      font-weight: var(--font-weight);

      &:hover, &:hover ::slotted(svg) {
        cursor: pointer;
        background-color: var(--stroke-color);
        color: var(--fill-color);
        fill: var(--fill-color);
      }
    }

    ::slotted(svg) {
      fill: var(--stroke-color);
      width: var(--icon-side-length);
      height: var(--icon-side-length);
    }
  </style>

  <button>
    <slot></slot>
  </button>
`;

export default class OpenModalButton extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  handleEvent(event) {
    event.preventDefault();

    const openModalEvent = new CustomEvent("open-modal", {
      bubbles: true,
      composed: true,
    });

    this.dispatchEvent(openModalEvent);
  }

  connectedCallback() {
    this.shadowRoot.querySelector("button").addEventListener("click", this);
  }
}

customElements.define("open-modal-button", OpenModalButton);
