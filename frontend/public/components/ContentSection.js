const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    p {
      margin: 0;
      padding: 0;
    }

    section {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      height: 100%;
      gap: var(--vertical-spacing);
    }

    section > header {
      display: flex;
      align-items: center;
      justify-content: space-between;

      h2 {
        font-size: 1.75rem;
      }
    }
  </style>

  <section>
    <dialog>
      <header>
        <h3>Título del modal</h3>
        <close-modal-button></close-modal-button>
      </header>
      <slot name="form"></slot>
    </dialog>

    <header>
      <h2>Título de la sección</h2>
      <open-modal-button>
        <slot name="button-icon"></slot>
        <p>Acción principal</p>
      </open-modal-button>
    </header>
    
    <slot name="content"></slot>
  </section>
`;

export default class ContentSection extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: "open" });
    this.shadowRoot.append(template.content.cloneNode(true));
  }

  static get observedAttributes() {
    return ["hide-modal", "title", "button-text"];
  }

  set hideModal(value) {
    if (value) {
      this.shadowRoot.querySelector("dialog").remove();
    }

    this.setAttribute("hide-modal", value);
  }

  set title(value) {
    const title = this.shadowRoot.querySelector("h2");
    title.innerHTML = value;

    this.setAttribute("title", value);
  }

  set buttonText(value) {
    const buttonText = this.shadowRoot.querySelector("open-modal-button p");
    buttonText.innerHTML = value;

    this.setAttribute("button-text", value);
  }
}

customElements.define("content-section", ContentSection);
