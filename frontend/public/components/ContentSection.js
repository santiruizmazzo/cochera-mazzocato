const template = document.createElement("template");
template.innerHTML = /*html*/ `
  <style>
    p, h2 {
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

    .title {
      display: flex;
      align-items: center;
      gap: 1rem;
    }

    go-back-icon {
      display: none;
    }

    .error-state {
      display: flex;
      align-items: center;
      justify-content: center;
      
      h2 {
        font-size: 4rem;
        color: rgb(198, 11, 11);
      }
    }
  </style>

  <section>
    <custom-modal>
      <slot name="modal-content"></slot>
    </custom-modal>

    <header>
      <div class="title">
        <go-back-icon></go-back-icon>
        <h2>Título de la sección</h2>
      </div>
      
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
  }

  set title(value) {
    const title = this.shadowRoot.querySelector("h2");
    title.innerHTML = value;

    this.setAttribute("title", value);
  }

  set buttonIcon(iconId) {
    const NAMESPACE_URL = "http://www.w3.org/2000/svg";
    const svg = document.createElementNS(NAMESPACE_URL, "svg");
    const use = document.createElementNS(NAMESPACE_URL, "use");

    use.setAttribute("href", `#${iconId}`);
    svg.appendChild(use);

    svg.setAttribute("slot", "button-icon");
    this.appendChild(svg);
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

  set modalTitle(title) {
    const modal = this.shadowRoot.querySelector("custom-modal");
    modal.title = title;
  }

  set modalContent(element) {
    element.setAttribute("slot", "modal-content");
    this.appendChild(element);
  }

  set showGoBackIcon(value) {
    if (value) {
      this.shadowRoot.querySelector("go-back-icon").style.display = "flex";
    }
  }

  set errorMessage(error) {
    const section = this.shadowRoot.querySelector("section");
    const errorMessage = document.createElement("h2");
    errorMessage.innerHTML = error.message;
    section.className = "error-state";
    section.innerHTML = "";
    section.appendChild(errorMessage);
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
