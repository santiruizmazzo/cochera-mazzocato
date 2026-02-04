import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor(params) {
    super(params);
  }

  getHtml() {
    return /*html*/ `
      <section class="tenants-view">
        <dialog id="new-tenant-modal">
          <header>
            <h3>Registrar inquilino</h3>
            <close-modal-button></close-modal-button>
          </header>
          <tenant-form mode="create"></tenant-form>
        </dialog>

        <header class="section-header">
          <h2>Inquilinos registrados</h2>
          <open-modal-button>
            <svg><use href="#person_add"></svg>
            Nuevo inquilino
          </open-modal-button>
        </header>
        
        <tenants-collection></tenants-collection>
      </section>
    `;
  }

  setUpJavascript() {
    const view = document.querySelector(".tenants-view");
    const modal = document.querySelector("#new-tenant-modal");
    const form = document.querySelector("tenant-form");

    view.addEventListener("open-modal", () => modal.showModal());
    view.addEventListener("close-modal", () => {
      modal.close();
      form.clear();
    });
    view.addEventListener("tenants:created", () => modal.close());
  }
}
