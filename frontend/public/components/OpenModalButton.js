const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    
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

  // handleEvent(event) {
  //   const newEvent = new Event(this.type, {
  //     bubbles: true,
  //     composed: true,
  //     cancelable: true,
  //   });
  //   this.dispatchEvent(newEvent);
  // }

  connectedCallback() {
    this.shadowRoot.querySelector("button").addEventListener("click", this);
  }
}

customElements.define("open-modal-button", OpenModalButton);
