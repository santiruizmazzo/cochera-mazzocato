import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  getHtml() {
    return /*html*/ `
      <section class="tenants-view">
        <custom-modal title="Registrar inquilino">
          <tenant-form mode="create"></tenant-form>
        </custom-modal>

        <header class="section-header">
          <h2>Inquilinos registrados</h2>
          <open-modal-button>
            <svg><use href="#person_add"></svg>
            Nuevo inquilino
          </open-modal-button>
        </header>
        
        <tenants-collection></tenants-collection>
        <div></div>
      </section>
    `;
  }

  setUpJavascript() {
    const view = document.querySelector(".tenants-view");
    const modal = document.querySelector("custom-modal");

    view.addEventListener("tenants:created", () => modal.close());
  }
}
