const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    button {
      border: none;
      background-color: transparent;
      display: flex;
      align-items: center;
      justify-content: center;

      &:hover {
        cursor: pointer;
        background-color: var(--clr-main-darker);
      }
    }

    svg {
      --icon-side-length: 1.5rem;

      fill: var(--clr-contrast);
      width: var(--icon-side-length);
      height: var(--icon-side-length);
    }
  </style>

  <button>
    <svg viewBox="0 -960 960 960">
      <path d="m256-200-56-56 224-224-224-224 56-56 224 224 224-224 56 56-224 224 224 224-56 56-224-224z"/>
    </svg>
  </button>
`;

export default class CloseModalButton extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  handleEvent(event) {
    event.preventDefault();

    const closeModalEvent = new CustomEvent("close-modal", {
      bubbles: true,
      composed: true,
    });

    this.dispatchEvent(closeModalEvent);
  }

  connectedCallback() {
    this.shadowRoot.querySelector("button").addEventListener("click", this);
  }
}

customElements.define("close-modal-button", CloseModalButton);
