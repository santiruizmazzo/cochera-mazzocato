const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    :host {
      display: none;
    }

    :host([open]) {
      display: contents;
    }

    #backdrop {
      position: fixed;
      inset: 0;
      background: rgba(0, 0, 0, 0.509);
      z-index: 998;
    }

    #content {
      position: fixed;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      z-index: 999;

      display: flex;
      flex-direction: column;
      gap: 0.85rem;
      padding: 1.3rem;
      min-width: 25dvw;
      border: none;
      background-color: var(--clr-main);
    }

    header {
      display: grid;
      grid-template-columns: 1fr 6fr 1fr;

      h3 {
        text-align: center;
        grid-column: 2 / 3;
        margin: 0;
        font-weight: bold;
        font-size: 1.3rem;
      }

      close-modal-button {
        grid-column: 3 / 4;
        justify-self: end;
        align-self: center;
      }
    }
  </style>

  <div id="backdrop"></div>
  <div id="content">
    <header>
      <h3>Título del modal</h3>
      <close-modal-button></close-modal-button>
    </header>
    <slot></slot>
  </div>
`;

export default class CustomModal extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  static get observedAttributes() {
    return ["title"];
  }

  set title(value) {
    this.shadowRoot.querySelector("h3").innerHTML = value;
    this.setAttribute("title", value);
  }

  get title() {
    return this.getAttribute("title");
  }

  show() {
    this.setAttribute("open", true);
  }

  close() {
    this.removeAttribute("open");
  }

  connectedCallback() {
    this.title = this.title;

    this.parentElement.addEventListener("open-modal", () => this.show());
    this.addEventListener("close-modal", () => this.close());
  }
}

customElements.define("custom-modal", CustomModal);
