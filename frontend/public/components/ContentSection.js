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
    <custom-modal>
      <slot name="modal-content"></slot>
    </custom-modal>

    <header>
      <h2>Título de la sección</h2>
      <open-modal-button>
        <slot name="button-icon"></slot>
        <p>Acción principal</p>
      </open-modal-button>
    </header>
    
    <slot name="content"></slot>
    <div></div>
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

  get modal() {
    return this.shadowRoot.querySelector("custom-modal");
  }

  set hideModal(value) {
    if (value) {
      this.shadowRoot.querySelector("dialog").remove();
    }

    this.setAttribute("hide-modal", value);
  }

  set hideButton(value) {
    if (value) {
      this.shadowRoot.querySelector("open-modal-button").remove();
    }

    this.setAttribute("hide-button", value);
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

  set content(element) {
    element.setAttribute("slot", "content");
    this.appendChild(element);
  }

  set modalContent(element) {
    element.setAttribute("slot", "modal-content");
    this.appendChild(element);
  }

  connectedCallback() {
    if (this.hideModal) {
      return;
    }
    const section = this.shadowRoot.querySelector("section");
    const modal = this.shadowRoot.querySelector("custom-modal");
    const form = this.shadowRoot
      .querySelector('slot[name="modal-content"]')
      .assignedElements()[0];

    section.addEventListener("open-modal", (event) => {
      if (event.detail) {
        form.loadData(event.detail);
      }
      modal.show();
    });
    section.addEventListener("close-modal", () => modal.close());
  }
}

customElements.define("content-section", ContentSection);
